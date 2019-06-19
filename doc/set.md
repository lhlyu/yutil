## Set

#### 方法


1. 新建一个Set
> NewSet() *Set


2. 新建一个Set并且限制Set的值类型
> NewSetTypeLimit(t string) *Set


3. 清空Set中所有的元素
> (this *Set) Clear()


4. 获取一个Set的长度
> (this *Set) Size() int


5. 添加Set中的一个（或多个）元素
> (this *Set) Add(v ...interface{})  


6. 删除Set中的一个（或多个）元素
> (this *Set) Del(v ...interface{}) 


7. 判断一个元素是否在Set中
> (this *Set) Exist(key interface{}) bool 


8. 判断两个Set是否相等，相等返回true
> (this *Set) Equals(set *Set) bool 


9. 判断set是否为空，空为true
> (this *Set) Empty() bool 


10. 并集  返回一个新的Set
> (this *Set) Union(set *Set) *Set 


11. 交集  返回一个新的Set
> (this *Set) Intersection(set *Set) *Set 


12. 差集(前者减去后者)  返回一个新的Set
> (this *Set) Difference(set *Set) *Set 


13. set 转 数组
> (this *Set) ToArray() []interface{}


14. 打印set
> (this *Set) Print()

