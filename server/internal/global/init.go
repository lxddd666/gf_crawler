// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
	"runtime"
	"strings"
)

// ParsedProxyList 解析后的代理列表
var ParsedProxyList []string

// ProxyInfo 代理信息结构体
type ProxyInfo struct {
	IpAddress string `json:"ip_address"`
	Port      int    `json:"port"`
}

func Init(ctx context.Context) {
	// 设置gf运行模式
	SetGFMode(ctx)

	// 设置服务日志处理
	glog.SetDefaultHandler(LoggingServeLogHandler)

	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err：%+v", err)
		return
	}

	fmt.Printf("欢迎使用HotGo！\r\n当前运行环境：%v, 运行根路径为：%v \r\nHotGo版本：v%v, gf版本：%v \n", runtime.GOOS, gfile.Pwd(), consts.VersionApp, gf.VERSION)

	// 初始化链路追踪
	InitTrace(ctx)

	// 设置缓存适配器
	cache.SetAdapter(ctx)

	// 初始化功能库配置
	service.SysConfig().InitConfig(ctx)

	// 加载超管数据
	service.AdminMember().LoadSuperAdmin(ctx)

	// 订阅集群同步
	SubscribeClusterSync(ctx)
	// 初始化代理池塘
	// 解析代理文件
	parseJsonFile(ctx)

	// 初始化代理池
	InitProxyPool()

	UserAgentList = NewUserAgentList()

	ProxySuccess = make([]string, 0)
}

func InitProxyPool() {

	ProxyList = NewSafeProxyList(proxyList())
}

func NewUserAgentList() []string {
	return []string{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60",
		"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.50",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0",
		"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.133 Safari/534.16",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/2.0 Safari/536.11",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; LBBROWSER)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E; LBBROWSER",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 SE 2.X MetaSr 1.0",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SV1; QQDownload 732; .NET4.0C; .NET4.0E; SE 2.X MetaSr 1.0)",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Maxthon/4.4.3.4000 Chrome/30.0.1599.101 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 UBrowser/4.0.3214.0 Safari/537.36",
		"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/5.0 (iPod; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/5.0 (iPad; U; CPU OS 4_2_1 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8C148 Safari/6533.18.5",
		"Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/5.0 (Linux; U; Android 2.2.1; zh-cn; HTC_Wildfire_A3333 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 3.0; en-us; Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		"Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en) AppleWebKit/534.1+ (KHTML, like Gecko) Version/6.0.0.337 Mobile Safari/534.1+",
		"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.0; U; en-US) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/233.70 Safari/534.6 TouchPad/1.0",
		"Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 NokiaN97-1/20.0.019; Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/525 (KHTML, like Gecko) BrowserNG/7.1.18124",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Titan)",
		"Mozilla/4.0 (compatible; MSIE 6.0; ) Opera/UCWEB7.0.2.37/28/999",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 SE 2.X MetaSr 1.0",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.26 Safari/537.36 Core/1.63.5680.400 QQBrowser/10.2.1852.400",
		"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:65.0) Gecko/20100101 Firefox/65.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.3 Safari/605.1.15",
		"Mozilla/5.0 (Linux; Android 4.2.1; M040 Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; M351 Build/KTU84P) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 5.0.2; zh-cn; X900 Build/CBXCNOP5500912251S) AppleWebKit/533.1 (KHTML, like Gecko)Version/4.0 MQQBrowser/5.4 TBS/025489 Mobile Safari/533.1 V1_AND_SQ_6.0.0_300_YYB_D QQ/6.0.0.2605 NetType/WIFI WebP/0.3.0 Pixel/1440",
		"Mozilla/5.0 (Linux; U; Android 5.0.2; zh-cn; NX511J Build/LRX22G) AppleWebKit/533.1 (KHTML, like Gecko)Version/4.0 MQQBrowser/8.8 TBS/88888888 Mobile Safari/533.1 MicroMessenger/6.3.8.56_re6b2553.680 NetType/ctlte Language/zh_CN MicroMessenger/6.3.8.56_re6b2553.680 NetType/ctlte Language/zh_CN",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) CriOS/31.0.1650.18 Mobile/11B554a Safari/8536.25",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_2_1 like Mac OS X) AppleWebKit/602.4.6 (KHTML, like Gecko) Mobile/14D27 QQ/6.7.1.416 V1_IPH_SQ_6.7.1_1_APP_A Pixel/750 Core/UIWebView NetType/4G QBWebViewType/1",
	}
}

