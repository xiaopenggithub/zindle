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
	booksRows                = "`id`,`title`,`description`,`quantity`,`cover`,`sort`,`create_by`,`update_by`,`created_at`,`updated_at`,`deleted_at`"
	booksRowsExpectAutoSet   = "`title`,`description`,`quantity`,`cover`,`sort`,`create_by`,`update_by`"
	booksRowsWithPlaceHolder = "`title`=?,`description`=?,`quantity`=?,`cover`=?,`sort`=?,`create_by`=?,`update_by`=?"

	cacheBooksIdPrefix = "cache#books#id#"
)

type (
	BooksModel interface {
		Insert(data Books) (sql.Result, error)
		FindOne(id int64) (*Books, error)
		Update(data Books) error
		Delete(id int64) error
		List(req utils.ListReq) ([]*Books, int, error)
		DeleteBatch(ids string) error
		CheckDuplicate(fieldname string) (Books, error)
		Counts(id int64) (int64, int64, int64, error)
	}

	defaultBooksModel struct {
		sqlc.CachedConn
		table string
	}

	Books struct {
		Id          int64  `db:"id" json:"id"`                   // ID
		Title       string `db:"title" json:"title"`             // 标题
		Description string `db:"description" json:"description"` // 简介
		Quantity    int64  `db:"quantity" json:"quantity"`       // 数量
		Cover       string `db:"cover" json:"cover"`             // 封面
		Sort        int64  `db:"sort" json:"sort"`               // 排序
		CreateBy    string `db:"create_by" json:"create_by"`     // 创建者
		UpdateBy    string `db:"update_by" json:"update_by"`     // 更新者
		CreatedAt   string `db:"created_at" json:"created_at"`   // 创建时间
		UpdatedAt   string `db:"updated_at" json:"updated_at"`   // 更新时间
		DeletedAt   string `db:"deleted_at" json:"deleted_at"`   // 删除时间

	}
)

// 图书
func NewBooksModel(conn sqlx.SqlConn, c cache.CacheConf) BooksModel {
	return &defaultBooksModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`books`",
	}
}

func (m *defaultBooksModel) Insert(data Books) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?,?,?,?,?,?,?)", m.table, booksRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Title, data.Description, data.Quantity, data.Cover, data.Sort, data.CreateBy, data.UpdateBy)

	return ret, err
}

func (m *defaultBooksModel) FindOne(id int64) (*Books, error) {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, id)
	var resp Books
	err := m.QueryRow(&resp, booksIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", booksRows, m.table)
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

func (m *defaultBooksModel) Update(data Books) error {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, booksRowsWithPlaceHolder)
		return conn.Exec(query, data.Title, data.Description, data.Quantity, data.Cover, data.Sort, data.CreateBy, data.UpdateBy, data.Id)
	}, booksIdKey)
	return err
}

func (m *defaultBooksModel) Delete(id int64) error {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `deleted_at`=? where `id` = ?", m.table)
		return conn.Exec(query, time.Now().Format("2006-01-02 15:04:05"), id)
	}, booksIdKey)
	return err
}

func (m *defaultBooksModel) List(req utils.ListReq) ([]*Books, int, error) {
	total := 0

	// 条件处理
	whereCondition := "where " + softDeleteFlag
	if req.Keyword != "" {
		whereCondition += "and `title` like '%" + req.Keyword + "%'"
	}

	orderBy := "order by sort asc"
	items := make([]*Books, 0)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT ? OFFSET ?", booksRows, m.table, whereCondition, orderBy)
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
func (m *defaultBooksModel) DeleteBatch(ids string) error {
	query := fmt.Sprintf("update %s set `deleted_at`=? where `id` in (%s)", m.table, ids)
	_, err := m.ExecNoCache(query, time.Now().Format("2006-01-02 15:04:05"))

	// 删除缓存
	idArr := strings.Split(ids, ",")
	for _, v := range idArr {
		systemUserIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, v)
		m.DelCache(systemUserIdKey)
	}
	return err
}
func (m *defaultBooksModel) CheckDuplicate(fieldname string) (Books, error) {
	var resp Books
	queryString := fmt.Sprintf("select %s from %s where %s and `name` = ? limit 1", booksRows, m.table, softDeleteFlag)
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

func (m *defaultBooksModel) Ta(fieldname string) (Books, error) {
	var resp Books

	err := m.CachedConn.Transact(func(session sqlx.Session) error {
		return nil
	})
	if err != nil {
		// ...
	}

	return resp, nil
}

func (m *defaultBooksModel) Counts(id int64) (totalBook int64, totalBorrow int64, totalShouldReturn int64, err error) {
	whereCondition := "where " + softDeleteFlag
	queryTotalBook := fmt.Sprintf("SELECT sum(`quantity`) FROM %s %s", m.table, whereCondition)
	err = m.CachedConn.QueryRowNoCache(&totalBook, queryTotalBook)
	if err != nil {
		totalBook = 0
	}

	queryTotalBorrow := fmt.Sprintf("SELECT count(1) FROM %s %s", "book_orders", whereCondition)
	err = m.CachedConn.QueryRowNoCache(&totalBorrow, queryTotalBorrow)
	if err != nil {
		totalBorrow = 0
	}

	queryTotalShouldReturn := fmt.Sprintf("SELECT count(1) FROM %s %s", "book_orders", whereCondition+" and `return_date`='2006-01-02 15:04:05' ")
	err = m.CachedConn.QueryRowNoCache(&totalShouldReturn, queryTotalShouldReturn)
	if err != nil {
		totalShouldReturn = 0
	}

	return totalBook, totalBorrow, totalShouldReturn, nil
}
