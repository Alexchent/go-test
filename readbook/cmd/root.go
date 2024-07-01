/*
Copyright © 2024 chentao02 <chentao02@zenmen.com>
This file is part of CLI application foo.
*/
package cmd

import (
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
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "目标文件")
	rootCmd.PersistentFlags().StringVarP(&chapter, "chapter", "", "", "章节目录文件")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "output", "输出文件路径")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
