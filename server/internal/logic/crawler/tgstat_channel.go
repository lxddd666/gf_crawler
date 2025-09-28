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
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/crawlerin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/simple"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/proxy"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"

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

func (s *sCrawlerTgstatChannel) CrawlerChannelUrl(ctx context.Context, in *crawlerin.TgstatChannelCrawlerChannelInp) (err error) {
	// 创建默认收集器
	c := colly.NewCollector()
	// 向 API 发送请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("发送请求到：", r.URL)
	})
	list := make([]entity.TgstatChannelCrawlerUrl, 0)
	// 处理 API 返回的数据
	c.OnResponse(func(r *colly.Response) {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(gconv.String(r.Body)))
		if err != nil {
			return
		}
		// 匹配包含指定类名的a标签
		doc.Find("a.list-group-item.list-group-item-action.px-2.pl-lg-3.border-hover-info-right-3px").Each(func(i int, s *goquery.Selection) {
			href := s.AttrOr("href", "")
			text := strings.TrimSpace(s.Text())
			fmt.Printf("找到链接: href=%s, text=%s\n", href, text)
			if href != "" {
				href = "https://tgstat.com" + href
				for _, countryUrl := range consts.RatingCountryList {
					getUrl := href
					// 国家替换
					getUrl = gstr.Replace(getUrl, "tgstat.com", countryUrl)

					// 添加 public / private
					for _, disclosure := range consts.PrivateDisclosure {
						disclosureUrl := getUrl
						if disclosure != "" {
							disclosureUrl = addPathPrefix(disclosureUrl, disclosure)
						}

						for _, sort := range consts.Sort {
							disclosureUrl, _ = replaceSortParam(disclosureUrl, sort)
							list = append(list, entity.TgstatChannelCrawlerUrl{Url: disclosureUrl, Type: consts.CrawlerRating, Status: consts.TaskWait})
						}
					}
				}
			}
			// 如果需要将数据添加到list中，可以创建对应的结构体
		})
		// 保存收集的数据到数据库 - 分批插入，每次300个
		if len(list) > 0 {
			batchSize := 300
			for i := 0; i < len(list); i += batchSize {
				end := i + batchSize
				if end > len(list) {
					end = len(list)
				}
				batch := list[i:end]
				// 使用正确的DAO操作TgstatChannelCrawlerUrl表
				_, err := dao.TgstatChannelCrawlerUrl.Ctx(ctx).FieldsEx(dao.TgstatChannelCrawlerUrl.Columns().Id).InsertIgnore(batch)
				if err != nil {
					fmt.Printf("批量插入第%d-%d条记录失败: %v\n", i+1, end, err)
				} else {
					fmt.Printf("成功插入第%d-%d条记录，共%d条\n", i+1, end, end-i)
				}
			}
		}

	})
	c.SetRequestTimeout(20 * time.Second)
	// 访问起始页面
	c.Visit("https://tgstat.com/ratings/channels?sort=members")
	return
}

func addPathPrefix(url, prefix string) string {
	// 分割URL，获取路径和查询参数
	parts := strings.SplitN(url, "?", 2)
	path := parts[0]

	// 如果有查询参数，添加回去
	var params string
	if len(parts) > 1 {
		params = "?" + parts[1]
	}

	return path + prefix + params
}

func replaceSortParam(rawURL, newValue string) (string, error) {
	// 解析URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// 获取查询参数
	queryParams := parsedURL.Query()

	// 设置新的sort参数值
	queryParams.Set("sort", newValue)

	// 更新查询参数
	parsedURL.RawQuery = queryParams.Encode()

	return parsedURL.String(), nil
}

