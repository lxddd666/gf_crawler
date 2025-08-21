// Package crawler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package crawler

import (
	"context"
	"hotgo/api/admin/tgstatchannel"
	"hotgo/internal/model/input/crawlerin"
	"hotgo/internal/service"
)

var (
	TgstatChannel = cTgstatChannel{}
)

type cTgstatChannel struct{}

// List 查看tgstat频道列表
func (c *cTgstatChannel) List(ctx context.Context, req *tgstatchannel.ListReq) (res *tgstatchannel.ListRes, err error) {
	list, totalCount, err := service.CrawlerTgstatChannel().List(ctx, &req.TgstatChannelListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*crawlerin.TgstatChannelListModel{}
	}

	res = new(tgstatchannel.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出tgstat频道列表
func (c *cTgstatChannel) Export(ctx context.Context, req *tgstatchannel.ExportReq) (res *tgstatchannel.ExportRes, err error) {
	err = service.CrawlerTgstatChannel().Export(ctx, &req.TgstatChannelListInp)
	return
}

// Edit 更新tgstat频道
func (c *cTgstatChannel) Edit(ctx context.Context, req *tgstatchannel.EditReq) (res *tgstatchannel.EditRes, err error) {
	err = service.CrawlerTgstatChannel().Edit(ctx, &req.TgstatChannelEditInp)
	return
}

// View 获取指定tgstat频道信息
func (c *cTgstatChannel) View(ctx context.Context, req *tgstatchannel.ViewReq) (res *tgstatchannel.ViewRes, err error) {
	data, err := service.CrawlerTgstatChannel().View(ctx, &req.TgstatChannelViewInp)
	if err != nil {
		return
	}

	res = new(tgstatchannel.ViewRes)
	res.TgstatChannelViewModel = data
	return
}

// Delete 删除tgstat频道
func (c *cTgstatChannel) Delete(ctx context.Context, req *tgstatchannel.DeleteReq) (res *tgstatchannel.DeleteRes, err error) {
	err = service.CrawlerTgstatChannel().Delete(ctx, &req.TgstatChannelDeleteInp)
	return
}

func (c *cTgstatChannel) CrawlerChannelUrl(ctx context.Context, req *tgstatchannel.CrawlerChannelUrlReq) (res *tgstatchannel.CrawlerChannelUrlRes, err error) {
	err = service.CrawlerTgstatChannel().CrawlerChannelUrl(ctx, &req.TgstatChannelCrawlerChannelInp)
	return
}

func (c *cTgstatChannel) StartRating(ctx context.Context, req *tgstatchannel.StartRatingReq) (res *tgstatchannel.StartRatingRes, err error) {
	err = service.CrawlerTgstatChannel().StartRating(ctx)
	return
}
