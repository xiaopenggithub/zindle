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
	activitysRows                = "`address`,`cover`,`create_by`,`created_at`,`deleted_at`,`description`,`id`,`quantity`,`sort`,`time_end`,`time_start`,`title`,`update_by`,`updated_at`"
	activitysRowsExpectAutoSet   = "`address`,`cover`,`create_by`,`description`,`quantity`,`sort`,`time_end`,`time_start`,`title`,`update_by`"
	activitysRowsWithPlaceHolder = "`address`=?,`cover`=?,`create_by`=?,`description`=?,`quantity`=?,`sort`=?,`time_end`=?,`time_start`=?,`title`=?,`update_by`=?"

	cacheActivitysIdPrefix = "cache#activitys#id#"
	softDeleteFlag         = "`deleted_at`='2006-01-02 15:04:05'"
)

type (
	ActivitysModel interface {
		Insert(data Activitys) (sql.Result, error)
		FindOne(id int64) (*Activitys, error)
		Update(data Activitys) error
		Delete(id int64) error
		List(req utils.ListReq) ([]*Activitys, int, error)
		DeleteBatch(ids string) error
		CheckDuplicate(fieldname string) (Activitys, error)
	}

	defaultActivitysModel struct {
		sqlc.CachedConn
		table string
	}

	Activitys struct {
		Address     string `db:"address" json:"address"`         // 活动地址
		Cover       string `db:"cover" json:"cover"`             // 封面
		CreateBy    string `db:"create_by" json:"create_by"`     // 创建者
		CreatedAt   string `db:"created_at" json:"created_at"`   // 创建时间
		DeletedAt   string `db:"deleted_at" json:"deleted_at"`   // 删除时间
		Description string `db:"description" json:"description"` // 简介
		Id          int64  `db:"id" json:"id"`                   // ID
		Quantity    int64  `db:"quantity" json:"quantity"`       // 数量
		Sort        int64  `db:"sort" json:"sort"`               // 排序
		TimeEnd     string `db:"time_end" json:"time_end"`       // 结束时间
		TimeStart   string `db:"time_start" json:"time_start"`   // 开始时间
		Title       string `db:"title" json:"title"`             // 标题
		UpdateBy    string `db:"update_by" json:"update_by"`     // 更新者
		UpdatedAt   string `db:"updated_at" json:"updated_at"`   // 更新时间

	}
)

// 付款信息
func NewActivitysModel(conn sqlx.SqlConn, c cache.CacheConf) ActivitysModel {
	return &defaultActivitysModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`activities`",
	}
}

func (m *defaultActivitysModel) Insert(data Activitys) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?,?,?,?,?,?,?,?,?,?)", m.table, activitysRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Address, data.Cover, data.CreateBy, data.Description, data.Quantity, data.Sort, data.TimeEnd, data.TimeStart, data.Title, data.UpdateBy)

	return ret, err
}

func (m *defaultActivitysModel) FindOne(id int64) (*Activitys, error) {
	activitysIdKey := fmt.Sprintf("%s%v", cacheActivitysIdPrefix, id)
	var resp Activitys
	err := m.QueryRow(&resp, activitysIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", activitysRows, m.table)
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

func (m *defaultActivitysModel) Update(data Activitys) error {
	activitysIdKey := fmt.Sprintf("%s%v", cacheActivitysIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, activitysRowsWithPlaceHolder)
		return conn.Exec(query, data.Address, data.Cover, data.CreateBy, data.Description, data.Quantity, data.Sort, data.TimeEnd, data.TimeStart, data.Title, data.UpdateBy, data.Id)
	}, activitysIdKey)
	return err
}

func (m *defaultActivitysModel) Delete(id int64) error {
	activitysIdKey := fmt.Sprintf("%s%v", cacheActivitysIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `deleted_at`=? where `id` = ?", m.table)
		return conn.Exec(query, time.Now().Format("2006-01-02 15:04:05"), id)
	}, activitysIdKey)
	return err
}

func (m *defaultActivitysModel) List(req utils.ListReq) ([]*Activitys, int, error) {
	total := 0

	// 条件处理
	whereCondition := "where " + softDeleteFlag
	if req.Keyword != "" {
		whereCondition += "and `title` like '%" + req.Keyword + "%'"
	}

	orderBy := "order by sort asc"
	items := make([]*Activitys, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT ? OFFSET ?", activitysRows, m.table, whereCondition, orderBy)
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
func (m *defaultActivitysModel) DeleteBatch(ids string) error {
	query := fmt.Sprintf("update %s set `deleted_at`=? where `id` in (%s)", m.table, ids)
	_, err := m.ExecNoCache(query, time.Now().Format("2006-01-02 15:04:05"))

	// 删除缓存
	idArr := strings.Split(ids, ",")
	for _, v := range idArr {
		systemUserIdKey := fmt.Sprintf("%s%v", cacheActivitysIdPrefix, v)
		m.DelCache(systemUserIdKey)
	}
	return err
}
func (m *defaultActivitysModel) CheckDuplicate(fieldname string) (Activitys, error) {
	var resp Activitys
	queryString := fmt.Sprintf("select %s from %s where %s and `name` = ? limit 1", activitysRows, m.table, softDeleteFlag)
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
