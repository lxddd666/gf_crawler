// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgstatChannelDao is the data access object for table hg_tgstat_channel.
type TgstatChannelDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns TgstatChannelColumns // columns contains all the column names of Table for convenient usage.
}

// TgstatChannelColumns defines and stores column names for table hg_tgstat_channel.
type TgstatChannelColumns struct {
	Id            string // id
	Title         string // 标题
	Subscribers   string // 频道人数
	PostReach     string // 点赞数
	CitationIndex string // 索引数
	Type          string // 类型
	Avatar        string // 头像url
	TelegramLink  string // telegram地址
	CreatedAt     string // 创建时间
	UpdatedAt     string // 修改时间
	DeletedAt     string //
	Private       string // 是否私有群
	Comment       string // 评论
}

// tgstatChannelColumns holds the columns for table hg_tgstat_channel.
var tgstatChannelColumns = TgstatChannelColumns{
	Id:            "id",
	Title:         "title",
	Subscribers:   "subscribers",
	PostReach:     "post_reach",
	CitationIndex: "citation_index",
	Type:          "type",
	Avatar:        "avatar",
	TelegramLink:  "telegram_link",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	Private:       "private",
	Comment:       "comment",
}

// NewTgstatChannelDao creates and returns a new DAO object for table data access.
func NewTgstatChannelDao() *TgstatChannelDao {
	return &TgstatChannelDao{
		group:   "default",
		table:   "hg_tgstat_channel",
		columns: tgstatChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgstatChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgstatChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgstatChannelDao) Columns() TgstatChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgstatChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgstatChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgstatChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
