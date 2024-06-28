# shell 脚本


### 1. 从txt文件中提取章节信息
赋予可执行权限
```
chmod +x ./gen-chapter.sh
```

从`test.txt`文件中提取章节信息
```bash
./gen-chapter.sh ../storage/test.txt  yt.chapter
```


### 按行分割文件
截取 `test.txt` 的第591行到638行，写入到 `001.txt` 文件中
```bash
sed -n '591,638p' ytsmx.txt >> 001.txt
```

将 `test.txt` 分割，每个文件5000行，新文件以 `hi` 为前缀
```bash
split -l 5000 storage/test.txt storage/hi
```