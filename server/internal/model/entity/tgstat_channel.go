// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgstatChannel is the golang structure for table tgstat_channel.
type TgstatChannel struct {
	Id            int64       `json:"id"            orm:"id"             description:"id"`
	Title         string      `json:"title"         orm:"title"          description:"标题"`
	Subscribers   int64       `json:"subscribers"   orm:"subscribers"    description:"频道人数"`
	PostReach     string      `json:"postReach"     orm:"post_reach"     description:"点赞数"`
	CitationIndex float64     `json:"citationIndex" orm:"citation_index" description:"索引数"`
	Type          string      `json:"type"          orm:"type"           description:"类型"`
	Avatar        string      `json:"avatar"        orm:"avatar"         description:"头像url"`
	TelegramLink  string      `json:"telegramLink"  orm:"telegram_link"  description:"telegram地址"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"修改时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:""`
}
