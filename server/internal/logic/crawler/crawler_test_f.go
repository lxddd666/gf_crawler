package crawler

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"time"
)

func (s *sCrawlerTgstatChannel) TestCrawler(ctx context.Context) error {
	// 创建默认收集器
	c := colly.NewCollector()
	// 向 API 发送请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("发送请求到：", r.URL)
	})
	// 处理 API 返回的数据
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Body)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(gconv.String(r.Body)))
		if err != nil {
			return
		}
		fmt.Println(doc)

	})
	c.SetRequestTimeout(20 * time.Second)
	// 访问起始页面
	c.Visit("https://hk.investing.com/currencies/xau-usd")

	return nil
}
