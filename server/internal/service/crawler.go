// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/crawlerin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ICrawlerTgstatChannel interface {
		// Model tgstat频道ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取tgstat频道列表
		List(ctx context.Context, in *crawlerin.TgstatChannelListInp) (list []*crawlerin.TgstatChannelListModel, totalCount int, err error)
		// Export 导出tgstat频道
		Export(ctx context.Context, in *crawlerin.TgstatChannelListInp) (err error)
		// Edit 修改/新增tgstat频道
		Edit(ctx context.Context, in *crawlerin.TgstatChannelEditInp) (err error)
		// Delete 删除tgstat频道
		Delete(ctx context.Context, in *crawlerin.TgstatChannelDeleteInp) (err error)
		// View 获取tgstat频道指定信息
		View(ctx context.Context, in *crawlerin.TgstatChannelViewInp) (res *crawlerin.TgstatChannelViewModel, err error)
		// CrawlerChannelUrl 抓取tgstat频道
		CrawlerChannelUrl(ctx context.Context, in *crawlerin.TgstatChannelCrawlerChannelInp) (err error)
		// StartRating 开始爬虫
		StartRating(ctx context.Context) error
	}
)

var (
	localCrawlerTgstatChannel ICrawlerTgstatChannel
)

func CrawlerTgstatChannel() ICrawlerTgstatChannel {
	if localCrawlerTgstatChannel == nil {
		panic("implement not found for interface ICrawlerTgstatChannel, forgot register?")
	}
	return localCrawlerTgstatChannel
}

func RegisterCrawlerTgstatChannel(i ICrawlerTgstatChannel) {
	localCrawlerTgstatChannel = i
}
