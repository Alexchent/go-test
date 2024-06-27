#!/bin/bash

file="$1"
echo 指定文本文本 "$1"
# echo 章节识别符"$2"

[ ! -f $file ] && { echo "file $file does not exist"; exit 1; }

cat ${file} | iconv -f GBK -t UTF-8 | grep "\d章" -n | tee ${file}.chapter

# cat 倚天收美行.txt | iconv -f GBK -t UTF-8 | grep "\d章" -n >> 倚天收美行.chapter

