package main

import (
	"context"
	"fmt"
	"github.com/theone-daxia/crawler/collect"
	"log"
	"regexp"
	"time"

	"github.com/chromedp/chromedp"
)

var titleRe = regexp.MustCompile(`<div class="small_toplink__GmZhY"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {
	//getFromPP() // 获取澎湃新闻 card title
	//getFromDB() // 获取豆瓣图书信息
	getByChromedp() // 利用
}

func getFromPP() {
	url := "https://www.thepaper.cn/channel_108856"
	var f collect.Fetcher = &collect.BaseFetch{}

	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}

	matches := titleRe.FindAllSubmatch(body, -1)
	for _, m := range matches {
		fmt.Println("fetch card news:", string(m[1]))
	}
}

func getFromDB() {
	url := "https://book.douban.com/subject/1007305/"
	var f collect.Fetcher = &collect.BrowserFetch{}

	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}

	fmt.Println(string(body))
}

func getByChromedp() {
	// 1. 创建谷歌浏览器实例
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 2. 设置 context 超时时间
	ctx, cancel = context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	// 3. 爬取页面，等待某一个元素出现，接着模拟鼠标点击，最后获取数据
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		chromedp.Value(`#example-After textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
