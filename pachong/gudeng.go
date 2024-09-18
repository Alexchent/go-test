package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// golang goquery selector(选择器) 示例大全 https://cloud.tencent.com/developer/article/1196783

func ScrapeFromUrl(url string) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http.get err:" + err.Error())
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	work(res.Body)
}

func ScrapeFromFile(url string) {
	file, err := os.Open(url)
	if err != nil {
		fmt.Println("http.get err:" + err.Error())
		return
	}
	defer file.Close()

	work(file)
}

func work(r io.Reader) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatalf("goquery.NewDocumentFromReader err: %s", err.Error())
	}

	subtitle := doc.Find("#pg-title-no-subtitle").Text()
	fmt.Println(subtitle)
	filename := subtitle + ".txt"
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// Find the review items
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		// 获取p标签，且id属性包含id的元素
		// content := s.Find("p[id*=id]").Text()
		s.Find("p[id*=id]").Each(func(ii int, sc *goquery.Selection) {
			content := sc.Text()
			fmt.Println(content)
			fmt.Println(strings.Repeat("=", 20))
			_, err2 := fd.WriteString(content + "\r\n")
			if err2 != nil {
				return
			}
		})
	})
}

func main() {
	//url := "https://www.gutenberg.org/cache/epub/25271/pg25271-images.html"
	//url := "https://www.gutenberg.org/cache/epub/23962/pg23962-images.html"
	//ScrapeFromUrl(url)
	ScrapeFromFile("xixiangji.html")
}
