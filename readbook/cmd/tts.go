/*
Copyright © 2024 chentao02 <chentao02@zenmen.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

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
		walker := 8
		ch := make(chan string, walker)
		for i := 1; i <= walker; i++ {
			go GenVoice(ch)
		}

		err := filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
			ext := filepath.Ext(path)
			if ext != ".txt" {
				return err
			}
			ch <- path
			return err
		})
		if err != nil {
			return
		}
	},
}

func init() {
	ttsCmd.Flags().StringVarP(&input, "input", "i", "", "扫描的目录")
	rootCmd.AddCommand(ttsCmd)
}

func GenVoice(ch chan string) {
	for file := range ch {
		fmt.Println(fmt.Sprintf(`say -o %s.wav --data-format=alaw -f %s`, file, file))
		//cmd := exec.Command("say", fmt.Sprintf(`-o %s.wav`, file), "--data-format=alaw", "-f", file)
		cmd := exec.Command("say", "-o", fmt.Sprintf(`%s.wav`, file), "--data-format=alaw", "-f", file)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("combined out err=" + err.Error())
		} else {
			fmt.Println(string(out))
		}
	}
}
