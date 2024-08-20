package main

import (
	"github.com/b3nguang/DingTalkLiveDownload/m3u8_plugin"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

func M3u8Down(title, playbackUrl string, Thread int) {
	m3u8 := m3u8_plugin.NewDownloader()
	m3u8.SetUrl(playbackUrl)
	m3u8.SetMovieName(title)
	m3u8.SetNumOfThread(Thread)
	m3u8.SetIfShowTheBar(true)
	m3u8.DefaultDownload()
}

func startChrome() {
	fmt.Println("🔥 正在启动Chrome")
	opts := append(
		chromedp.DefaultExecAllocatorOptions[3:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	)

	parentCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(parentCtx)
	defer cancel()

	var requests []string

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			url := ev.Request.URL
			//fmt.Println(url)
			if url != "" && urlContains(ev.Request.URL, "/live_hp/") && urlContains(ev.Request.URL, "_normal.m3u8") {
				requests = append(requests, ev.Request.URL)
				//fmt.Println(ev.Request.URL)
				M3u8Down("钉钉直播视频", ev.Request.URL, 10)
				//cancel()
			}
		} // 这里添加关闭 switch 语句的大括号
	})

	// 启动浏览器，访问目标网址
	err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate(`https://n.dingtalk.com/dingding/live-room/index.html?roomId=xxx&liveUuid=xxx`),  //修改这里
		chromedp.Sleep(10*time.Second), // 适当等待以加载和捕获请求

	)
	if err != nil {
		log.Fatal(err)
	}

	// 打印捕获的请求
	for _, req := range requests {
		fmt.Println(req)
	}

	// 等待用户按下回车键后退出
	fmt.Println("按下回车键退出...")
	fmt.Scanln()
}

func urlContains(url, substr string) bool {
	return strings.Contains(url, substr)
}

func main() {
	startChrome()
}
