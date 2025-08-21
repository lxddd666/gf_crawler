// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgstatChannel is the golang structure of table hg_tgstat_channel for DAO operations like Where/Data.
type TgstatChannel struct {
	g.Meta        `orm:"table:hg_tgstat_channel, do:true"`
	Id            interface{} // id
	Title         interface{} // 标题
	Subscribers   interface{} // 频道人数
	PostReach     interface{} // 点赞数
	CitationIndex interface{} // 索引数
	Type          interface{} // 类型
	Avatar        interface{} // 头像url
	TelegramLink  interface{} // telegram地址
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 修改时间
	DeletedAt     *gtime.Time //
	Private       interface{} // 是否私有群
	Comment       interface{} // 评论
}