func parseJsonFile(ctx context.Context) {
	// 代理文件路径
	proxyFilePath := "internal/global/proxy_file/proxies.json"

	// 检查文件是否存在
	if !gfile.Exists(proxyFilePath) {
		g.Log().Warningf(ctx, "代理文件不存在: %s", proxyFilePath)
		return
	}

	// 读取JSON文件内容
	jsonContent := gfile.GetContents(proxyFilePath)
	if jsonContent == "" {
		g.Log().Warning(ctx, "代理文件内容为空")
		return
	}

	// 解析JSON数据
	var proxyInfos []ProxyInfo
	if err := gjson.DecodeTo(jsonContent, &proxyInfos); err != nil {
		g.Log().Errorf(ctx, "解析代理JSON文件失败: %v", err)
		return
	}

	// 转换为ip:port格式的字符串数组
	ParsedProxyList = make([]string, 0, len(proxyInfos))
	for _, proxy := range proxyInfos {
		proxyStr := fmt.Sprintf("%s:%d", proxy.IpAddress, proxy.Port)
		ParsedProxyList = append(ParsedProxyList, proxyStr)
	}

	g.Log().Infof(ctx, "成功解析代理文件，共加载 %d 个代理", len(ParsedProxyList))
}

func proxyList() []string {
	// 优先使用从JSON文件解析的代理列表

	// 如果没有解析到代理，使用硬编码的备用代理列表
	g.Log().Warning(gctx.New(), "未找到有效的代理文件，使用硬编码的备用代理列表")
	// https://fineproxy.org/cn/free-proxy/
	proxtys := []string{
		"98.191.0.47:4145",
		"198.8.94.170:4145",
		"208.102.51.6:58208",
		"184.181.178.33:4145",
		"192.111.130.2:4145",
		"192.111.139.163:19404",
		"70.166.65.160:4145",
		"174.77.111.197:4145",
		"174.77.111.198:49547",
		"174.64.199.82:4145",
		"116.98.50.140:1080",
		"8.211.42.167:4145",
		"98.178.72.21:10919",
		"98.188.47.132:4145",
		"98.188.47.150:4145",
		"47.252.11.233:8181",
		"98.175.31.195:4145",
		"98.170.57.249:4145",
		"98.170.57.231:4145",
		"72.195.34.58:4145",
		"72.195.34.59:4145",
		"72.195.34.60:27391",
		"72.195.114.169:4145",
		"72.195.114.184:4145",
		"72.195.34.41:4145",
		"72.195.34.42:4145",
		"70.166.167.55:57745",
		"72.49.49.11:31034",
		"72.195.34.35:27360",
		"69.61.200.104:36181",
		"67.201.33.10:25283",
		"199.116.114.11:4145",
		"66.42.224.229:41679",
		"72.211.46.124:4145",
		"24.249.199.4:4145",
		"24.249.199.12:4145",
		"47.252.18.37:999",
		"47.252.18.37:45",
		"8.211.51.115:104",
		"142.54.236.97:4145",
		"142.54.228.193:4145",
		"47.91.89.3:8080",
		"45.119.55.129:1080",
		"47.90.167.27:8880",
		"103.82.27.24:10001",
		"192.111.137.34:18765",
		"192.111.137.35:4145",
		"192.111.137.37:18762",
		"192.111.138.29:4145",
		"192.111.139.162:4145",
		"192.111.139.165:4145",
		"192.252.209.155:14455",
		"192.252.211.197:14921",
		"89.169.36.109:1080",
		"8.209.96.245:45",
		"57.129.81.201:1080",
		"222.59.173.105:44077",
		"68.71.251.134:4145",
		"198.177.253.13:4145",
		"199.187.210.54:4145",
		"192.111.130.5:17002",
		"192.111.135.17:18302",
		"192.111.135.18:18301",
		"103.82.25.14:10001",
		"72.214.108.67:4145",
		"162.253.68.97:4145",
		"8.220.200.221:8080",
		"192.252.215.5:16137",
		"184.178.172.13:15311",
		"70.166.167.38:57728",
		"198.8.94.174:39078",
		"192.252.210.233:4145",
		"72.223.188.92:4145",
		"72.211.46.99:4145",
		"143.110.217.153:1080",
		"67.201.58.190:4145",
		"68.71.242.118:4145",
		"199.102.104.70:4145",
		"98.170.57.241:4145",
		"72.37.216.68:4145",
		"142.54.239.1:4145",
		"67.201.35.145:4145",
		"206.220.175.2:4145",
		"199.58.185.9:4145",
		"184.170.249.65:4145",
		"184.170.245.148:4145",
		"199.58.184.97:4145",
		"98.191.0.37:4145",
		"47.238.226.127:1024",
		"72.211.46.99:4145",
		"46.4.88.72:9050",
		"72.223.188.67:4145",
		"72.207.113.97:4145",
		"68.71.249.153:48606",
		"199.102.106.94:4145",
		"199.102.105.242:4145",
		"199.102.107.145:4145",
		"72.206.74.126:4145",
		"98.182.147.97:4145",
		"107.181.168.145:4145",
		"72.37.217.3:4145",
		"192.252.215.2:4145",
		"68.71.249.158:4145",
		"68.71.252.38:4145",
		"85.111.94.98:15833",
		"127.0.0.1:7890",
		"127.0.0.1:7890",
		"127.0.0.1:7890",
		"192.252.209.158:4145",
		"98.190.239.3:4145",
		"199.116.112.6:4145",
		"184.170.251.30:11288",
		"98.182.171.161:4145",
		"192.111.129.150:4145",
		"192.252.214.17:4145",
		"174.75.211.193:4145",
		"198.177.252.24:4145",
		"142.54.237.38:4145",
		"68.71.245.206:4145",
		"98.175.31.222:4145",
		"68.71.243.14:4145",
		"68.71.240.210:4145",
		"68.71.254.6:4145",
		"72.207.109.5:4145",
		"74.119.144.60:4145",
		"98.181.137.80:4145",
		"68.1.210.163:4145",
		"98.181.137.83:4145",
		"68.71.247.130:4145",
		"74.119.147.209:4145",
		"192.252.216.81:4145",
		"184.178.172.17:4145",
		"72.205.0.67:4145",
		"198.177.254.157:4145",
		"192.252.220.92:17328",
		"184.181.217.201:4145",
		"184.181.217.206:4145",
		"184.181.217.210:4145",
		"184.181.217.213:4145",
		"184.181.217.220:4145",
		"198.177.254.157:4145",
		"192.252.220.92:17328",
		"184.181.217.201:4145",
		"184.181.217.206:4145",
		"184.181.217.210:4145",
		"184.181.217.213:4145",
		"184.181.217.220:4145",
		"184.178.172.11:4145",
		"184.178.172.14:4145",
		"184.178.172.18:15280",
		"184.178.172.23:4145",
		"184.178.172.25:15291",
		"184.178.172.26:4145",
		"184.178.172.28:15294",
		"184.181.217.194:4145",
		"184.178.172.3:4145",
		"184.178.172.5:15303",
		"174.64.199.79:4145",
		"174.77.111.196:4145",
		"198.177.254.131:4145",
		"68.1.210.189:4145",
		"72.205.0.93:4145",
		"192.252.216.86:4145",
		"192.252.211.193:4145",
		"68.71.241.33:4145",
		"67.201.39.14:4145",
		"174.75.211.222:4145",
		"192.252.220.89:4145",
		"107.152.98.5:4145",
		"198.8.84.3:4145",
		"142.54.232.6:4145",
		"142.54.235.9:4145",
		"142.54.229.249:4145",
		"142.54.237.34:4145",
		"142.54.231.38:4145",
		"68.71.251.134:4145",
		"198.177.253.13:4145",
		"199.187.210.54:4145",
		"192.111.135.17:18302",
		"72.214.108.67:4145",
		"103.82.25.14:10001",
		"198.8.94.174:39078",
		"192.252.215.5:16137",
	}
	if len(ParsedProxyList) > 0 {
		g.Log().Infof(gctx.New(), "使用JSON文件中的代理列表，共 %d 个代理", len(ParsedProxyList))
		ParsedProxyList = append(ParsedProxyList, "127.0.0.1:7890",
			"127.0.0.1:7890",
			"127.0.0.1:7890",
			"127.0.0.1:7890",
			"127.0.0.1:7890")
		proxtys = append(proxtys, ParsedProxyList...)
	}
	return proxtys
}

