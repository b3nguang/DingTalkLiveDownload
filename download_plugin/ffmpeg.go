package download_plugin

import (
	"fmt"
	"os"
	"os/exec"
)

func ffmpeg(ts string) {
	fmt.Println("正在转换ts为mp4...")
	cmd := exec.Command("ffmpeg", "-i", "video/"+ts+".ts", "-c:v", "copy", "-c:a", "copy", "-f", "mp4", "-y", ts+".mp4")
	err := cmd.Run()
	if err != nil {
		return
	}
	fmt.Println(ts + ".mp4 转换完成\n")
	os.Remove("video/" + ts + ".ts")
}
