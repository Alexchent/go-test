# shell 脚本


### 1. 从txt文件中提取章节信息
从`test.txt`文件中提取章节信息
```bash
./gen-chapter.sh ../storage/test.txt  yt.chapter
```


## sed 提取文件内容
截取 `test.txt` 的第591行到638行，写入到 `001.txt` 文件中
```bash
sed -n '591,638p' ytsmx.txt >> 001.txt
```

## split 分割文件
将 `test.txt` 分割，每个文件5000行，新文件以 `hi` 为前缀
```bash
split -l 5000 storage/test.txt storage/hi
```

1. 将《穿越之修仙.txt》按行分割，新文件已 cyxx. 为前缀
```bash
split -l 1000 穿越之修仙.txt cyxx.
```

2. 将《穿越之修仙.txt》按大小分割，新文件已 cyxx. 为前缀
```bash
split -b 50k 穿越之修仙.txt cyxx.
```