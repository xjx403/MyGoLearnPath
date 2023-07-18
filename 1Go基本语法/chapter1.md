# Go基础语法第一章
> 前言： Go基础语法在很多都是通用的，本笔记不再继续缀述，仅对于Go语法上特有的部分进行总结提炼，便于有基础的开发者快速掌握Go语言。

## 第一节 环境搭建
go语言的编译器配置教程，网上已经有非常多的教程，在此仍然不过多赘述。
1. 链接推荐 菜鸟教程：https://www.runoob.com/go/go-environment.html
2. 安装成功演示：(仅针对windows11)
   ```bash
   go help       （查看go编译器命令）
   go version     （查看go编译器版本）
   ```
go语言的开发环境目前主流有JetBrain的GoLANG（需付费），微软的VScode编译器（开源，需要自行配置编译环境和代码提示插件），以及LiteIDE（开源）。目前选择VScode作为主要开发环境。

## 第二节 hello world
当完成如上环境搭建之后，便可以开始运行你自己的第一个go 程序。
```go
package main

import "fmt"

func main() {
	fmt.Println("hello world!")
}
```
1. 第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
2. 下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
3. 下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
4. 下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
使用 fmt.Print("hello, world\n") 可以得到相同的结果。
Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
>注意，可以看到，Go语法里是使用换行符作为一句代码的结束，而不是分号`;`,也就是说，Go代码每一行末尾不需要添加分号。

## 第三节 数据类型
> 基本变量中，go支持bool int float string类型，也支持数组，指针，结构体，MAP，**切片**等类型。但是没有对`char`类型的支持。
- **int**
  > Go对int类型的支持，类似于C，范围很广。包括：byte/int8/uint8, int16/uint16, int（32位或者64位）/int32/uint32, int64/uint64。 U代表无符号，数字代表bit数。
- **float**
  > Go提供了对复数的支持。 float类型包括：float32, float64, complex64(32位实数和虚数s), complex128.
- **string**
  > Go语言默认使用UTF-8作为`string`类型的存储方式，也就是意味着，对于中文，每个字符采用三个字节存储，对于英语，每个字符采用一个字节存储。
  例子如下：
  ```go
  package main
  import "fmt"

  func main() {
    var str string = "E人"
    fmt.Println(len(str))
    str += "e"
    fmt.Printf("str: %s, len = %d\n", str, len(str))
  }
  ```
  输出如下：
  <p>4 <p\>
  <p>str: E人e, len = 5 <p>

- **指针类型**