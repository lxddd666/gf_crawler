// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgstatChannelCrawlerUrl is the golang structure for table tgstat_channel_crawler_url.
type TgstatChannelCrawlerUrl struct {
	Id        int64       `json:"id"        description:"id"`
	Url       string      `json:"url"       description:"url"`
	Status    int         `json:"status"    description:"采集状态0未开始1执行完成-1执行失败"`
	Type      string      `json:"type"      description:"采集类型"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
	Comment   string      `json:"comment"   description:""`
}
