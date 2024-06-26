/*
Copyright © 2024 chentao02 <chentao02@zenmen.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "txt文本切割",
	Long:  `txt文本切割`,
	Run: func(cmd *cobra.Command, args []string) {
		err := SplitFile(file, 5, 1, 20)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("finished")
	},
}

func init() {
	splitCmd.Flags().StringVarP(&way, "way", "w", "line", "切割文件的方式，选择（line|chapter）中的一种")
	rootCmd.AddCommand(splitCmd)
}

func SplitFile(f string, line int, startLine, endLine int) error {
	open, err := os.Open(f)
	defer open.Close()

	utf8Reader := transform.NewReader(open, simplifiedchinese.GBK.NewDecoder()) // 编码 GBK转为UTF-8
	fileScanner := bufio.NewScanner(utf8Reader)
	lines := 1
	var builder strings.Builder
	builder.Grow(1000000000)
	for fileScanner.Scan() {
		// 将收到的GBK内容转换成utf-8
		context := strings.TrimSpace(fileScanner.Text())
		if len(context) == 0 {
			continue
		}
		if lines < startLine {
			//fmt.Println("忽略：" + context + "\n")
			lines++
			continue
		}
		//context = strings.Replace(context, "，", "", -1)
		//context = strings.Replace(context, "。", "", -1)
		builder.WriteString(context)
		if lines%line == 0 {
			fmt.Println("==================")
			content := builder.String()
			fmt.Println(content)
			builder.Reset()
		}
		if lines >= endLine {
			return nil
		}
		lines++
	}
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

func AppendContent(filename, content string) {
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// 打开文件失败处理
		log.Fatal(err.Error())
	} else {
		_, err := fd.WriteString(content + "\n")
		if err != nil {
			return
		}
	}
}
