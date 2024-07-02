package download_plugin

import (
	"fmt"
	"net/url"
)

func process(target_url string) {
	// 解析 URL
	parsedURL, err := url.Parse(target_url)
	if err != nil {
		fmt.Println("解析 URL 时出错:", err)
		return
	}
	// 提取查询参数中的 roomId 和 liveUuid
	queryParams := parsedURL.Query()
	roomId := queryParams.Get("roomId")
	liveUuid := queryParams.Get("liveUuid")
	if roomId == "" || liveUuid == "" {
		fmt.Println("URL 中缺少 roomId 或 liveUuid 参数，退出...")
		return
	}

	// 调用函数
	// 假设你有一个处理这些信息的函数
	getLiveRoomPublicInfo(roomId, liveUuid, 10)
}
