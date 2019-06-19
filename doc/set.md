## Set
```
package ytype
```


#### 方法


- 新建一个Set
```
NewSet() *Set
```


- 新建一个Set并且限制Set的值类型(例:int string bool...)
```
NewSetTypeLimit(t string) *Set
```


- 清空Set中所有的元素
```
(this *Set) Clear()
```


- 获取一个Set的长度
```
(this *Set) Size() int
```


- 添加Set中的一个（或多个）元素
```
(this *Set) Add(v ...interface{})  
```


- 删除Set中的一个（或多个）元素
```
(this *Set) Del(v ...interface{}) 
```


- 判断一个元素是否在Set中
```
(this *Set) Exist(key interface{}) bool 
```


- 判断两个Set是否相等，相等返回true
```
(this *Set) Equals(set *Set) bool 
```


- 判断set是否为空，空为true
```
(this *Set) Empty() bool 
```


- 并集  返回一个新的Set
```
(this *Set) Union(set *Set) *Set 
```


- 交集  返回一个新的Set
```
(this *Set) Intersection(set *Set) *Set 
```


- 差集(前者减去后者)  返回一个新的Set
```
(this *Set) Difference(set *Set) *Set 
```


- set 转 数组
```
(this *Set) ToArray() []interface{}
```


- 打印set
```
(this *Set) Print()
```

