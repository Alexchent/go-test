package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	file "github.com/alexchen/go_test/File"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// 下载书籍
func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	run(homeDir+"/武侠之无尽恶人", "https://www.qixiashu.com/yue/62/62422/", 133837, 134262)
	//StartWork(939847, 940506, "https://www.qixiashu.com/yue/465/465479/", homeDir+"/影视实际从攻略女主开始")
	//StartWork(133837, 134262, "https://www.qixiashu.com/yue/62/62422/", homeDir+"/武侠之无尽恶人")
}

// 并发下载
func run(saveDir, host string, begin, end int) {
	file.CreateDateDir(saveDir)
	fmt.Println(saveDir)
	var wg sync.WaitGroup
	worker := 8
	ch := make(chan int, worker)
	// 启动4个消费者
	wg.Add(worker)
	for i := 0; i <= worker; i++ {
		go download(host, saveDir, ch, &wg)
	}

	// 生产者
	for begin <= end {
		ch <- begin
		begin++
	}
	close(ch)
	wg.Wait()
}

func download(host, saveDir string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for id := range ch {
		url := host + strconv.Itoa(id) + ".html"
		//fmt.Println(url)
		getContent(url, saveDir)
	}
}

// StartWork 开始工作 begin起始页 dir保存目录
func StartWork(begin, end int, host, saveDir string) {
	var url string
	file.CreateDateDir(saveDir)
	fmt.Println(saveDir)
	for begin <= end {
		url = host + strconv.Itoa(begin) + ".html"
		//fmt.Println(url)
		getContent(url, saveDir)
		begin++
		//time.Sleep(time.Second * 1) // 限制爬虫的速率
	}
}

func getContent(url, savePath string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	title := doc.Find("h1").Text()
	if strings.Contains(title, "章") {
		fmt.Println("get " + title)
		bread, _ := doc.Find("#htmlContent").Html()
		file.AppendContent(savePath+"/"+title+".html", "<div id=\"htmlContent\" class=\"contentbox clear\">")
		file.AppendContent(savePath+"/"+title+".html", bread)
		file.AppendContent(savePath+"/"+title+".html", "</div>")
		javascript := "<script type=\"text/javascript\">\n\tvar odiv=document.getElementById('htmlContent')\n\tvar aDiv=odiv.getElementsByTagName('p')\n\n    //var aDiv = document.getElementsByTagName('p');\n    var arr = [];\n    for(var i=0;i<aDiv.length;i++)\n    {\n        arr.push(aDiv[i]);\n    }\n    arr.sort(function(a,b){return a.getAttribute('data-id') - b.getAttribute('data-id')});\n\t\n    for(var i=0;i<arr.length;i++)\n    {\n\n\nodiv.appendChild(arr[i]);\n\t\n    }\n\n</script>"
		file.AppendContent(savePath+"/"+title+".html", javascript)
	}
	return
}
