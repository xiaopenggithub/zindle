package model

import (
	"backend/common/utils"
	"backend/service/backend/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
)

var (
	bookOrdersRows                = "`id`,`member_id`,`book_id`,`return_date`,`created_at`,`updated_at`,`deleted_at`"
	bookOrdersRowsExpectAutoSet   = "`member_id`,`book_id`,`return_date`"
	bookOrdersRowsWithPlaceHolder = "`member_id`=?,`book_id`=?,`return_date`=?"

	cacheBookOrdersIdPrefix = "cache#bookOrders#id#"
)

type (
	BookOrdersModel interface {
		Insert(data BookOrders) (sql.Result, error)
		FindOne(id int64) (*BookOrders, error)
		FindOneOrder(id int64, readerId int64) (*BookOrders, error)
		CheckBorrow(id int64, readerId int64) (*BookOrders, error)
		Update(data BookOrders) error
		Delete(id int64) error
		List(req utils.ListReq) ([]*BookOrders, int, error)
		ListBorrow(req utils.ListReq, readerId int64) ([]*BookOrders, int, error)
		DeleteBatch(ids string) error
		CheckDuplicate(fieldname string) (BookOrders, error)
		ReturnBook(data BookOrders) error
		BorrowBook(bookId int64, readerId int64) error
		NearReturnBook(nearTime string) ([]*BookOrders, error)
	}

	defaultBookOrdersModel struct {
		sqlc.CachedConn
		table string
	}

	BookOrders struct {
		Id         int64  `db:"id" json:"id"`                   // ID
		MemberId   int64  `db:"member_id" json:"member_id"`     // 用户ID
		BookId     int64  `db:"book_id" json:"book_id"`         // 书ID
		ReturnDate string `db:"return_date" json:"return_date"` // 还书日期
		CreatedAt  string `db:"created_at" json:"created_at"`   // 创建时间(借书日期)
		UpdatedAt  string `db:"updated_at" json:"updated_at"`   // 更新时间
		DeletedAt  string `db:"deleted_at" json:"deleted_at"`   // 删除时间

	}
)

// 图书订单
func NewBookOrdersModel(conn sqlx.SqlConn, c cache.CacheConf) BookOrdersModel {
	return &defaultBookOrdersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`book_orders`",
	}
}

func (m *defaultBookOrdersModel) Insert(data BookOrders) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?,?,?)", m.table, bookOrdersRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.MemberId, data.BookId, data.ReturnDate)

	return ret, err
}

