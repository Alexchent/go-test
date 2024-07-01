/*
Copyright © 2024 chentao02 <chentao02@zenmen.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

// getChapterCmd represents the getChapter command
var getChapterCmd = &cobra.Command{
	Use:   "getChapter",
	Short: "提取章节信息",
	Long: `提取章节信息

example:
 	1. go run main.go getChapter -o gg.sh --chapter shell/storage/yt.chapter -f shell/storage/test.txt `,
	Run: func(cmd *cobra.Command, args []string) {
		ch, err := GetChapter(chapter)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = GenderChapter(ch, file, output)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(getChapterCmd)
}

type Chapter struct {
	Start int
	End   int
	Title string
}

type Chapters map[int]Chapter

func GetChapter(f string) (chapters Chapters, err error) {
	chapters = make(Chapters)
	open, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer open.Close()

	fileScanner := bufio.NewScanner(open)
	var line int
	for i := 1; fileScanner.Scan(); i++ {
		context := strings.TrimSpace(fileScanner.Text())
		ss := strings.Split(context, ":")
		if len(ss) < 2 {
			fmt.Println(context)
			return
		}
		line, _ = strconv.Atoi(ss[0])
		title := strings.Replace(ss[1], " ", "-", -1)
		//title = strings.Replace(title, "（", "", -1)
		//title = strings.Replace(title, "）", "", -1)
		chapters[i] = Chapter{Start: line, Title: title}
		last, ok := chapters[i-1]
		if ok {
			last.End = line - 1
			chapters[i-1] = last
		}
	}
	return chapters, err
}

// GenderChapter 生成shell脚本
func GenderChapter(chapters Chapters, f, save string) error {
	fd, err := os.OpenFile(save, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	for _, c := range chapters {
		cmd := fmt.Sprintf("sed -n '%d,%dp' %s | iconv -f GBK -t UTF-8 >> tmp/%s.txt", c.Start, c.End, f, c.Title)
		fd.WriteString(cmd + "\n")
	}
	return err
}
