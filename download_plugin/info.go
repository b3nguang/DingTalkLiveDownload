package download_plugin

import (
	"DingTalkLiveDownload/common"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func getLiveRoomPublicInfo(roomId, liveUuid string, Thread int) {
	// 构造URL
	urlStr := "https://lv.dingtalk.com/getOpenLiveInfo?roomId=" + roomId + "&liveUuid=" + liveUuid
	urlObj, _ := url.Parse(urlStr)

	// 创建请求
	req, _ := http.NewRequest("GET", urlObj.String(), nil)

	// 读取Cookies.json文件
	jsonCookies, _ := os.ReadFile("cookies.json")
	var cookies map[string]string
	_ = json.Unmarshal(jsonCookies, &cookies)

	// 添加Cookies到请求
	var cookieStr strings.Builder
	for name, value := range cookies {
		cookieStr.WriteString(fmt.Sprintf("%s=%s; ", name, value))
	}
	cookieHeader := cookieStr.String()
	CookiepcSession := cookies["LV_PC_SESSION"]
	// 设置请求头
	req.Header.Set("Host", common.Host)
	req.Header.Set("Cookie", cookieHeader)
	req.Header.Set("Cookie", "PC_SESSION="+CookiepcSession)
	req.Header.Set("Sec-Ch-Ua", common.SecChUa)
	req.Header.Set("Sec-Ch-Ua-Mobile", common.SecChUaMobile)
	req.Header.Set("Sec-Ch-Ua-Platform", common.SecChUaPlatform)
	req.Header.Set("Dnt", common.Dnt)
	req.Header.Set("Upgrade-Insecure-Requests", common.UpgradeInsecureRequests)
	req.Header.Set("User-Agent", common.UserAgent)
	req.Header.Set("Accept", common.Accept)
	req.Header.Set("Sec-Fetch-Site", common.SecFetchSite)
	req.Header.Set("Sec-Fetch-Mode", common.SecFetchMode)
	req.Header.Set("Sec-Fetch-User", common.SecFetchUser)
	req.Header.Set("Sec-Fetch-Dest", common.SecFetchDest)
	req.Header.Set("Accept-Language", common.AcceptLanguage)

	// 发送请求
	client := &http.Client{}
	resp, _ := client.Do(req)

	// 关闭响应
	defer resp.Body.Close()

	// 读取响应内容
	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	title := result["openLiveDetailModel"].(map[string]interface{})["title"].(string)
	playbackUrl := result["openLiveDetailModel"].(map[string]interface{})["playbackUrl"].(string)

	// 获取当前时间并格式化
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	title += "_" + currentTime

	fmt.Println("标题:", title)
	fmt.Println("请求网址:", playbackUrl)
	M3u8Down(title, playbackUrl, Thread)
}
