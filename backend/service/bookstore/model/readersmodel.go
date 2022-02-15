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
	readersRows                = "`id`,`username`,`password`,`nickname`,`phone`,`email`,`status`,`login_ip`,`login_date`,`remark`,`created_at`,`updated_at`,`deleted_at`"
	readersRowsExpectAutoSet   = "`username`,`password`,`nickname`,`phone`,`email`,`status`,`login_ip`,`login_date`,`remark`"
	readersRowsWithPlaceHolder = "`username`=?,`password`=?,`nickname`=?,`phone`=?,`email`=?,`status`=?,`login_ip`=?,`login_date`=?,`remark`=?"

	cacheReadersIdPrefix = "cache#readers#id#"
	softDeleteFlag       = "`deleted_at`='2006-01-02 15:04:05'"
)

type (
	ReadersModel interface {
		Insert(data Readers) (sql.Result, error)
		FindOne(id int64) (*Readers, error)
		Update(data Readers) error
		Delete(id int64) error
		List(req utils.ListReq) ([]*Readers, int, error)
		DeleteBatch(ids string) error
		CheckDuplicate(fieldname string, email string) (Readers, error)
		FindOneByUserName(username string) (Readers, error)
		FindOneByUserEmail(email string) (Readers, error)
		FindByIds(ids string) ([]*Readers, error)
	}

	defaultReadersModel struct {
		sqlc.CachedConn
		table string
	}
	//ALTER TABLE `readers` ADD unique(`username`);
	//ALTER TABLE `readers` ADD unique(`email`);
	Readers struct {
		Id        int64  `db:"id" json:"id"`                 // 用户ID
		Username  string `db:"username" json:"username"`     // 用户账号
		Password  string `db:"password" json:"password"`     // 密码
		Nickname  string `db:"nickname" json:"nickname"`     // 用户昵称
		Phone     string `db:"phone" json:"phone"`           // 手机
		Email     string `db:"email" json:"email"`           // 用户邮箱
		Status    int64  `db:"status" json:"status"`         // 帐号状态（0正常 1停用）
		LoginIp   string `db:"login_ip" json:"login_ip"`     // 最后登录IP
		LoginDate string `db:"login_date" json:"login_date"` // 最后登录时间
		Remark    string `db:"remark" json:"remark"`         // 备注
		CreatedAt string `db:"created_at" json:"created_at"` // 创建时间
		UpdatedAt string `db:"updated_at" json:"updated_at"` // 更新时间
		DeletedAt string `db:"deleted_at" json:"deleted_at"` // 删除时间

	}
)

// 读者
func NewReadersModel(conn sqlx.SqlConn, c cache.CacheConf) ReadersModel {
	return &defaultReadersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`readers`",
	}
}

func (m *defaultReadersModel) Insert(data Readers) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?,?,?,?,?,?,?,?,?)", m.table, readersRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Username, data.Password, data.Nickname, data.Phone, data.Email, data.Status, data.LoginIp, data.LoginDate, data.Remark)

	return ret, err
}

func (m *defaultReadersModel) FindOne(id int64) (*Readers, error) {
	readersIdKey := fmt.Sprintf("%s%v", cacheReadersIdPrefix, id)
	var resp Readers
	err := m.QueryRow(&resp, readersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", readersRows, m.table)
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

func (m *defaultReadersModel) Update(data Readers) error {
	readersIdKey := fmt.Sprintf("%s%v", cacheReadersIdPrefix, data.Id)
	readerNameKey := fmt.Sprintf("%s%v", cacheReadersIdPrefix, data.Username)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, readersRowsWithPlaceHolder)
		return conn.Exec(query, data.Username, data.Password, data.Nickname, data.Phone, data.Email, data.Status, data.LoginIp, data.LoginDate, data.Remark, data.Id)
	}, readersIdKey, readerNameKey)
	return err
}

func (m *defaultReadersModel) Delete(id int64) error {
	readersIdKey := fmt.Sprintf("%s%v", cacheReadersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `deleted_at`=? where `id` = ?", m.table)
		return conn.Exec(query, time.Now().Format("2006-01-02 15:04:05"), id)
	}, readersIdKey)
	return err
}

func (m *defaultReadersModel) List(req utils.ListReq) ([]*Readers, int, error) {
	total := 0

	// 条件处理
	whereCondition := "where " + softDeleteFlag
	if req.Keyword != "" {
		whereCondition += "and `username` like '%" + req.Keyword + "%'"
	}

	orderBy := "order by id desc"
	items := make([]*Readers, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT ? OFFSET ?", readersRows, m.table, whereCondition, orderBy)
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
func (m *defaultReadersModel) FindByIds(ids string) ([]*Readers, error) {
	// 条件处理
	whereCondition := "where " + softDeleteFlag
	whereCondition += "and `id` in (?)"

	orderBy := "order by id desc"
	items := make([]*Readers, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s ", readersRows, m.table, whereCondition, orderBy)

	//获取记录
	err := m.CachedConn.QueryRowsNoCache(&items, query, ids)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return items, nil
}

func (m *defaultReadersModel) DeleteBatch(ids string) error {
	query := fmt.Sprintf("update %s set `deleted_at`=? where `id` in (%s)", m.table, ids)
	_, err := m.ExecNoCache(query, time.Now().Format("2006-01-02 15:04:05"))

	// 删除缓存
	idArr := strings.Split(ids, ",")
	for _, v := range idArr {
		systemUserIdKey := fmt.Sprintf("%s%v", cacheReadersIdPrefix, v)
		m.DelCache(systemUserIdKey)
	}
	return err
}
func (m *defaultReadersModel) CheckDuplicate(fieldname string, email string) (Readers, error) {
	var resp Readers
	queryString := fmt.Sprintf("select %s from %s where %s and ( `username` = ? or `email` = ? ) limit 1", readersRows, m.table, softDeleteFlag)
	err := m.CachedConn.QueryRowNoCache(&resp, queryString, fieldname, email)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, model.ErrNotFound
	default:
		return resp, err
	}
}

func (m *defaultReadersModel) FindOneByUserName(username string) (Readers, error) {
	//readerIdKey := fmt.Sprintf("%s%v", cacheReadersIdPrefix, username)
	var resp Readers
	//err := m.QueryRow(&resp, readerIdKey, func(conn sqlx.SqlConn, v interface{}) error {
	//	query := fmt.Sprintf("select %s from %s where %s and `username` = ? limit 1", readersRows, m.table, softDeleteFlag)
	//	return conn.QueryRow(v, query, username)
	//})
	queryString := fmt.Sprintf("select %s from %s where %s and `username` = ? limit 1", readersRows, m.table, softDeleteFlag)
	err := m.CachedConn.QueryRowNoCache(&resp, queryString, username)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, model.ErrNotFound
	default:
		return resp, err
	}
}
func (m *defaultReadersModel) FindOneByUserEmail(email string) (Readers, error) {
	var resp Readers
	queryString := fmt.Sprintf("select %s from %s where %s and `email` = ? limit 1", readersRows, m.table, softDeleteFlag)
	err := m.CachedConn.QueryRowNoCache(&resp, queryString, email)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, model.ErrNotFound
	default:
		return resp, err
	}
}
