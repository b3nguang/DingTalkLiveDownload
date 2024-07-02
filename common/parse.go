package common

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Readfile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Open %s error, %v\n", filename, err)
		os.Exit(0)
	}
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			content = append(content, scanner.Text())
		}
	}
	return content, nil
}

func Parse(Info *Info) {
	//var UrlList []string

	if Info.URL == "" && Info.UrlFile == "" {
		fmt.Println("未提供 URL 或 URL 文件路径，退出...")
		flag.Usage()
		os.Exit(0)
	}

	if Info.URL != "" {
		urls := strings.Split(Info.URL, ",")
		TmpUrls := make(map[string]struct{})
		for _, url := range urls {
			if _, ok := TmpUrls[url]; !ok {
				TmpUrls[url] = struct{}{}
				if url != "" {
					URLs = append(URLs, url)
				}
			}
		}
	}
	if Info.UrlFile != "" {
		urls, err := Readfile(Info.UrlFile)
		if err == nil {
			TmpUrls := make(map[string]struct{})
			for _, url := range urls {
				if _, ok := TmpUrls[url]; !ok {
					TmpUrls[url] = struct{}{}
					if url != "" {
						URLs = append(URLs, url)
					}
				}
			}
		}
	}
}
