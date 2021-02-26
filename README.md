# beego

learn beego

# 安装

go环境搭建及beego框架源码下载 参考文献：https://studygolang.com/articles/34

如果安装出现bee command not found问题 参考文献：https://www.cnblogs.com/meixudong/p/10229325.html，也可以把GOPATH/bin下的bee.exe拷贝到 GOROOT/bin里面,也就是你的go语言环境变量的bin里面,然后运行 bee

# 单元测试

#### 安装goConvey
本身自带自动化测试，但是不支持断言、mock，所以需要依赖外部包，常用 goconvey gomonkey

goConvey是作为外层框架进行单元测试管理。

1、安装
$ go get github.com/smartystreets/goconvey

2、写测试用例

3、切换至测试文件目录，运行测试用例 $ go test -v 

参考文献：https://blog.csdn.net/qq_38959696/article/details/106468440