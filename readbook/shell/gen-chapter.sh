#!/bin/bash

file=$1
echo 指定文本文本 $1
echo 输出的文件 $2
voice_folder=storage

[ ! -f $file ] && { echo "file $file does not exist"; exit 1; }

[ ! -d $voice_folder ] && { echo "folder $voice_folder does not exist"; exit 1; }

cat $file | iconv -f GBK -t UTF-8 | grep "\d章" -n >> ${voice_folder}/$2

# cat 倚天收美行.txt | iconv -f GBK -t UTF-8 | grep "\d章" -n >> 倚天收美行.chapter

# cat test.txt | iconv -f GBK -t UTF-8 | grep "\d章" | tee yt2.txt