// LoggingServeLogHandler 服务日志处理
// 需要将异常日志保存到服务日志时可以通过SetHandlers设置此方法
func LoggingServeLogHandler(ctx context.Context, in *glog.HandlerInput) {
	in.Next(ctx)

	err := g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				panic(err)
			}
		}()

		// web服务日志不做记录，因为会导致重复记录
		r := g.RequestFromCtx(ctx)
		if r != nil && r.Server != nil && in.Logger.GetConfig().Path == r.Server.Logger().GetConfig().Path {
			return
		}

		conf, err := service.SysConfig().GetLoadServeLog(ctx)
		if err != nil {
			return
		}

		if conf == nil {
			return
		}

		if !conf.Switch {
			return
		}

		if in.LevelFormat == "" || !gstr.InArray(conf.LevelFormat, in.LevelFormat) {
			return
		}

		if in.Stack == "" {
			in.Stack = in.Logger.GetStack()
		}

		if len(in.Content) == 0 {
			in.Content = gstr.StrLimit(gvar.New(in.Values).String(), consts.MaxServeLogContentLen)
		}

		var data entity.SysServeLog
		data.TraceId = gctx.CtxId(ctx)
		data.LevelFormat = in.LevelFormat
		data.Content = in.Content
		data.Stack = gjson.New(charset.ParseStack(in.Stack))
		data.Line = strings.TrimRight(in.CallerPath, ":")
		data.TriggerNs = in.Time.UnixNano()
		data.Status = consts.StatusEnabled

		if gstr.Contains(in.Content, `exception recovered`) {
			data.LevelFormat = "PANI"
		}

		if data.Stack.IsNil() {
			data.Stack = gjson.New(consts.NilJsonToString)
		}

		if conf.Queue {
			err = queue.Push(consts.QueueServeLogTopic, data)
		} else {
			err = service.SysServeLog().RealWrite(ctx, data)
		}
	})

	if err != nil {
		g.Dump("LoggingServeLogHandler err:", err)
	}
}

// InitTrace 初始化链路追踪
func InitTrace(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, "jaeger.switch").Bool() {
		return
	}

	tp, err := jaeger.Init(simple.AppName(ctx), g.Cfg().MustGet(ctx, "jaeger.endpoint").String())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		_ = tp.Shutdown(ctx)
		g.Log().Debug(ctx, "jaeger closed ..")
	})
}

// SetGFMode 设置gf运行模式
func SetGFMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "system.mode").String()
	if len(mode) == 0 {
		mode = gmode.NOT_SET
	}

	var modes = []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}

	// 如果是有效的运行模式，就进行设置
	if validate.InSlice(modes, mode) {
		gmode.Set(mode)
	}
}
