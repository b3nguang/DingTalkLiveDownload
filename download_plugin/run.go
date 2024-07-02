package download_plugin

import (
	"DingTalkLiveDownload/common"
	"fmt"
	"os"
)

func Run(Info *common.Info) {
	if _, err := os.Stat("cookies.json"); err == nil {
		// 文件存在，跳过startChrome()
		fmt.Println("cookies.json存在\n")
	} else {
		// 文件不存在，执行startChrome(
		fmt.Println("cookies.json不存在\n")
		startChrome()
	}
	for _, url := range common.URLs {
		process(url)
	}
}
