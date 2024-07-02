/*
Copyright © 2024 chentao02 <chentao02@zenmen.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"bufio"
	"fmt"
	myFile "github.com/alexchen/go_test/File"
	"github.com/alexchen/go_test/util"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// chapterCmd represents the getChapter command
var chapterCmd = &cobra.Command{
	Use:   "chapter",
	Short: "提取章节信息",
	Long: `提取章节信息，保存到指定的目录中

example:
 	1. go run main.go chapter -f storage/test.txt -o output`,
	Run: func(cmd *cobra.Command, args []string) {
		if !CheckFlag() {
			return
		}
		err := GetChapter(file, output)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(chapterCmd)
}

type Chapter struct {
	Start int
	End   int
	Title string
}

type Chapters map[int]Chapter

func GetChapterMap(f string) (chapters Chapters, err error) {
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

// GenderShellForSplitByChapter 生成shell脚本
func GenderShellForSplitByChapter(chapters Chapters, f, save string) error {
	fd, err := os.OpenFile(save, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	for i := 1; i < len(chapters); i++ {
		c := chapters[i]
		//cmd := fmt.Sprintf("sed -n '%d,%dp' %s | iconv -f GBK -t UTF-8 >> tmp/%s.txt", c.Start, c.End, f, c.Title)
		cmd := fmt.Sprintf("sed -n '%d,%dp' %s | iconv -f GBK -t UTF-8 >> tmp/%s.txt | say -o %s.wav --data-format=alaw -f tmp/%s.txt",
			c.Start, c.End, f, c.Title, c.Title, c.Title)
		fd.WriteString(cmd + "\n")
	}
	return err
}

// GetChapter 生成txt文件暂存章节信息
func GetChapter(f, output string) error {
	filenameOnly := myFile.GetFileNameOnly(f)
	if output != "" {
		output = output + "/" + filenameOnly + ".chapter"
	} else {
		output = filenameOnly + ".chapter"
	}
	str := util.DetectEncoding(file)
	var command string
	if str == util.GBK {
		command = `shell/gen-chapter-iconv.sh`
	} else {
		command = `shell/gen-chapter.sh`
	}
	fmt.Println(fmt.Sprintf(`shell/gen-chapter.sh %s %s`, f, output))
	cmd := exec.Command("/bin/bash", "-C", command, f, output)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	} else {
		fmt.Println(string(out))
	}
	return err
}
