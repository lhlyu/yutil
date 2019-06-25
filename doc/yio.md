## Yio
```
package yfunc
```


#### 方法


- 一次性读取文件所有内容
```
ReadFileAll(filePath string) string
```


- 逐行读取文件内容
```
ReadFileLines(filePath string) []string
```


- 逐行读取文件内容，去除空行和首尾空格
```
ReadFileLinesTrim(filePath string) []string
```

