package middleware

import (
	"backend/common/errorx"
	"encoding/json"
	"errors"
	"fmt"
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

type CheckLoginMiddleware struct {
}

func NewCheckLoginMiddleware() *CheckLoginMiddleware {
	return &CheckLoginMiddleware{}
}

func (m *CheckLoginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleIdNumber := json.Number(fmt.Sprintf("%v", r.Context().Value("roleId")))
		roleId, err := roleIdNumber.Int64()
		fmt.Println(roleId, r.RequestURI, r.Method)

		if err != nil {
			httpx.Error(w, errorx.NewCodeError(401, fmt.Sprintf("%v", err), ""))
			return
		}
		// todo: 权限放到redis
		if ok, err := check(strconv.Itoa(int(roleId)), r.RequestURI, r.Method); ok == true {
			next(w, r)
		} else {
			httpx.Error(w, errorx.NewCodeError(403, fmt.Sprintf("%v", err), ""))
			return
		}
	}
}

// https://github.com/Blank-Xu/sqlx-adapter
func finalizer(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
func check(sub, obj, act string) (bool, error) {
	m, err := model.NewModelFromString(`
											[request_definition]
											r = sub, obj, act
											
											[policy_definition]
											p = sub, obj, act
											
											[role_definition]
											g = _, _
											
											[policy_effect]
											e = some(where (p.eft == allow))
											
											[matchers]
											m = r.sub == p.sub && ParamsMatch(r.obj,p.obj) && r.act == p.act
											`)

	if err != nil {
		logx.Infof("error: model: %s", err)
		return false, errors.New("model-err")
	}
	//todo: 如何直接引用配置中的信息
	db, err := sqlx.Connect("mysql", "root:123456@tcp(192.168.1.2:3306)/db_zeromicro?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true")
	db.SetMaxIdleConns(20)
	db.SetMaxIdleConns(10)
	runtime.SetFinalizer(db, finalizer)
	a, err := sqlxadapter.NewAdapter(db, "system_role_apis")
	if err != nil {
		//panic(err)
		logx.Infof("error: NewAdapter: %s", err)
		return false, errors.New("NewAdapter-err")
	}

	e, err := casbin.NewEnforcer(m, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)

	if err != nil {
		logx.Infof("error: NewEnforcer: %s", err)
		return false, errors.New("NewEnforcer-err")
	}
	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// handle err
		logx.Infof("error: Enforce: %s", err)
		return false, errors.New("Enforce-err")
	}

	if ok == true {
		// permit alice to read data1
		//fmt.Println("-------check permit--------")
		return true, nil
	} else {
		// deny the request, show an error
		//fmt.Println("-------check deny--------")
		return false, errors.New("未授权访问")
	}
}
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	//fmt.Println("---------ParamsMatchFunc----------")
	name1 := args[0].(string)
	name2 := args[1].(string)
	//fmt.Println(name1, name2)

	return ParamsMatch(name1, name2), nil
}
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	//fmt.Println(fullNameKey1, key2)
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}
