package model

import (
	"backend/common/utils"
	"backend/service/backend/model"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
)

var (
	verifyCodesRows                = "`id`,`created_at`,`updated_at`,`deleted_at`,`account`,`code`,`type`,`status`"
	verifyCodesRowsExpectAutoSet   = "`account`,`code`,`type`,`status`"
	verifyCodesRowsWithPlaceHolder = "`account`=?,`code`=?,`type`=?,`status`=?"

	cacheVerifyCodesIdPrefix = "cache#verifyCodes#id#"
)

type (
	VerifyCodesModel interface {
		Insert(data VerifyCodes) (sql.Result, error)
		FindOne(id int64) (*VerifyCodes, error)
		Update(data VerifyCodes) error
		Delete(id int64) error
		CheckCode(account string, code string, availableTime string) (VerifyCodes, error)
		List(req utils.ListReq) ([]*VerifyCodes, int, error)
		DeleteBatch(ids string) error
		CheckDuplicate(fieldname string) (VerifyCodes, error)
	}

	defaultVerifyCodesModel struct {
		sqlc.CachedConn
		table string
	}

	VerifyCodes struct {
		Id        int64  `db:"id" json:"id"`                 //
		CreatedAt string `db:"created_at" json:"created_at"` //
		UpdatedAt string `db:"updated_at" json:"updated_at"` //
		DeletedAt string `db:"deleted_at" json:"deleted_at"` //
		Account   string `db:"account" json:"account"`       // 号码(手机或邮箱)
		Code      string `db:"code" json:"code"`             // 验证码
		Type      int64  `db:"type" json:"type"`             // 类型0手机1邮箱
		Status    int64  `db:"status" json:"status"`         // 状态0未验证1已验证2验证错误

	}
)

// 验证码
func NewVerifyCodesModel(conn sqlx.SqlConn, c cache.CacheConf) VerifyCodesModel {
	return &defaultVerifyCodesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`verify_codes`",
	}
}

func (m *defaultVerifyCodesModel) Insert(data VerifyCodes) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?,?,?,?)", m.table, verifyCodesRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Account, data.Code, data.Type, data.Status)

	return ret, err
}

func (m *defaultVerifyCodesModel) FindOne(id int64) (*VerifyCodes, error) {
	verifyCodesIdKey := fmt.Sprintf("%s%v", cacheVerifyCodesIdPrefix, id)
	var resp VerifyCodes
	err := m.QueryRow(&resp, verifyCodesIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", verifyCodesRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVerifyCodesModel) Update(data VerifyCodes) error {
	verifyCodesIdKey := fmt.Sprintf("%s%v", cacheVerifyCodesIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, verifyCodesRowsWithPlaceHolder)
		return conn.Exec(query, data.Account, data.Code, data.Type, data.Status, data.Id)
	}, verifyCodesIdKey)
	return err
}

func (m *defaultVerifyCodesModel) Delete(id int64) error {
	verifyCodesIdKey := fmt.Sprintf("%s%v", cacheVerifyCodesIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `deleted_at`=? where `id` = ?", m.table)
		return conn.Exec(query, time.Now().Format("2006-01-02 15:04:05"), id)
	}, verifyCodesIdKey)
	return err
}

func (m *defaultVerifyCodesModel) List(req utils.ListReq) ([]*VerifyCodes, int, error) {
	total := 0

	// 条件处理
	whereCondition := "where " + softDeleteFlag
	if req.Keyword != "" {
		whereCondition += "and `account` like '%" + req.Keyword + "%'"
	}

	orderBy := "order by id desc"
	items := make([]*VerifyCodes, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT ? OFFSET ?", verifyCodesRows, m.table, whereCondition, orderBy)
	queryCount := fmt.Sprintf("SELECT count(1) FROM %s %s", m.table, whereCondition)
	err := m.CachedConn.QueryRowNoCache(&total, queryCount)

	// 查询错误
	if err != nil {
		return items, total, err
	}

	// 没有记录
	if total == 0 {
		return items, total, nil
	}

	//获取记录
	err = m.CachedConn.QueryRowsNoCache(&items, query, req.PageSize, req.PageSize*(req.Page-1))
	if err != nil {
		logx.Errorf("usersSex.findOne error, sex=%d, err=%s", req.Page, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, total, model.ErrNotFound
		}
		return nil, total, err
	}

	return items, total, nil
}
func (m *defaultVerifyCodesModel) DeleteBatch(ids string) error {
	query := fmt.Sprintf("update %s set `deleted_at`=? where `id` in (%s)", m.table, ids)
	_, err := m.ExecNoCache(query, time.Now().Format("2006-01-02 15:04:05"))

	// 删除缓存
	idArr := strings.Split(ids, ",")
	for _, v := range idArr {
		systemUserIdKey := fmt.Sprintf("%s%v", cacheVerifyCodesIdPrefix, v)
		m.DelCache(systemUserIdKey)
	}
	return err
}
func (m *defaultVerifyCodesModel) CheckDuplicate(fieldname string) (VerifyCodes, error) {
	var resp VerifyCodes
	queryString := fmt.Sprintf("select %s from %s where %s and `name` = ? limit 1", verifyCodesRows, m.table, softDeleteFlag)
	err := m.CachedConn.QueryRowNoCache(&resp, queryString, fieldname)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, model.ErrNotFound
	default:
		return resp, err
	}
}
func (m *defaultVerifyCodesModel) CheckCode(account string, code string, availableTime string) (VerifyCodes, error) {
	var resp VerifyCodes
	queryString := fmt.Sprintf("select %s from %s where %s and `status` = 0 and `code` = ? and `account` = ? and `created_at` > ? order by `id` desc limit 1", verifyCodesRows, m.table, softDeleteFlag)
	//fmt.Println(queryString)
	err := m.CachedConn.QueryRowNoCache(&resp, queryString, code, account, availableTime)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, model.ErrNotFound
	default:
		return resp, err
	}
}
