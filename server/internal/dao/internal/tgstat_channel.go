// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgstatChannelDao is the data access object for the table hg_tgstat_channel.
type TgstatChannelDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  TgstatChannelColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// TgstatChannelColumns defines and stores column names for the table hg_tgstat_channel.
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
}

// tgstatChannelColumns holds the columns for the table hg_tgstat_channel.
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
}

// NewTgstatChannelDao creates and returns a new DAO object for table data access.
func NewTgstatChannelDao(handlers ...gdb.ModelHandler) *TgstatChannelDao {
	return &TgstatChannelDao{
		group:    "default",
		table:    "hg_tgstat_channel",
		columns:  tgstatChannelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TgstatChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TgstatChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TgstatChannelDao) Columns() TgstatChannelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TgstatChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TgstatChannelDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TgstatChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
