#Upload-Image2Qiniu
---
## 

这是基于beego框架的图片文件上传项目，目前已完成：

+ 文件的简单上传
+ 上传成功会短时间内保存用户的上传信息
+ 上传成功会返回上传文件的公网地址

## 2016-04-19

### 
嗯...被Beego上的官方文档坑了，问题出现在[Cache模块](http://beego.me/docs/module/cache.md)中。
其中他提到了Cache模块缓存数据应该这么写：

```go
bm.Put("astaxie", 1, 10)
```

### 
坑就出现在这里，我按照这样写法，写进去的数据无法Get出来，会有error，然后我去找

```
github.com/astaxie/cache/cache_test.go
```

### 
发现他是这样进行测试的：

```go
timeoutDuration := 10 * time.Second
bm.Put("astaxie", 1, timeoutDuration)
```

### 
嗯，问题就出现在Put方法的**最后一个参数**的类型。应该是time.Duration，而不是一个int，同样我们也可以查找

```
github.com/astaxie/cache/file.go
```

### 
中的Put原型方法，可见第3个参数类型赫然写着**timeout time.Duration**

### 
后来PR了beedoc，现在没有这个问题了 :)

## 
Usage:

```
$ bee run 
```