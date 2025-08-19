// Package crawlerin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package crawlerin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgstatChannelUpdateFields 修改tgstat频道字段过滤
type TgstatChannelUpdateFields struct {
	Title         string  `json:"title"         dc:"标题"`
	Subscribers   int64   `json:"subscribers"   dc:"频道人数"`
	PostReach     string  `json:"postReach"     dc:"点赞数"`
	CitationIndex float64 `json:"citationIndex" dc:"索引数"`
	Type          string  `json:"type"          dc:"类型"`
	Avatar        string  `json:"avatar"        dc:"头像url"`
	TelegramLink  string  `json:"telegramLink"  dc:"telegram地址"`
}

// TgstatChannelInsertFields 新增tgstat频道字段过滤
type TgstatChannelInsertFields struct {
	Title         string  `json:"title"         dc:"标题"`
	Subscribers   int64   `json:"subscribers"   dc:"频道人数"`
	PostReach     string  `json:"postReach"     dc:"点赞数"`
	CitationIndex float64 `json:"citationIndex" dc:"索引数"`
	Type          string  `json:"type"          dc:"类型"`
	Avatar        string  `json:"avatar"        dc:"头像url"`
	TelegramLink  string  `json:"telegramLink"  dc:"telegram地址"`
}

// TgstatChannelEditInp 修改/新增tgstat频道
type TgstatChannelEditInp struct {
	entity.TgstatChannel
}

func (in *TgstatChannelEditInp) Filter(ctx context.Context) (err error) {
	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type TgstatChannelEditModel struct{}

// TgstatChannelDeleteInp 删除tgstat频道
type TgstatChannelDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgstatChannelDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgstatChannelDeleteModel struct{}

// TgstatChannelViewInp 获取指定tgstat频道信息
type TgstatChannelViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgstatChannelViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgstatChannelViewModel struct {
	entity.TgstatChannel
}

// TgstatChannelListInp 获取tgstat频道列表
type TgstatChannelListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"id"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TgstatChannelListInp) Filter(ctx context.Context) (err error) {
	return
}

// TgstatChannelCrawlerChannelInp 爬虫tgstat频道
type TgstatChannelCrawlerChannelInp struct {
}

func (in *TgstatChannelCrawlerChannelInp) Filter(ctx context.Context) (err error) {
	return
}

type TgstatChannelListModel struct {
	Id            int64       `json:"id"            dc:"id"`
	Title         string      `json:"title"         dc:"标题"`
	Subscribers   int64       `json:"subscribers"   dc:"频道人数"`
	PostReach     string      `json:"postReach"     dc:"点赞数"`
	CitationIndex float64     `json:"citationIndex" dc:"索引数"`
	Type          string      `json:"type"          dc:"类型"`
	Avatar        string      `json:"avatar"        dc:"头像url"`
	TelegramLink  string      `json:"telegramLink"  dc:"telegram地址"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"修改时间"`
}

// TgstatChannelExportModel 导出tgstat频道
type TgstatChannelExportModel struct {
	Id            int64       `json:"id"            dc:"id"`
	Title         string      `json:"title"         dc:"标题"`
	Subscribers   int64       `json:"subscribers"   dc:"频道人数"`
	PostReach     string      `json:"postReach"     dc:"点赞数"`
	CitationIndex float64     `json:"citationIndex" dc:"索引数"`
	Type          string      `json:"type"          dc:"类型"`
	Avatar        string      `json:"avatar"        dc:"头像url"`
	TelegramLink  string      `json:"telegramLink"  dc:"telegram地址"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"修改时间"`
}
