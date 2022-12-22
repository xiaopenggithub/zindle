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
	activityOrderssRows                = "`id`,`member_id`,`activity_id`,`status`,`seat_number`,`created_at`,`updated_at`,`deleted_at`"
	activityOrderssRowsExpectAutoSet   = "`member_id`,`activity_id`,`status`,`seat_number`"
	activityOrderssRowsWithPlaceHolder = "`member_id`=?,`activity_id`=?,`status`=?,`seat_number`=?"

	cacheActivityOrderssIdPrefix = "cache#activityOrderss#id#"
)

type (
	ActivityOrderssModel interface {
		Insert(data ActivityOrderss) (sql.Result, error)
		FindOne(id int64) (*ActivityOrderss, error)
		FindOneOrder(id int64, memberId int64) (*ActivityOrderss, error)
		Update(data ActivityOrderss) error
		Delete(id int64) error
		List(req utils.ListReq) ([]*ActivityOrderss, int, error)
		ListMy(req utils.ListReq, memberId int64) ([]*ActivityOrderss, int, error)
		DeleteBatch(ids string) error
		CheckDuplicate(fieldname string) (ActivityOrderss, error)
	}

	defaultActivityOrderssModel struct {
		sqlc.CachedConn
		table string
	}

	ActivityOrderss struct {
		Id         int64  `db:"id" json:"id"`                   // ID
		MemberId   int64  `db:"member_id" json:"member_id"`     // 用户ID
		ActivityId int64  `db:"activity_id" json:"activity_id"` // 活动ID
		Status     int64  `db:"status" json:"status"`           // 0报名1签到2取消3超时失约
		SeatNumber int64  `db:"seat_number" json:"seat_number"` // 座位
		CreatedAt  string `db:"created_at" json:"created_at"`   // 创建时间(借书日期)
		UpdatedAt  string `db:"updated_at" json:"updated_at"`   // 更新时间
		DeletedAt  string `db:"deleted_at" json:"deleted_at"`   // 删除时间

	}
)

// 活动预订信息
func NewActivityOrderssModel(conn sqlx.SqlConn, c cache.CacheConf) ActivityOrderssModel {
	return &defaultActivityOrderssModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`activity_orders`",
	}
}

func (m *defaultActivityOrderssModel) Insert(data ActivityOrderss) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?,?,?,?)", m.table, activityOrderssRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.MemberId, data.ActivityId, data.Status, data.SeatNumber)

	return ret, err
}

func (m *defaultActivityOrderssModel) FindOne(id int64) (*ActivityOrderss, error) {
	activityOrderssIdKey := fmt.Sprintf("%s%v", cacheActivityOrderssIdPrefix, id)
	var resp ActivityOrderss
	err := m.QueryRow(&resp, activityOrderssIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", activityOrderssRows, m.table)
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

func (m *defaultActivityOrderssModel) FindOneOrder(id int64, memberId int64) (*ActivityOrderss, error) {
	activityOrderssIdKey := fmt.Sprintf("%s%v", cacheActivityOrderssIdPrefix, id)
	var resp ActivityOrderss
	err := m.QueryRow(&resp, activityOrderssIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where (`id` = ? and `member_id` = ?) limit 1", activityOrderssRows, m.table)
		return conn.QueryRow(v, query, id, memberId)
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

func (m *defaultActivityOrderssModel) Update(data ActivityOrderss) error {
	activityOrderssIdKey := fmt.Sprintf("%s%v", cacheActivityOrderssIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, activityOrderssRowsWithPlaceHolder)
		return conn.Exec(query, data.MemberId, data.ActivityId, data.Status, data.SeatNumber, data.Id)
	}, activityOrderssIdKey)
	return err
}

func (m *defaultActivityOrderssModel) Delete(id int64) error {
	activityOrderssIdKey := fmt.Sprintf("%s%v", cacheActivityOrderssIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `deleted_at`=? where `id` = ?", m.table)
		return conn.Exec(query, time.Now().Format("2006-01-02 15:04:05"), id)
	}, activityOrderssIdKey)
	return err
}

func (m *defaultActivityOrderssModel) List(req utils.ListReq) ([]*ActivityOrderss, int, error) {
	total := 0

	// 条件处理
	whereCondition := "where " + softDeleteFlag
	//if req.Keyword != "" {
	//	whereCondition += "and `name` like '%" + req.Keyword + "%'"
	//}

	orderBy := "order by id desc"
	items := make([]*ActivityOrderss, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT ? OFFSET ?", activityOrderssRows, m.table, whereCondition, orderBy)
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

func (m *defaultActivityOrderssModel) ListMy(req utils.ListReq, memberId int64) ([]*ActivityOrderss, int, error) {
	total := 0

	// 条件处理
	whereCondition := "where " + softDeleteFlag
	whereCondition += fmt.Sprintf(" and `member_id`=%v ", memberId)
	//if req.Keyword != "" {
	//	whereCondition += "and `name` like '%" + req.Keyword + "%'"
	//}

	orderBy := "order by id desc"
	items := make([]*ActivityOrderss, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT ? OFFSET ?", activityOrderssRows, m.table, whereCondition, orderBy)
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

func (m *defaultActivityOrderssModel) DeleteBatch(ids string) error {
	query := fmt.Sprintf("update %s set `deleted_at`=? where `id` in (%s)", m.table, ids)
	_, err := m.ExecNoCache(query, time.Now().Format("2006-01-02 15:04:05"))

	// 删除缓存
	idArr := strings.Split(ids, ",")
	for _, v := range idArr {
		systemUserIdKey := fmt.Sprintf("%s%v", cacheActivityOrderssIdPrefix, v)
		m.DelCache(systemUserIdKey)
	}
	return err
}
func (m *defaultActivityOrderssModel) CheckDuplicate(fieldname string) (ActivityOrderss, error) {
	var resp ActivityOrderss
	queryString := fmt.Sprintf("select %s from %s where %s and `name` = ? limit 1", activityOrderssRows, m.table, softDeleteFlag)
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
