package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/antchfx/xquery/html"
)

var (
	expr = regexp.MustCompile("http://www.meizitu.com/a/([0-9]+)\\.html")
)

func getAllUrls() []string {
	// http://www.meizitu.com/a/more_1.html
	// http://www.meizitu.com/a/more_72.html
	var urls []string
	var url string
	for i := 0; i < 72; i++ {
		url = "http://www.meizitu.com/a/more_" + strconv.Itoa(i+1) + ".html"
		urls = append(urls, url)
	}
	links := GetSubUrlFromPage(urls)
	return links
}

func GetSubUrlFromPage(p []string) []string {
	var links []string
	for i := 0; i < len(p); i++ {
		doc, _ := htmlquery.LoadURL(p[i])
		for _, n := range htmlquery.Find(doc, "//div[@class='inWrap']//a/@href") {
			link := htmlquery.SelectAttr(n, "href")
			fmt.Printf("#Line 33: link = %s\n", link)
			link = expr.FindString(link) //使用 regexp来获取 网页链接地址
			fmt.Printf("link = %s\n", link)
			if link != "" {
				links = append(links, link)
			}
		}
	}
	return links
}

func parseHtml(url string) int {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range htmlquery.Find(doc, "//div[@class='postContent']//img/@src") {
		img_url := htmlquery.SelectAttr(n, "src")
		alt := htmlquery.SelectAttr(n, "alt")
		fmt.Printf("Alt = %s\n", alt)
		go download(img_url, alt)
	}

	return 0
}

func main() {
	urls := getAllUrls()
	for _, url := range urls {
		parseHtml(url)
	}
}
