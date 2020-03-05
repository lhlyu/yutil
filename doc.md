## 文档

##### Base64

- [v2.0.0] 字符串base64编码

`yutil.Base64.Encode(data string) string`


- [v2.0.0] 对象进行base64编码

`yutil.Base64.Encode4Obj(data interface{}) string`

- [v2.0.0] base64字符串解码

`yutil.Base64.Decode(data string) string`

- [v2.0.0] base64字符串反序列化对象

`yutil.Base64.Decode2Obj(data string,v interface{})`

##### Md5

- [v2.0.0] md5加密

`yutil.Md5.Encode(data string) string`

- [v2.0.0] md5 加密附带盐值

`yutil.Md5.EncodeWithSalt(data, salt string) string`


##### Sha1

- [v2.0.0] sha1加密

`yutil.Sha1.Encode(data string) string`

- [v2.0.0] sha1 加密附带盐值

`yutil.Sha1.EncodeWithSalt(data, salt string) string`

##### Json

- [v2.0.0] 将对象转成json字符串

`yutil.Json.Marshal(v interface{}) string`

- [v2.0.0] 将字符串反序列化为对象

`yutil.Json.Unmarshal(s string,v interface{})`


##### Regexp

- [v2.0.0] 正则切割字符串

`yutil.Regexp.SplitAll(s, pattern string) []string`

- [v2.0.0] 正则替换字符串

`yutil.Regexp.ReplaceAll(s, pattern string, repl string) string`

- [v2.0.0] 正则获取符合的所有字符串

`yutil.Regexp.FindAll(s, pattern string) []string`

##### File

- [v2.0.0] 判断文件是否存在

`yutil.File.IsExists(filePath string) bool`

- [v2.0.0] 一次性读取文件所有内容

`yutil.File.ReadAll(filePath string) string`

- [v2.0.0] 按行读取

`yutil.File.ReadLines(filePath string) []string`

- [v2.0.0] 逐行读取文件内容，去除空行和首尾空格

`yutil.File.ReadLinesTrim(filePath string) []string`

- [v2.0.0] 自定义处理每一行

`yutil.File.ReadLine(filePath string, f func(line string))`


##### Time 

- [v2.0.0] 获取当前时间并格式化成 2006-01-02 15:04:05

`yutil.Time.Now() string`

- [v2.0.0] 获取当前时间并格式化成 2006-01-02

`yutil.Time.NowDate() string`

- [v2.0.0] 将2006-01-02 15:04:05字符串转换成时间

`yutil.Time.Parse(s string) time.Time`

- [v2.0.0] 将2006-01-02字符串转换成时间

`yutil.Time.DateParse(s string) time.Time`

- [v2.0.0] 打印时间间隔 ,params任意参数都会导致函数重置

`yutil.Time.Interval(params ...interface{})`


##### String 

- [v2.0.0] 模板渲染

`yutil.String.TemplateParse(tmpl string, v interface{}) string`


##### Random 

- [v2.0.0] 随机指定长度的字符串，(范围包含: 大小写数字下划线)

`yutil.Random.String(length int) string`


##### Ip

- [v2.0.0] 获取客户端Ip

`yutil.Ip.ClientIp(req *http.Request) string`

- [v2.0.0] 将Ip转成uint32

`yutil.Ip.Tolong(ipstr string) uint32`


##### Runtime

- [v2.0.0] 获取系统类型

`yutil.Runtime.GetOSName() string`

- [v2.0.0] 获取调用栈上的函数信息(函数名、文件名、调用行数)

`yutil.Runtime.CurrentFuncInfo(skip int) (funcName string, fileName string, line int)`


##### Convert

- [v2.0.0] 将任意切片转成[]interface{}

`yutil.Convert.ToSlinceInterface(slice interface{}) []interface{}`

- [v2.0.0] 将任意类型转string

`yutil.Convert.ToString(v interface{}) string`

- [v2.0.0] 将任意类型转int

`yutil.Convert.ToInt(v interface{}) int`

- [v2.0.0] 将任意类型转interface

`yutil.Convert.ToInterface(v interface{}) interface{}`


##### Slice 

- [v2.0.0] 切片洗牌

`yutil.Slice.Shuffle(v interface{})`


##### Other 

- [v2.0.0] 比较两个版本号: v1 > v2 => 1; v1 == v2 => 0; v1 < v2 => -1

> 版本号示例: v1.0 v0.01 v1.0.0.0.1 v0.0.1 ...

`yutil.Other.CompareVersion(v1, v2 string) (int, error)`
