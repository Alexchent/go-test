package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	file "github.com/alexchen/go_test/File"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 下载书籍
func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	down1(homeDir)
}

func down1(saveDir string) {
	begin, end := 939847, 940506
	host := "https://www.qixiashu.com/yue/465/465479/"
	dir := saveDir + "/影视实际从攻略女主开始"
	StartWork(begin, end, host, dir)
}

func down2(saveDir string) {
	begin, end := 133733, 138155
	host := "https://www.qixiashu.com/yue/62/62422/"
	dir := saveDir + "/武侠之无尽恶人"
	StartWork(begin, end, host, dir)
}

// StartWork 开始工作 begin起始页 dir保存目录
func StartWork(begin, end int, host, saveDir string) {
	var url string
	file.CreateDateDir(saveDir)
	fmt.Println(saveDir)
	for begin < end {
		url = host + strconv.Itoa(begin) + ".html"
		fmt.Println(url)
		getContent(url, saveDir)
		begin++
		//time.Sleep(time.Second * 1) // 限制爬虫的速率
	}
}

func getContent(url, savePath string) {
	res, err := http.Get(url)
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
}
