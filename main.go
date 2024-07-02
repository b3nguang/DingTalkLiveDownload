package main

import (
	"DingTalkLiveDownload/common"
	"DingTalkLiveDownload/download_plugin"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var Info common.Info
	common.Flag(&Info)
	common.Parse(&Info)
	download_plugin.Run(&Info)
	fmt.Printf("[*] 下载结束,耗时: %s\n", time.Since(start))
}
