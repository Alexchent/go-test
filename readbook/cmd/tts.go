/*
Copyright © 2024 chentao02 <chentao02@zenmen.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// ttsCmd represents the tts command
var ttsCmd = &cobra.Command{
	Use:   "tts",
	Short: "文本转语音服务",
	Long:  `文本转语音服务`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(input); os.IsNotExist(err) {
			fmt.Println("目标文件不存在，请传入正确的file")
			return
		}
		npmCPU := runtime.NumCPU()
		walker := npmCPU / 4 * 3
		ch := make(chan string, walker)
		wg.Add(walker)
		for i := 1; i <= walker; i++ {
			go GenVoice(ch)
		}
		color.Blue("启动%d个walker", npmCPU)
		err := filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
			ext := filepath.Ext(path)
			if ext != ".txt" {
				return err
			}
			ch <- path
			return err
		})
		if err != nil {
			color.Red("目录遍历事变" + err.Error())
			return
		}
		close(ch) // 关闭通过，for消费channel会在消费完后安全退出
		wg.Wait() // 等待消费goroutine退出
		color.Red("finish")
		return
	},
}

func init() {
	ttsCmd.Flags().StringVarP(&input, "input", "i", "", "扫描的目录")
	rootCmd.AddCommand(ttsCmd)
}

func GenVoice(ch <-chan string) {
	for file := range ch {
		fmt.Println("开始执行", fmt.Sprintf(`say -o %s.wav --data-format=alaw -f %s`, file, file))
		out := fmt.Sprintf(`%s.wav`, file)
		cmd := exec.Command("say", "-o", out, "--data-format=alaw", "-f", file)
		_, err := cmd.CombinedOutput()
		if err != nil {
			color.Red("combined out err=" + err.Error())
		}
		color.Green("%s:生成完毕", out)
	}
	wg.Done()
}
