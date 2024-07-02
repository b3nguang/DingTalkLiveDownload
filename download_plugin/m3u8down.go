package download_plugin

import "DingTalkLiveDownload/m3u8_plugin"

func M3u8Down(title, playbackUrl string, Thread int) {
	m3u8 := m3u8_plugin.NewDownloader()
	m3u8.SetUrl(playbackUrl)
	m3u8.SetMovieName(title)
	m3u8.SetNumOfThread(Thread)
	m3u8.SetIfShowTheBar(true)
	if m3u8.DefaultDownload() {
		ffmpeg(title)
	}
}
