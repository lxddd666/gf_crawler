// Package crawler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package crawler

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/crawlerin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCrawlerTgstatChannel struct{}

func NewCrawlerTgstatChannel() *sCrawlerTgstatChannel {
	return &sCrawlerTgstatChannel{}
}

func init() {
	service.RegisterCrawlerTgstatChannel(NewCrawlerTgstatChannel())
}

// Model tgstat频道ORM模型
func (s *sCrawlerTgstatChannel) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgstatChannel.Ctx(ctx), option...)
}

// List 获取tgstat频道列表
func (s *sCrawlerTgstatChannel) List(ctx context.Context, in *crawlerin.TgstatChannelListInp) (list []*crawlerin.TgstatChannelListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(crawlerin.TgstatChannelListModel{})

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.TgstatChannel.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgstatChannel.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.TgstatChannel.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取tgstat频道列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出tgstat频道
func (s *sCrawlerTgstatChannel) Export(ctx context.Context, in *crawlerin.TgstatChannelListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(crawlerin.TgstatChannelExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出tgstat频道-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []crawlerin.TgstatChannelExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增tgstat频道
func (s *sCrawlerTgstatChannel) Edit(ctx context.Context, in *crawlerin.TgstatChannelEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(crawlerin.TgstatChannelUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改tgstat频道失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(crawlerin.TgstatChannelInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增tgstat频道失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除tgstat频道
func (s *sCrawlerTgstatChannel) Delete(ctx context.Context, in *crawlerin.TgstatChannelDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除tgstat频道失败，请稍后重试！")
		return
	}
	return
}

// View 获取tgstat频道指定信息
func (s *sCrawlerTgstatChannel) View(ctx context.Context, in *crawlerin.TgstatChannelViewInp) (res *crawlerin.TgstatChannelViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取tgstat频道信息，请稍后重试！")
		return
	}
	return
}

// CrawlerChannel 抓取tgstat频道
func (s *sCrawlerTgstatChannel) CrawlerChannel(ctx context.Context, in *crawlerin.TgstatChannelCrawlerChannelInp) (err error) {
	// 创建默认收集器
	c := colly.NewCollector()
	// 向 API 发送请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("发送请求到：", r.URL)
	})
	// 处理 API 返回的数据
	c.OnResponse(func(r *colly.Response) {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(gconv.String(r.Body)))
		if err != nil {
			return
		}
		doc.Find("a[target='_blank']").Each(func(i int, s *goquery.Selection) {
			results := map[string]string{
				"链接地址": s.AttrOr("href", ""),
				"图片地址": s.Find("img").AttrOr("src", ""),
				"标题":   s.Find(".font-16.text-dark").Text(),
				"订阅人数": extractSubscribers(s.Find(".font-14.text-dark").Text()),
				"分类标签": s.Find(".bg-light.px-1").Text(),
				//"完整HTML": getOuterHTML(a),
			}
			for k, v := range results {
				fmt.Println(k, ":", v)
			}
		})

	})
	c.SetRequestTimeout(20 * time.Second)
	// 访问起始页面
	c.Visit("https://tgstat.com/ratings/channels?sort=members")
	return
}

func extractSubscribers(text string) string {
	if idx := strings.Index(text, "subscribers"); idx != -1 {
		return strings.TrimSpace(text[:idx])
	}
	return text
}
