// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package genrouter

import "hotgo/internal/controller/crawler"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, crawler.TgstatChannel) // tgstat频道
}
