## yutil

golang 工具包

#### 使用

`go get -v github.com/lhlyu/yutil`

#### TODO

- String
- SQL 
- Random
- 加密
- 文件
- 时间

#### 工具包全局设置

```go
// 是否不忽略错误,调用将会打印错误
func NotIgnore()
```

#### 常量定义

```go
const (
	DEFAULT_DATE_FORMAT = "2006-01-02"
	DEFAULT_TIME_FORMAT = "2006-01-02 15:04:05"
	ALPHA = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
)
```

#### 更新日志

- [v1.0.0](./changelogs/v1.0.0.md) 

```text
1. base64   
2. ip       
3. json     
4. 文件     
5. 切片      
6. 类型转换
```

- [v1.0.1](./changelogs/v1.0.1.md) 

```text
1. 常量定义
2. 加密
3. 随机
4. 文件[+]
5. 时间
```

