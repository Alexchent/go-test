/*
Copyright © 2024 chentao02 <chentao02@zenmen.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	file    string
	chapter string
	output  string
	way     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "read",
	Short: "文本转语音服务",
	Long: `将txt文件切割并转化为wav格式的语音文件

利用了MacOS系统的say命令 
	say -o hi.wav --data-format=alaw 北京欢迎您
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 全局flag
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "目标文件，必传")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "输出文件路径")
	//err := rootCmd.MarkFlagRequired("file")
	//fmt.Println("root init")
	//if err != nil {
	//	fmt.Println("root err" + err.Error())
	//	return
	//}
}

func CheckFlag() bool {
	if output != "" {
		if _, err := os.Stat(output); os.IsNotExist(err) {
			fmt.Println(fmt.Sprintf("目录%s不存在，请指定正确的output", output))
			return false
		}
	}
	if file == "" {
		fmt.Println("file必传")
		return false
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("目标文件不存在，请传入正确的file")
		return false
	}
	return true
}
