// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgstatChannelCrawlerUrlDao is the data access object for table hg_tgstat_channel_crawler_url.
type TgstatChannelCrawlerUrlDao struct {
	table   string                         // table is the underlying table name of the DAO.
	group   string                         // group is the database configuration group name of current DAO.
	columns TgstatChannelCrawlerUrlColumns // columns contains all the column names of Table for convenient usage.
}

// TgstatChannelCrawlerUrlColumns defines and stores column names for table hg_tgstat_channel_crawler_url.
type TgstatChannelCrawlerUrlColumns struct {
	Id        string // id
	Url       string // url
	Status    string // 采集状态0未开始1执行完成-1执行失败
	Type      string // 采集类型
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string //
	Comment   string //
}

// tgstatChannelCrawlerUrlColumns holds the columns for table hg_tgstat_channel_crawler_url.
var tgstatChannelCrawlerUrlColumns = TgstatChannelCrawlerUrlColumns{
	Id:        "id",
	Url:       "url",
	Status:    "status",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Comment:   "comment",
}

// NewTgstatChannelCrawlerUrlDao creates and returns a new DAO object for table data access.
func NewTgstatChannelCrawlerUrlDao() *TgstatChannelCrawlerUrlDao {
	return &TgstatChannelCrawlerUrlDao{
		group:   "default",
		table:   "hg_tgstat_channel_crawler_url",
		columns: tgstatChannelCrawlerUrlColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgstatChannelCrawlerUrlDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgstatChannelCrawlerUrlDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgstatChannelCrawlerUrlDao) Columns() TgstatChannelCrawlerUrlColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgstatChannelCrawlerUrlDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgstatChannelCrawlerUrlDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgstatChannelCrawlerUrlDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
