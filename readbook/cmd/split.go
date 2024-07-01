/*
Copyright © 2024 chentao02 <chentao02@zenmen.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"bufio"
	"fmt"
	myFile "github.com/alexchen/go_test/File"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// splitCmd represents the split command
var line int
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "txt文本切割",
	Long:  `txt文本切割`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(output); os.IsNotExist(err) {
			fmt.Println(fmt.Sprintf("目录%s不存在", output))
			return
		}
		switch way {
		case "line":
			err := SplitFile(file, output, line)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "chapter":
			fmt.Println(fmt.Sprintf("暂不支持%s分割", way))
			return
		default:
			fmt.Println(fmt.Sprintf("暂不支持%s分割", way))
			return
		}
		fmt.Println("split finished")
	},
}

func init() {
	splitCmd.Flags().StringVarP(&way, "way", "w", "line", "切割文件的方式，选择（line|chapter）中的一种")
	splitCmd.Flags().StringVarP(&file, "file", "f", "", "目标文件")
	splitCmd.Flags().StringVarP(&output, "output", "o", "output", "输出文件路径")
	splitCmd.Flags().IntVarP(&line, "line", "n", 100, "切割行数")
	// file 必传
	err := splitCmd.MarkFlagRequired("file")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rootCmd.AddCommand(splitCmd)
}

func SplitFile(f, output string, line int) error {
	open, err := os.Open(f)
	defer open.Close()

	//utf8Reader := transform.NewReader(open, simplifiedchinese.GBK.NewDecoder()) // 编码 GBK转为UTF-8
	fileScanner := bufio.NewScanner(open)

	var builder strings.Builder
	//builder.Grow(1000000)
	i, chapterNo := 1, 1
	for fileScanner.Scan() {
		context := strings.TrimSpace(fileScanner.Text())
		//if len(context) == 0 {
		//	continue
		//}
		//context = strings.Replace(context, "，", "", -1)
		//context = strings.Replace(context, "。", "", -1)
		builder.WriteString(context + "\n")
		if i%line == 0 {
			//fmt.Println("==================")
			fmt.Println(i)
			content := builder.String()
			//fmt.Println(content)
			myFile.AppendContent(output+fmt.Sprintf("/%d.txt", chapterNo), content)
			builder.Reset()
			chapterNo++
		}
		i++
	}
	content := builder.String()
	myFile.AppendContent(output+fmt.Sprintf("/%d.txt", chapterNo), content)

	fmt.Println("总计：", i)
	return err
}

func SplitFileAndTTs(f string, line int) error {
	open, err := os.Open(f)
	defer open.Close()

	utf8Reader := transform.NewReader(open, simplifiedchinese.GBK.NewDecoder()) // 编码 GBK转为UTF-8
	fileScanner := bufio.NewScanner(utf8Reader)
	lines := 1
	var builder strings.Builder
	builder.Grow(1000000000)
	for fileScanner.Scan() {
		context := strings.TrimSpace(fileScanner.Text())
		if len(context) == 0 {
			continue
		}
		if lines < 9 {
			lines++
			continue
		}
		// 将收到的GBK内容转换成utf-8
		//context = strings.Replace(context, "，", "", -1)
		//context = strings.Replace(context, "。", "", -1)
		builder.WriteString(context)
		//if lines/line == 0 {
		content := builder.String()
		//fmt.Println(content)
		Tts(content, strconv.Itoa(lines))
		builder.Reset()
		//}
		if lines >= 11 {
			return nil
		}
		lines++
	}
	return err
}

func Tts(content string, title string) {
	fmt.Println(fmt.Sprintf(`say -o %s.wav --data-format=alaw "%s"`, title, content))
	cmd := exec.Command("say", fmt.Sprintf(`-o %s.wav`, title), "--data-format=alaw", fmt.Sprintf(`"%s"`, content))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("combined out err=" + err.Error())
	} else {
		fmt.Println(string(out))
	}
}
