// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgstatChannelCrawlerUrl is the golang structure of table hg_tgstat_channel_crawler_url for DAO operations like Where/Data.
type TgstatChannelCrawlerUrl struct {
	g.Meta    `orm:"table:hg_tgstat_channel_crawler_url, do:true"`
	Id        interface{} // id
	Url       interface{} // url
	Status    interface{} // 采集状态0未开始1执行完成-1执行失败
	Type      interface{} // 采集类型
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time //
	Comment   interface{} //
}
