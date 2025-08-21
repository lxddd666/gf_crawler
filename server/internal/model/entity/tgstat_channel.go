// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgstatChannel is the golang structure for table tgstat_channel.
type TgstatChannel struct {
	Id            int64       `json:"id"            description:"id"`
	Title         string      `json:"title"         description:"标题"`
	Subscribers   int64       `json:"subscribers"   description:"频道人数"`
	PostReach     string      `json:"postReach"     description:"点赞数"`
	CitationIndex float64     `json:"citationIndex" description:"索引数"`
	Type          string      `json:"type"          description:"类型"`
	Avatar        string      `json:"avatar"        description:"头像url"`
	TelegramLink  string      `json:"telegramLink"  description:"telegram地址"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"修改时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:""`
	Private       int         `json:"private"       description:"是否私有群"`
	Comment       string      `json:"comment"       description:"评论"`
}
