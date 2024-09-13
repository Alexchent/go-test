package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// golang goquery selector(选择器) 示例大全 https://cloud.tencent.com/developer/article/1196783

func ExampleScrape() {
	// Request the HTML page.
	//res, err := http.Get("https://www.gutenberg.org/cache/epub/25271/pg25271-images.html")
	res, err := http.Get("https://www.gutenberg.org/cache/epub/23962/pg23962-images.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	subtitle := doc.Find("#pg-title-no-subtitle").Text()
	fmt.Println(subtitle)
	filename := subtitle + ".txt"
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// Find the review items
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		// 获取p标签，且id属性包含id的元素
		//node := s.Find("p[id*=id]")
		//a, ok := node.Attr("id")
		//if ok {
		//	fmt.Println(a)
		//}
		content := s.Find("p[id*=id]").Text()
		_, err2 := fd.WriteString(content + "\r\n")
		if err2 != nil {
			return
		}
	})
}

type P struct {
	fds map[string]*os.File
}

func (i *P) Append(file, content string) error {
	fd, ok := i.fds[file]
	if !ok {
		fd, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		i.fds[file] = fd
	}
	_, err := fd.WriteString(content + "\n")
	return err
}

func main() {
	ExampleScrape()
}
