package crawler

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"log"
	"strings"
	"time"
)

// CrawlerChromedpChannel 抓取tgstat频道
func (s *sCrawlerTgstatChannel) CrawlerChromedpChannel(ctx context.Context, in *entity.TgstatChannelCrawlerUrl) (err error) {

	defer func() {
		if err != nil {
			dao.TgstatChannelCrawlerUrl.Ctx(ctx).Fields(dao.TgstatChannelCrawlerUrl.Columns().Status, dao.TgstatChannelCrawlerUrl.Columns().Comment).WherePri(in.Id).Data(do.TgstatChannelCrawlerUrl{Comment: err.Error(), Status: consts.TaskFail}).Update()
		} else {
			dao.TgstatChannelCrawlerUrl.Ctx(ctx).Fields(dao.TgstatChannelCrawlerUrl.Columns().Status).WherePri(in.Id).Data(do.TgstatChannelCrawlerUrl{Status: consts.TaskSuccess}).Update()
		}
	}()

	if in.Url == "" {
		return gerror.New("请输入正确的url")
	}

	list := make([]entity.TgstatChannel, 0)

	// 爬虫
	body, err := StartChromedp(in.Url)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(gconv.String(body)))
	if err != nil {
		return
	}
	doc.Find("div[class='card peer-item-row mb-2 ribbon-box border']").Each(func(i int, s *goquery.Selection) {
		collect := entity.TgstatChannel{}

		headSelect := s.Find("a[target='_blank']")
		collect.Title = headSelect.Find(".font-16.text-dark").Text()
		collect.Subscribers = extractSubscribers(headSelect.Find(".font-14.text-dark").Text())
		collect.Avatar = headSelect.Find("img").AttrOr("src", "")
		// 从完整链接中提取 @ 后面的部分
		collect.TelegramLink, collect.Private = extractTelegramLink(headSelect.AttrOr("href", ""))
		collect.Type = strings.TrimSpace(headSelect.Find("span[class='border rounded bg-light px-1']").Text())
		// 提取 PostReach (1 post reach)
		postReachText := s.Find(".col.col-4.pt-1").Eq(1).Find("h4.font-16.font-sm-18").Text()
		collect.PostReach = extractNumber(postReachText)

		// 提取 CitationIndex (citation index)
		citationText := s.Find(".col.col-4.pt-1").Eq(2).Find("h4.font-16.font-sm-18").Text()
		collect.CitationIndex = gconv.Float64(extractNumber(citationText))

		// 将收集的数据添加到列表中
		list = append(list, collect)
	})
	if len(list) == 0 {
		err = gerror.New("没有找到数据")
		return
	}
	// 保存收集的数据到数据库 - 分批插入，每次300个
	for _, cralwer := range list {
		_, _ = s.Model(ctx).FieldsEx(dao.TgstatChannel.Columns().Id).Insert(cralwer)
	}

	return
}

func StartChromedp(url string) (html string, err error) {
	soket5Proxy, flag := global.ProxyList.GetRandom()
	defer func() {
		if err != nil && gstr.Contains(err.Error(), "SOCKS") {
			global.ProxyList.Remove(soket5Proxy)
			fmt.Println("剩余代理数量:" + gconv.String(global.ProxyList.Size()))
		} else {
			global.ProxySuccess = append(global.ProxySuccess, soket5Proxy)
		}
		if flag == false {
			fmt.Println("成功代理数量:" + gconv.String(len(global.ProxySuccess)))
			fmt.Println(global.ProxySuccess)
			global.ProxyList = global.NewSafeProxyList(global.ProxySuccess)
		}
	}()
	if !flag {
		err = gerror.New("没有可用的代理")
		return
	}
	soket5ProxyV := "socks5://" + soket5Proxy

	opts := ChromedpOpt(soket5ProxyV)
	// 创建无头浏览器实例
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// 创建上下文
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 设置超时
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// 用于存储页面HTML的变量

	// 运行任务
	err = chromedp.Run(ctx,
		// 访问目标URL
		chromedp.Navigate(url),
		// 等待页面加载完成
		chromedp.WaitVisible("body", chromedp.ByQuery),
		// 等待更长时间确保内容加载
		chromedp.Sleep(5*time.Second),
		// 获取页面HTML
		chromedp.OuterHTML("html", &html),
	)
	if err != nil {
		log.Printf("ChromeDP运行错误: %v", err)
		return
	}
	return

}

// ChromedpOpt 无头浏览器爬虫
func ChromedpOpt(soket5Proxy string) []chromedp.ExecAllocatorOption {
	userAgent := global.RandomUserAgent()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-features", "VizDisplayCompositor"),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-plugins", true),
		chromedp.Flag("disable-images", true),
		chromedp.Flag("disable-javascript", false),
		chromedp.UserAgent(userAgent),
		// 设置代理（如果需要）
		chromedp.ProxyServer(soket5Proxy),
	)
	return opts
}
