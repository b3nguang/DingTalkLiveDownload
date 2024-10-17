package main

import (
	"github.com/b3nguang/DingTalkLiveDownload/m3u8_plugin"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
	"time"
)

func M3u8Down(title, playbackUrl string, thread int) {
	m3u8 := m3u8_plugin.NewDownloader()
	m3u8.SetUrl(playbackUrl)
	m3u8.SetMovieName(title)
	m3u8.SetNumOfThread(thread)
	m3u8.SetIfShowTheBar(true)
	m3u8.DefaultDownload()
	os.Exit(0)
}

func startChrome() {
	fmt.Println("🔥 正在启动Chrome")
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:], // 使用完整的默认选项
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false), // 如果需要无头模式则改为 true
		chromedp.Flag("disable-gpu", true),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	// 创建 Chrome 上下文
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 设置网络请求捕获逻辑
	var requests []string
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			url := ev.Request.URL
			if url != "" && isM3U8Url(url) {
				fmt.Println(url)
				requests = append(requests, url)
				log.Printf("捕获到 M3U8 URL: %s\n", url)
				M3u8Down("钉钉直播视频", url, 10)
				cancel() // 下载后关闭浏览器
			}
		}
	})

	// 启动浏览器并导航到指定页面
	err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate(`https://n.dingtalk.com/dingding/live-room/index.html?roomId=XMaJOOwJi5aVizzV&liveUuid=38d2f5a5-8918-46d9-bcfe-66f3b269b7fa`),
		chromedp.Sleep(10*time.Second), // 等待页面加载
	)
	if err != nil {
		log.Fatalf("浏览器运行错误: %v", err)
	}

	// 输出捕获的 M3U8 请求
	for _, req := range requests {
		fmt.Println("捕获的 M3U8 请求:", req)
	}

	fmt.Println("按下回车键退出...")
	fmt.Scanln()
}

// 更严格的 URL 捕获逻辑
func isM3U8Url(url string) bool {
	return strings.Contains(url, "/live_hp/") && strings.Contains(url, "auth_key")
}

func main() {
	startChrome()
}