func (m *defaultBookOrdersModel) FindOne(id int64) (*BookOrders, error) {
	bookOrdersIdKey := fmt.Sprintf("%s%v", cacheBookOrdersIdPrefix, id)
	var resp BookOrders
	err := m.QueryRow(&resp, bookOrdersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookOrdersRows, m.table)
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
func (m *defaultBookOrdersModel) FindOneOrder(id int64, readerId int64) (*BookOrders, error) {
	bookOrdersIdKey := fmt.Sprintf("%s%v", cacheBookOrdersIdPrefix, id)
	var resp BookOrders
	err := m.QueryRow(&resp, bookOrdersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where (`id` = ? and `member_id` = ?) limit 1", bookOrdersRows, m.table)
		return conn.QueryRow(v, query, id, readerId)
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
func (m *defaultBookOrdersModel) CheckBorrow(id int64, readerId int64) (*BookOrders, error) {
	//bookOrdersIdKey := fmt.Sprintf("%s%v", cacheBookOrdersIdPrefix, id)
	var resp BookOrders
	//err := m.QueryRow(&resp, bookOrdersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
	//	query := fmt.Sprintf("select %s from %s where (`book_id` = ? and `member_id` = ? and `return_date` = '2006-01-02 15:04:05') limit 1", bookOrdersRows, m.table)
	//	return conn.QueryRow(v, query, id, readerId)
	//})
	queryString := fmt.Sprintf("select %s from %s where (`book_id` = ? and `member_id` = ? and `return_date` = '2006-01-02 15:04:05') limit 1", bookOrdersRows, m.table)
	err := m.CachedConn.QueryRowNoCache(&resp, queryString, id, readerId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookOrdersModel) Update(data BookOrders) error {
	bookOrdersIdKey := fmt.Sprintf("%s%v", cacheBookOrdersIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bookOrdersRowsWithPlaceHolder)
		return conn.Exec(query, data.MemberId, data.BookId, data.ReturnDate, data.Id)
	}, bookOrdersIdKey)
	return err
}

func (m *defaultBookOrdersModel) Delete(id int64) error {
	bookOrdersIdKey := fmt.Sprintf("%s%v", cacheBookOrdersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `deleted_at`=? where `id` = ?", m.table)
		return conn.Exec(query, time.Now().Format("2006-01-02 15:04:05"), id)
	}, bookOrdersIdKey)
	return err
}

func (m *defaultBookOrdersModel) List(req utils.ListReq) ([]*BookOrders, int, error) {
	total := 0

	// 条件处理
	whereCondition := "where " + softDeleteFlag
	if req.Keyword != "" {
		whereCondition += "and `name` like '%" + req.Keyword + "%'"
	}

	orderBy := "order by id desc"
	items := make([]*BookOrders, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT ? OFFSET ?", bookOrdersRows, m.table, whereCondition, orderBy)
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
func (m *defaultBookOrdersModel) ListBorrow(req utils.ListReq, readerId int64) ([]*BookOrders, int, error) {
	total := 0

	// 条件处理
	whereCondition := "where " + softDeleteFlag
	whereCondition += fmt.Sprintf(" and `member_id`=%v ", readerId)
	orderBy := "order by id desc"
	items := make([]*BookOrders, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT ? OFFSET ?", bookOrdersRows, m.table, whereCondition, orderBy)
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
func (m *defaultBookOrdersModel) DeleteBatch(ids string) error {
	query := fmt.Sprintf("update %s set `deleted_at`=? where `id` in (%s)", m.table, ids)
	_, err := m.ExecNoCache(query, time.Now().Format("2006-01-02 15:04:05"))

	// 删除缓存
	idArr := strings.Split(ids, ",")
	for _, v := range idArr {
		systemUserIdKey := fmt.Sprintf("%s%v", cacheBookOrdersIdPrefix, v)
		m.DelCache(systemUserIdKey)
	}
	return err
}
func (m *defaultBookOrdersModel) CheckDuplicate(fieldname string) (BookOrders, error) {
	var resp BookOrders
	queryString := fmt.Sprintf("select %s from %s where %s and `name` = ? limit 1", bookOrdersRows, m.table, softDeleteFlag)
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

func (m *defaultBookOrdersModel) NearReturnBook(nearTime string) ([]*BookOrders, error) {
	items := make([]*BookOrders, 0)
	whereCondition := "where " + softDeleteFlag
	whereCondition += " and `created_at`< ? "
	orderBy := "order by id desc"
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", bookOrdersRows, m.table, whereCondition, orderBy)
	err := m.CachedConn.QueryRowsNoCache(&items, query, nearTime)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return items, nil
}
func (m *defaultBookOrdersModel) ReturnBook(data BookOrders) error {
	bookOrdersIdKey := fmt.Sprintf("%s%v", cacheBookOrdersIdPrefix, data.Id)
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, data.BookId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		err0 := conn.Transact(func(session sqlx.Session) error {
			// 1.归还
			query := fmt.Sprintf("update %s set `return_date`=? where (`id` = ? and `return_date`='2006-01-02 15:04:05')", m.table)
			e, err1 := conn.Exec(query, time.Now().Format("2006-01-02 15:04:05"), data.Id)
			if err1 != nil {
				return err1
			}
			returnAffected, _ := e.RowsAffected()

			// 2.恢复库存
			if returnAffected > 0 {
				query := fmt.Sprintf("update %s set `quantity`=`quantity`+1 where id=?", "`books`")
				_, err2 := conn.Exec(query, data.BookId)
				if err2 != nil {
					return err2
				}
			} else {
				return errors.New("归还失败")
			}
			return nil
		})

		return nil, err0
	}, bookOrdersIdKey, booksIdKey)

	return err
}

func (m *defaultBookOrdersModel) BorrowBook(bookId int64, readerId int64) error {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, bookId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		err0 := conn.Transact(func(session sqlx.Session) error {
			// 1.生成订单
			query := fmt.Sprintf("insert into %s (%s) values (?,?,?)", m.table, bookOrdersRowsExpectAutoSet)
			ret, err0 := m.ExecNoCache(query, readerId, bookId, "2006-01-02 15:04:05")
			affected, _ := ret.RowsAffected()
			id, _ := ret.LastInsertId()
			fmt.Println(ret, err0, affected, id)
			// 2.减扣库存 todo 检查负值情况，为0后不再借
			if err0 != nil {
				return errors.New("借书失败")
			} else {
				query := fmt.Sprintf("update %s set `quantity`=`quantity`-1 where id=?", "`books`")
				_, err2 := conn.Exec(query, bookId)
				if err2 != nil {
					return err2
				}
			}
			return nil
		})

		return nil, err0
	}, booksIdKey)

	return err
}
