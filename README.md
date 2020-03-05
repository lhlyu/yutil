## yutil

golang 工具包

#### 当前版本

`v2.0.0`

- 注: 不兼容v1

#### 使用

`go get -v github.com/lhlyu/yutil/v2`

#### 文档

[链接](./doc.md)

#### 错误处理

```text
默认情况是不会打印错误信息,如果需要打印，调用下面这个方法:

yutil.NotIgnore()

错误将打印在控制台

自定义错误处理

yutil.SetHandler(f func(fileName, funcName string, err error))

```