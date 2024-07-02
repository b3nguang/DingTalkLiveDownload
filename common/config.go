package common

var author = "b3nguang"

var (
	Host                    = "lv.dingtalk.com"
	UserAgent               = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"
	Dnt                     = "1"
	SecChUa                 = `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`
	SecChUaMobile           = "?0"
	SecChUaPlatform         = "macOS"
	UpgradeInsecureRequests = "1"
	Accept                  = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
	SecFetchSite            = "none"
	SecFetchMode            = "navigate"
	SecFetchUser            = "?1"
	SecFetchDest            = "document"
	AcceptLanguage          = "zh-CN,zh;q=0.9"
)

type Info struct {
	URL     string
	UrlFile string
}

var URLs = []string{}
