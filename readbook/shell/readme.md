# 常用shell命令

## split 分割文件

1. 将《穿越之修仙.txt》按行分割，新文件已 cyxx. 为前缀
```
split -l 1000 穿越之修仙.txt cyxx.
```

2. 将《穿越之修仙.txt》按大小分割，新文件已 cyxx. 为前缀
```
split -b 50k 穿越之修仙.txt cyxx.
```