func (s *sCrawlerTgstatChannel) StartRating(ctx context.Context) (err error) {
	simple.SafeGo(gctx.New(), func(ctx context.Context) {
		var list []*entity.TgstatChannelCrawlerUrl
		err = dao.TgstatChannelCrawlerUrl.Ctx(ctx).WhereNot(dao.TgstatChannelCrawlerUrl.Columns().Status, consts.TaskSuccess).Scan(&list)
		if err != nil {
			return
		}
		concurrencyLimit := 10
		semaphore := make(chan struct{}, concurrencyLimit)
		wg := sync.WaitGroup{}
		for _, inp := range list {
			wg.Add(1)
			semaphore <- struct{}{}

			simple.SafeGo(gctx.New(), func(ctx context.Context) {
				defer func() {
					wg.Done()
				}()
				defer func() {
					// 释放信号量
					<-semaphore
				}()
				//s.CrawlerChannel(gctx.New(), inp)
				s.CrawlerChromedpChannel(ctx, inp)
				//time.Sleep(grand.D(5, 8) * time.Second)
			})
		}
		wg.Wait()
		return
	})
	return
}

// CrawlerChannel 抓取tgstat频道
func (s *sCrawlerTgstatChannel) CrawlerChannel(ctx context.Context, in *entity.TgstatChannelCrawlerUrl) (err error) {

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
	// 创建默认收集器
	c := colly.NewCollector()
	// 向 API 发送请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("发送请求到：", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err.Error())
	})
	list := make([]entity.TgstatChannel, 0)
	// 处理 API 返回的数据
	c.OnResponse(func(r *colly.Response) {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(gconv.String(r.Body)))
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
		//if len(list) > 0 {
		//	batchSize := 300
		//	for i := 0; i < len(list); i += batchSize {
		//		end := i + batchSize
		//		if end > len(list) {
		//			end = len(list)
		//		}
		//		batch := list[i:end]
		//		_, err = s.Model(ctx).FieldsEx(dao.TgstatChannel.Columns().Id).InsertIgnore(batch)
		//		if err != nil {
		//			fmt.Printf("TgstatChannel批量插入第%d-%d条记录失败: %v\n", i+1, end, err)
		//		} else {
		//			fmt.Printf("TgstatChannel成功插入第%d-%d条记录，共%d条\n", i+1, end, end-i)
		//		}
		//	}
		//}

	})
	c.SetRequestTimeout(20 * time.Second)
	// 访问起始页面
	// https://tgstat.com/ratings/channels?sort=members
	err = c.Visit(in.Url)
	return
}

func (s *sCrawlerTgstatChannel) CrawlerChannelV2(ctx context.Context, in *entity.TgstatChannelCrawlerUrl) (err error) {
	body, err := SendHttpReq(ctx, in.Url)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(gconv.String(body)))
	if err != nil {
		return
	}
	list := make([]entity.TgstatChannel, 0)

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
	//if len(list) > 0 {
	//	batchSize := 300
	//	for i := 0; i < len(list); i += batchSize {
	//		end := i + batchSize
	//		if end > len(list) {
	//			end = len(list)
	//		}
	//		batch := list[i:end]
	//		_, err = s.Model(ctx).FieldsEx(dao.TgstatChannel.Columns().Id).InsertIgnore(batch)
	//		if err != nil {
	//			fmt.Printf("TgstatChannel批量插入第%d-%d条记录失败: %v\n", i+1, end, err)
	//		} else {
	//			fmt.Printf("TgstatChannel成功插入第%d-%d条记录，共%d条\n", i+1, end, end-i)
	//		}
	//	}
	//}
	return
}

