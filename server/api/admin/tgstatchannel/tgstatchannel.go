// Package tgstatchannel
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package tgstatchannel

import (
	"hotgo/internal/model/input/crawlerin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询tgstat频道列表
type ListReq struct {
	g.Meta `path:"/tgstatChannel/list" method:"get" tags:"tgstat频道" summary:"获取tgstat频道列表"`
	crawlerin.TgstatChannelListInp
}

type ListRes struct {
	form.PageRes
	List []*crawlerin.TgstatChannelListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出tgstat频道列表
type ExportReq struct {
	g.Meta `path:"/tgstatChannel/export" method:"get" tags:"tgstat频道" summary:"导出tgstat频道列表"`
	crawlerin.TgstatChannelListInp
}

type ExportRes struct{}

// ViewReq 获取tgstat频道指定信息
type ViewReq struct {
	g.Meta `path:"/tgstatChannel/view" method:"get" tags:"tgstat频道" summary:"获取tgstat频道指定信息"`
	crawlerin.TgstatChannelViewInp
}

type ViewRes struct {
	*crawlerin.TgstatChannelViewModel
}

// EditReq 修改/新增tgstat频道
type EditReq struct {
	g.Meta `path:"/tgstatChannel/edit" method:"post" tags:"tgstat频道" summary:"修改/新增tgstat频道"`
	crawlerin.TgstatChannelEditInp
}

type EditRes struct{}

// DeleteReq 删除tgstat频道
type DeleteReq struct {
	g.Meta `path:"/tgstatChannel/delete" method:"post" tags:"tgstat频道" summary:"删除tgstat频道"`
	crawlerin.TgstatChannelDeleteInp
}

type DeleteRes struct{}

type CrawlerChannelReq struct {
	g.Meta `path:"/tgstatChannel/crawlerChannel" method:"post" tags:"tgstat频道" summary:"爬虫tgstat频道"`
	crawlerin.TgstatChannelCrawlerChannelInp
}

type CrawlerChannelRes struct {
}
