package common

import (
	"flag"
)

func Banner() {
	banner := `DingTalkLiveDownLoad  author: ` + author + `

`
	print(banner)
}
func Flag(Info *Info) {
	Banner()
	flag.StringVar(&Info.URL, "u", "", "需要下载的回放URL")
	flag.StringVar(&Info.UrlFile, "f", "", "包含需要下载的回放URL的文件路径")
	flag.Parse()
}