func SendHttpReq(ctx context.Context, url string) (response []byte, err error) {

	log.Println("请求内容 url：", url)
	// 创建SOCKS5拨号器
	//socks := RandomSocks5()
	var respErr error
	var socks5 string
	defer func() {
		if respErr != nil {
			global.ProxyList.Remove(socks5)
		}
	}()
	socks5, flag := global.ProxyList.GetRandom()
	if !flag {
		log.Println("无代理可用", url)
		err = gerror.New("无代理可用")
		return
	}
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:7890", nil, proxy.Direct)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			global.ProxyList.Remove(socks5)
			size := global.ProxyList.Size()
			fmt.Println(size)
		}
	}()
	// 创建HTTP传输层
	transport := &http.Transport{
		Dial: dialer.Dial,
	}

	// 创建HTTP客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	// 设置请求头，模拟真实浏览器
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Cookie", "_ym_uid=1755137155684405945; _ym_d=1755137155; tgstat_settings=cf59d85e5bee03100d6d9b8bf80743b9da88d0639c052902b46b493dc7f37dd1a%3A2%3A%7Bi%3A0%3Bs%3A15%3A%22tgstat_settings%22%3Bi%3A1%3Bs%3A19%3A%22%7B%22fp%22%3A%22DLuEcClm_i%22%7D%22%3B%7D; _tgstat_csrk=5bac79bfb8b8f5eefde1d4ffda6ca0bdc189ac134289e72e78e6271d5f2afa7ea%3A2%3A%7Bi%3A0%3Bs%3A12%3A%22_tgstat_csrk%22%3Bi%3A1%3Bs%3A32%3A%22lR-NkhQ8J69LcJ9fPRmbJNY4cXW5NY13%22%3B%7D; cf_clearance=WcDT.P2UjTWc5OyLs4iQCmjK0urkD7a6dikMtRPGN9M-1759043899-1.2.1.1-W5NhqSEVvFisG3DAbXVVjjMcAkuIjFTwIvoyeN9OkDPSPIUpbdm_K_w9dhq.KCpGghpy5Ta6XnrcTGuJ2U2FDbNNQd2V1gF3TaLMYjTrRq_AYEdZY5zMNZAbvU3xlD87jLDeytReGOANJuz_CAYoJe5F4dvydQ15vM19JKmSLnGq7sHhhk2I1Z_xqVm5PXYrrzLcs_1Cx2uZ2Z_gSXoCbwcJumawMhwfnO53SzV6xCE; _ga_ZEKJ7V8PH3=GS2.1.s1759043896$o11$g1$t1759043902$j54$l0$h0; _ga=GA1.2.282633308.1755137073; _gid=GA1.2.763353112.1759043911; _gat_gtag_UA_104082833_1=1; _ym_isad=2")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Referer", "https://tgstat.ru/en/crypto")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"140\", \"Not=A?Brand\";v=\"24\", \"Google Chrome\";v=\"140\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response = body
	return
}

func extractSubscribers(text string) int64 {
	text = strings.TrimSpace(text)
	if idx := strings.Index(text, "subscribers"); idx != -1 {

		return gconv.Int64(strings.ReplaceAll(strings.TrimSpace(text[:idx]), " ", ""))

	}
	return 0
}

// extractNumber 提取数字，处理带单位的数字（如 "337.4k", "3 277"）
func extractNumber(text string) string {
	text = strings.TrimSpace(text)
	if text == "" {
		return "0"
	}
	return strings.ReplaceAll(text, " ", "")
	// 处理带k单位的数字，如 "337.4k"
}

func extractTelegramLink(fullLink string) (string, int) {
	var respFlag int
	flag := gstr.Contains(fullLink, "@")
	if flag {
		respFlag = 1
	} else {
		respFlag = 0
	}
	// 从链接中提取 channel 和 stat 之间的值
	if strings.Contains(fullLink, "/channel/") && strings.Contains(fullLink, "/stat") {
		// 找到 /channel/ 的位置
		channelIndex := strings.Index(fullLink, "/channel/")
		if channelIndex != -1 {
			// 从 /channel/ 后面开始截取
			start := channelIndex + len("/channel/")
			// 找到 /stat 的位置
			statIndex := strings.Index(fullLink[start:], "/stat")
			if statIndex != -1 {
				// 提取 channel 和 stat 之间的值
				channelName := fullLink[start : start+statIndex]
				// 去掉 @ 符号
				return strings.TrimPrefix(channelName, "@"), respFlag
			}
		}
	}
	return fullLink, respFlag
}
