#!/bin/bash
# 将text.txt文件中的文字，逐行转换为语音，保存到voice目录中
# 命令示例 say -o hi.wav --data-format=alaw 北京欢迎您

text="$1"
voice_folder=voice

[ ! -f $text ] && { echo "file $text does not exist"; exit 1; }

[ ! -d $voice_folder ] && { echo "folder $voice_folder does not exist"; exit 1; }

while read i
do
    say -o ${voice_folder}/${i}.wav --data-format=alaw $i
done < $text

