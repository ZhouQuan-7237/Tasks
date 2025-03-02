# Go入门学习笔记——周全

## 学习目录

一. Go 语言简介

1. Go基础介绍
2. Go C系演变
3. Go语言特性
4. Go语言应用

二.Go执行原理

1. Go源码文件
2. Go执行原理
3. Go执行流程
4. Go常见指令

三.Go开发环境搭建

四.Go基础语法

1. 注释
2. 数据类型、类型转换
3. 变量、常量
4. 运算符、表达式
5. 输入输出
6. 分支逻辑
7. 循环逻辑
8. 数组
9. 切片
10. Map
11.  函数
12.  指针
13.  包管理
14.  结构体
15.  方法 and 接口
16.  错误处理

## 学习内容

### 一.Go语言简介

#### 1.Go基础介绍

**Go语言**（又称**Golang**）是一种开源编程语言，其设计初衷是解决现代软件开发中的效率和可维护性问题，尤其在并发处理和垃圾回收方面表现出色。Go为开发者提供了一种简单、强大且高效的编程工具。

#### 2.Go C系演变

![在这里插入图片描述](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/7368b57596ef2e376a65bd9ca101a695.png)

Go语言是C语言家族中的重要成员，继承了C语言的许多核心特性，并在并发编程和现代软件开发方面进行了创新。以下是Go C系历史演变的总结：

- **B语言（1969年）**：C语言的前身。
- **C语言（1972年）**：现代编程语言的基石，影响了众多后续语言。
- **Newsqueak（1989年）**：引入了早期的并发编程模型。
- **Alef（1993年）**：在C语言基础上加强了并发编程。
- **Limbo（1995年）**：专注于并发编程和分布式系统，推动了并发模型的发展。
- **Go语言（2009年）**：结合C语言高效性与现代编程语言特性，加入并发支持与垃圾回收。

#### 3.Go语言特性

* **简洁性**：语法简洁，易于学习和使用。

- **高效性**：继承了C语言的高效性，编译为机器码，执行速度快。

- **并发支持**：通过Goroutines和Channels实现高效并发编程，利用CSP模型简化并发编程。

  * **Goroutines**：轻量级线程，适合并发编程。

  * **Channels**：用于Goroutines之间的通信，确保数据安全。

  * **CSP模型**：基于Communicating Sequential Processes(CSP)并发模型，简化并发任务管理。

- **垃圾回收**：自动内存管理，减少内存泄漏和悬挂指针等问题。

- **标准库强大**：提供丰富的网络和系统标准库，支持分布式和去中心化特性。

- **静态类型**：在编译时进行类型检查，减少运行时错误。

#### 4.Go语言应用

* **Web开发**：Go适用于开发高性能的Web服务器和应用。
* **云服务和微服务**：Go适合开发云服务和微服务架构。eg.Docker和Kubernetes。
* **网络编程**：Go适合开发高性能的网络应用。eg.TCP服务器
* **系统工具和命令行工具**：Go适合开发各种系统工具和命令行应用。
* **数据处理和大数据**：Go适合数据处理和开发大数据应用。
* **DevOps工具**：Go适合开发许多DevOps工具。eg.Terraform

### 二.Go执行原理

#### 1.Go源码文件

有三种类型，如下图所示：

![img](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/4db25ba25eece8e1b4e2c893eeb21f8b.png)

1. **命令源码文件**

   - **特征**：包含`package main` 和 `main()` 函数的Go文件。
   - **作用**：Go程序的入口，生成可执行文件。
   - **注意**：同一代码包中最好不要放多个命令源码文件（多个命令源码文件虽然可以分别单独 `go run` 运行起来，但是无法通过 `go build` 和 `go install`编译）。

2. **库源码文件**

   - **特征**：不包含`package main` 和 `main()` 函数的普通Go文件。
   - **作用**：提供函数或类型供其他包调用，编译为 `.a` 归档文件。

3. **测试源码文件**

   - **特征**：以 `_test.go` 结尾，包含 `TestXxx`（功能测试）和 `BenchmarkXxx`（性能测试）函数。

     * 功能测试函数：以 Test 为名称前缀，只能接受 *testing.T 的参数

       ```go
       func TestXXX( t *testing.T) {
       
       }
       ```

     * 性能测试函数：以 Benchmark 为名称前缀，只能接受 *testing.B 的参数

       ```go
       func BenchmarkXXX( b *testing.B) {
       
       }
       ```
     
   - **作用**：通过 `go test` 运行，验证代码逻辑与性能。

#### 2.Go执行原理

* 编译阶段：

  * 编写代码  (输入）：开发者编写 `.go` 文件（命令源码文件、库源码文件、测试源码文件）。
  * 编译转换（处理）：编译器先将源码解析为抽象语法树（AST），进行类型检查和语义分析，然后生成 SSA 中间代码并优化，最终转换为目标平台的机器码。
  * 生成文件（输出）：生成平台相关的可执行文件（如 `.exe`）或库文件（`.a` ）。

* 运行阶段：

  * 初始化（垃圾回收器、调度器、init()函数）
  * 执行入口。调用`main()`函数（程序的入口）

  * 并发调度：运行时调度器将 Goroutine（G）分配到逻辑处理器（P）和操作系统线程（M）上执行。

    * GMP 模型：
      - G（Goroutine）：轻量级协程，初始栈 2KB。
      - M（Machine）  ：操作系统线程，由内核调度。
      - P（Processor） ：逻辑处理器，管理本地队列和上下文切换。

    * 调度策略：工作窃取（空闲 P 偷取其他 P 的任务）、抢占式调度（防止 Goroutine 独占 CPU）。

  * 垃圾回收（GC)：使用三色标记清除算法，并发标记内存对象，减少停顿时间。

#### 3.Go执行流程

```
编写代码 → 编译代码（生成机器码）→ 生成文件 → 运行时初始化（GC、调度器、init()）→ 执行 main() → 并发调度(GMP)和垃圾回收(GC) → 程序退出  
```

#### 4.Go常见命令

![image-20250302180055782](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250302180055782.png)

1.`go build`:

* 功能：编译Go代码包/源码文件，生成可执行文件或库文件

* 用法：

  ```go
  go build 编译当前目录下的包
  (在当前目录生成一个与目录同名的可执行文件（比如:如果目录名为app，则生成app或app.exe）)
  
  go build [目录] 编译指定定目录下包
  (比如：go build ./cmd/myapp,编译位于./cmd/myapp目录下的包)
  
  go build -o [文件名] 指定输出文件名
  (比如:go build -o myapp,在当前目录生成名为myapp的可执行文件)
  ```

  

2.`go run`  

* 功能：编译并直接运行Go源码文件，适用于快速测试和运行简单的Go程序

* 用法：

  ```go
  go run [文件名] 运行单个文件
  (比如：go run main.go)
  
  go run [文件名1] [文件名2]
  (比如：go run main.go hello.go)
  ```

  

3.`go test`     

* 功能：运行Go测试文件，确保代码正确性。

* 用法：

  ```go
  go test 运行当前目录下的所有测试文件,并输出测试结果
  go test -v 详细显示测试过程
  go test -run=[函数名] 运行特定测试函数
  (比如：go test -run=TestAdd，运行TestAdd测试函数)
  ```

* 注意：

  * 测试文件必须以`_test.go`结尾，并位于与被测试代码相同的包中。

  * 测试函数必须以`Test`开头，并接收`*testing.T`类型的参数。

    

4.`go fmt`      

* 功能：格式化Go源码文件（自动调整代码中的缩进、空格、换行等格式)使代码符合Go语言的标准风格

* 用法：

  ```go
  go fmt 格式化当前目录下的所有Go文件
  
  go fmt [目录] 编译指定定目录下包
  (比如：go fmt ./cmd/myapp,格式化位于./cmd/myapp目录下的包)
  ```

  

5.`go mod`

* 功能：用于管理Go模块

* 用法：

```         go
go mod init [模块名] 初始化一个新的Go模块
(比如：go mod init github.com/username/project,
在当前目录生成一个go.mod文件，声明模块路径和Go版本)

go mod tidy 整理依赖，添加缺失的模块，移除不需要的模块
```



6.`go install`

* 功能：编译Go程序并将生成的可执行文件安装到 `$GOPATH/bin`目录下，使其能够在命令行中直接运行。

* 用法：

  ```go
  go install 编译并安装当前目录中的包
  go install [包名]编译并安装指定目录中的包(比如:go install github.com/myusername/myapp)
  ```

  

7.`go vet` 

* 功能：静态代码分析工具，用于检查代码中的潜在错误。

* 用法：

```go
go vet [文件名]
比如：go vet *.go
```



8.`go version`    

* 功能：查看当前Go语言的版本。
* 用法：

 ```GO
go version
 ```



9.`go doc`      

* 功能：查看Go程序包里的文档

* 用法：

  ```go
  go doc [包名]
  比如：go doc fmt
  ```

  

10. `go env`   

* 功能：打印Go语言的环境变量信息

* 用法：

  ```go
  go env      
  ```



11.`go clean`  

* 功能：清理编译时生成的临时文件

* 用法：

  ```go
  go clean
  ```

  

12.`go generate` 

* 功能：执行Go代码中的 `//go:generate` 指令，自动生成代码。

* 用法：

  ```go
  go generate [文件名]
  比如：go generate main.go
  ```

  

13.`go get`    

* 功能：下载并安装指定的Go包及其依赖

* 用法：

  ```go
  go get [包名] 
  比如:go get github.com/gorilla/mux
  ```

  

14.`go list` 

* 功能：列出项目中的包或模块信息

* 用法：

  ```go
  go list [包名]
  go list -m all 
  ```

  

15.`go tool`  

* 功能：调用Go的底层工具

* 用法：

  ```go
  go tool 列出所有可以调用的工具
  go tool [工具名] 调用具体的工具
  (比如：go tool vet,调用Go的静态分析工具)
  ```

  

16.`go fix`    

* 功能：用于更新Go源代码

* 用法：

  ```go
  go fix 更新当前目录的代码
  go fix [包名]更新特定目录的代码
  (比如：go fix ./myapp)
  ```



17.`go bug`

* 功能：用于启动一个 Go 语言的 bug 报告流程

* 用法：

  ```go
  go bug
  ```

  

18.`go work`

* 功能：主要用于管理和构建 Go 工作区

* 用法：

  ```go
  go work init 初始化工作区
  
  go work use [模块] 向工作区添加模块
  比如：go work use ./module1 ./module2，将 module1 和 module2 这两个模块加入到工作区
  
  go work build 构建工作区中的所有模块
  ```

  

19.`go telemetry`

* 功能：用于收集与 Go 工具链、Go 编译器以及工具的使用情况相关的数据（遥测数据）。

* 用法：

  ```GO
  go telemetry enable 启用遥测数据收集
  go telemetry disable 禁用遥测数据收集
  ```

---

### 三.Go开发环境搭建

![image-20250206022848408](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250206022848408.png)

### 四. Go基础语法

#### 1. 注释

* 单行注释：从`//`开始，直到行末

* 多行注释：从`/*`开始，`*/`结束

  ```go
  // 这是一个单行注释
  
  /*
    这是一个多行注释
    这是一个多行注释
  */
  ```

#### 2.数据类型、类型转换                                                                                              

##### 2.1数据类型

###### 2.11整型

* 有符号整型：`int8`, `int16`, `int32`, `int64`, `int`
* 无符号整型：`uint8`, `uint16`, `uint32`, `uint64`, `uint`

> 有符号整型范围：`-2^(n-1) 到 2^(n-1)-1`  
>
> 无符号整型范围: ` 0 到 2^n-1`

```go
var a int = 1
var b int64 = 1000000
var c uint = 20
var d uint8 = 200
```

###### 2.12浮点型

* float32 ： 范围 约1.4e-45 到 约3.4e38

* float64 ：范围约4.9e-324 到 约1.8e308

```go
var pi float32 = 3.14
var e float64 = 3.1415926
```

###### 2.13字符类型

* byte（uint8 ）：代表了 ASCII 码的一个字符
* rune（int32）：代表 UTF-8 的一个字符

```go
//使用单引号 表示一个字符
var ch byte = 'A'
//在 ASCII 码表中，A 的值是 65,也可以这么定义
var ch byte = 65
//65使用十六进制表示是41，所以也可以这么定义 \x 总是紧跟着长度为 2 的 16 进制数
var ch byte = '\x41'
//65的八进制表示是101，所以使用八进制定义 \后面紧跟着长度为 3 的八进制数
var ch byte = '\101'

fmt.Printf("%c",ch)
```

>在书写 UTF-8字符时，需要在 16 进制数之前加上前缀`\u`或者`\U`。如果需要使用到 4 字节，则使用`\u`前缀，如果需要使用到 8 个字节，则使用`\U`前缀。

```go
var ch rune = '\u0041'
	var ch1 int64 = '\U00000041'
	//格式化说明符%c用于表示字符，%v或%d会输出用于表示该字符的整数，%U输出格式为 U+hhhh 的字符串。
	fmt.Printf("%c,%c,%U",ch,ch1,ch)
```

###### 2.14字符串型

> 字符串是由`不可更改`的字符组成的序列。

```go
var greeting string = "Hello, World!"
```

###### 2.15布尔型

> 布尔类型数据表示真或假，使用`bool`类型声明，且不能参与任何计算以及类型转换。

```go
var isA bool = true
var isB bool = false
```

> `==`,`>`,`<`，`<=`, `>=`,`&&(AND)`,`||(OR)`等都会产生bool值

```go
var aVar = 10
aVar == 5  // false
aVar == 10 // true
aVar != 5  // true
aVar != 10 // false
```

> `&&(AND)`,`||(OR)`的短路行为

如果运算符左边的值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值。(&&优先级高于||)

```GO
var a = 10
	//因为a>11已经不满足了，所以a < 30不会走，整个表达式为false
	if(a > 11 && a < 30){
		fmt.Println("正确")
	}else{
		fmt.Println("错误")
	}

	//因为a > 5已经满足了，所以a < 30不会走，整个表达式为true
	if(a > 5 || a < 30){
		fmt.Println("正确")
	}else{
		fmt.Println("错误")
	}
```

##### 2.2类型转换

Go语言只有显式类型转换，不支持隐式类型转换。

```go
//类型 B 的值 = 类型 B(类型 A 的值)
valueOfTypeB = type B(valueOfTypeA)
```

```go
a := 5.0
b := int(a)

fmt.Println(a, b) // 输出: 5.0 5
```

#### 3.变量、常量

##### 3.1变量

###### 3.11变量声明

> 在Go中，变量使用`var`关键字声明，后跟变量名和类型。

```go
var age int//声明了一个名为age的变量，类型为int
age = 30

var name string//声明了一个名为name的变量，类型为string
name = "Alice"
```

> 也可以在声明时赋值：

```go
var age int = 30
var name string = "Alice"
```

> Go语言支持类型推断，即编译器可以根据赋值自动推断变量类型：

```go
var age = 30          // age为int
var name = "Alice"    // name为string
```

> Go语言提供了批量声明的方式

```go
var (
    a int
    b string
    c []float32
)
```

> 在函数内部，也可以使用短变量声明语法`:=`来声明并初始化变量。

```go
age := 30
name := "Alice"
isA := true
```

等同于：

```go
var age int = 30
var name string = "Alice"
var A bool = true
```

> 当一个变量被声明之后，系统自动赋予它该类型的零值：

`int 为 0`，`float 为 0.0`，`bool 为 false`，`string 为空字符串`，`指针为 nil `

> 变量的命名规则遵循驼峰命名法，即首个单词小写，每个新单词的首字母大写，例如： startDate。

###### 3.12匿名变量

> 使用多重赋值时，如果不需要在左值中接受变量，可以使用匿名变量。

```go
package main

import (
	"fmt"
	"net"
)
func main() {

    //conn, err := net.Dial("tcp", "127.0.0.1:8080")
    //如果不想接收err的值，那么可以使用_表示，这就是匿名变量
    conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println(conn)
}
```

```go
package main

import (
	"fmt"
	"net"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	//匿名变量可以重复声明
	conn1, _ := net.Dial("tcp", "127.0.0.1:8080")
	// 匿名变量不可以直接开头
	// _ :=1
	fmt.Println(conn)
	fmt.Println(conn1)
}
```

##### 3.2常量

> 常量使用`const`关键字声明，其值在编译时确定，且不可修改。

```go
const Pi = 3.14159
const Greeting = "Hello, World!"
```

> 和变量声明一样，可以批量声明多个常量

```go
const (
    e  = 2.7182818
    pi = 3.1415926
)
```

#### 4.运算符、表达式

- **算术运算符**：`+`, `-`, `*`, `/`, `%`  

- **关系运算符**：`==`, `!=`, `>`, `<`, `>=`, `<=`  

- **逻辑运算符**：`&&`, `||`, `!`  

- **赋值运算符**：`=`, `+=`, `-=`, `*=`, `/=`  

  

  - **算数运算符**：`+`, `-`, `*`, `/`, `%`

    ```go
    var sum = 10 + 5
    var product = 10 * 5
    ```

  - **关系运算符**：`==`, `!=`, `<`, `>`, `<=`, `>=`

    ```go
    var isEqual = (10 == 5)
    var isGreater = (10 > 5)
    ```

  - **逻辑运算符**：`&&`，`||`，`!`

    ```go
    var result = (10 > 5 && 5 > 3)
    ```

  - **赋值运算符**：`=`, `+=`, `-=`, `*=`, `/=`, `%=`

    ```go
    var a int = 10
    a += 5  // a = a + 5
    ```

#### 5.输入输出

> Go 提供了 `fmt` 包来进行输入输出操作。

```go
package main // package 定义包名 main 包名,声明当前文件属于main包

import "fmt" // import 引用库 fmt 库名,导入fmt包，以便使用其提供的函数

func main() {	// func 定义函数 main 函数名
    var name string
    fmt.Print("Enter your name: ") // fmt 包名 . 调用 Print 函数,并且输出定义的字符串
    fmt.Scanln(&name)
    fmt.Println("Your name is:", name) 
}
```

#### 6.分支逻辑

##### 6.1if-else语句

```go
if x > 0 {
    fmt.Println("Positive")
} else if x < 0 {
    fmt.Println("Negative")
} else {
    fmt.Println("Zero")
}
```

##### 6.2switch语句

```go
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Other day")
}
```

> Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch。

#### 7.循环逻辑

> go语言中的循环语句只支持 for 关键字。

* 传统for循环

  ```go
  for i := 0; i < 4; i++ {
      fmt.Println(i)
  }
  ```

* **无条件 for 循环**（类似于 `while`）

  ```go
  i := 0
  for i < 4 {
      fmt.Println(i)
      i++
  }
  ```

* 无限循环

  ```go
  sum := 0
  for {
      sum++
  }
  ```

#### 8.数组

> 因为数组的长度是固定的，所以在Go语言中很少直接使用数组。

> 声明和初始化

```go
//先声明，后初始化
var arr [5]int//var 数组变量名 [元素数量]Type
arr[0] = 10
arr[1] = 20

// 声明并初始化
var arr2 = [3]string{"apple", "banana", "cherry"}

// 使用省略长度
arr3 := [...]float64{1.1, 2.2, 3.3}
```

#### 9.切片

> 切片（Slice）与数组一样，也是可以容纳若干类型相同的元素的容器。

与数组不同，切片可自动扩展,不需要指定长度，可认为切片是一个动态数组。

> 每个切片值都会将数组作为其底层数据结构，我们也把这样的数组称为切片的底层数组。

> 切片是对数组的一个连续片段的引用，所以切片是一个引用类型。

这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集。

> 需要注意，终止索引标识的项`不包括在切片内(左闭右开的区间)`。

##### 9.1声明切片

```go
//name 表示切片的变量名，Type 表示切片对应的元素类型。
var name []Type
```

```go
// 声明字符串切片
var strList []string
// 声明整型切片
var numList []int
// 声明一个空切片
var numListEmpty = []int{}
// 输出3个切片
fmt.Println(strList, numList, numListEmpty)
// 输出3个切片大小
fmt.Println(len(strList), len(numList), len(numListEmpty))
// 切片判定空的结果
fmt.Println(strList == nil)
fmt.Println(numList == nil)
fmt.Println(numListEmpty == nil)

// 声明切片时，Go 会自动为切片分配一个初始容量。
此时，切片 slice 被声明但没有分配内存。也就是说，它的初始值是 nil，且没有分配容量。
```

##### 9.2切片初始化

###### 9.21字面量初始化

```go
slice := []int{1, 2, 3, 4, 5}
fmt.Println(slice) // 输出：[1 2 3 4 5]    
```

###### 9.22使用 make 函数初始化切片

`make`需要指定类型、长度和可选的容量。

```go
slice1 := make([]int, 3) // 长度为 3，容量为3，元素初始值为 0
fmt.Println(slice) // 输出：[0 0 0]

slice2 := make([]int, 5, 10) // 长度为 5，容量为 10
fmt.Println(slice2) // 输出：[0 0 0 0 0]
```

##### 9.3创建切片

###### 9.31通过数组创建

切片和数组密不可分，如果将数组理解为一栋办公楼，那么切片就是把不同的连续楼层出租给使用者，出租的过程需要选择开始楼层和结束楼层，这个过程就会生成切片。

```go
slice [开始位置 : 结束位置]
//slice：表示目标切片对象；
//开始位置：对应目标切片对象的索引；
//结束位置：对应目标切片的结束索引。
```

```go
var a = [3]int{1, 2, 3}

fmt.Println(a, a[1:2]) //输出[1 2 3] [2]
```

```go
arr := [5]int{1, 2, 3, 4, 5} // 创建一个数组
slice := arr[1:4] // 创建一个切片，包含 arr 中索引为 1 到 3 的元素
slice2 := arr[1:] // 创建一个切片，包含 arr 中索引为 1 到 末尾 的元素
slice3 := arr[:3] // 创建一个切片，包含 arr 中索引为 开始 到 2 的元素

fmt.Println(slice) // 输出：[2 3 4]
fmt.Println(slice2) // 输出：[2 3 4 5]
fmt.Println(slice3) // 输出：[1 2 3]
```

###### 9.32通过make创建

```go
make( []Type, Length, Capacity )
//Type是指切片的元素类型
//Length（长度）指的是为这个类型分配多少个元素
//Capacity（容量）为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题
```

```go
slice := make([]int, 3) // 创建一个长度为 3 的切片，容量也为3，元素初始值为 0
fmt.Println(slice) // 输出：[0 0 0]
```

##### 9.4切片的容量与长度

* 长度（Length）：切片当前包含的元素个数，使用 `len(slice)` 获取。

* 容量（Capacity)：切片底层数组的大小，使用 `cap(slice)` 获取。

  ```go
  slice := []int{1, 2, 3, 4}
  fmt.Println(len(slice)) // 输出：4，切片的长度是 4
  fmt.Println(cap(slice)) // 输出：4，切片的容量是 4
  
  slice = append(slice, 5, 6)
  fmt.Println(len(slice)) // 输出：6，切片的长度变为 6
  fmt.Println(cap(slice)) // 输出：8，容量自动扩展为 8
  ```

##### 9.5切片操作

###### 9.51添加元素

Go 提供了内置的 `append` 函数来向切片添加元素。`append` 会返回一个新的切片，如果底层数组的容量不足以容纳新的元素，它会自动扩展容量。

```go
slice := []int{1, 2, 3}
fmt.Println(slice) // 输出：[1 2 3]

slice = append(slice, 4, 5)
fmt.Println(slice) // 输出：[1 2 3 4 5]
```

###### 9.52删除元素

Go语言没有内置的删除函数，但可以通过切片操作实现。

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    index := 2 // 删除索引为2的元素，即元素3

    s = append(s[:index], s[index+1:]...)
    fmt.Println("删除元素后的切片:", s) // 输出: [1 2 4 5]
}
```

###### 9.53修改元素

直接通过索引修改切片中的元素。

```go
package main

import "fmt"

func main() {
    s := []string{"apple", "banana", "cherry"}
    s[1] = "blueberry"
    fmt.Println("修改后的切片:", s) // 输出: [apple blueberry cherry]
}
```

###### 9.54切片截取

通过切片操作可以创建子切片，指定起始和结束索引。

```go
package main

import "fmt"

func main() {
    s := []int{10, 20, 30, 40, 50}

    sub1 := s[1:4]
    fmt.Println("sub1:", sub1) // 输出: [20 30 40]

    sub2 := s[:3]
    fmt.Println("sub2:", sub2) // 输出: [10 20 30]

    sub3 := s[2:]
    fmt.Println("sub3:", sub3) // 输出: [30 40 50]
}
```

###### 9.55复制切片

Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片中，如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制。

```go
copy( destSlice, srcSlice []T) int
```

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{5, 4, 3}
copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
```

```go
package main

import "fmt"

func main() {
    src := []int{1, 2, 3, 4, 5}
    dst := make([]int, len(src))

    copy(dst, src)
    fmt.Println("源切片:", src) //输出：源切片: [1 2 3 4 5]
    fmt.Println("目标切片:", dst) //输出：目标切片: [1 2 3 4 5]
}
```

#### 10.Map

> Map是键值对的无序集合，键和值可以是不同的类型。
>
> Map 最重要的是可以通过 key 来快速检索数据，key类似于索引，指向数据的值。

##### 10.1声明Map

```go
//mapname 为 map 的变量名。
//keytype 为键类型。
//valuetype 是键对应的值类型。
var mapname map[keytype]valuetype
//[keytype] 和 valuetype 之间允许有空格。
```

```go
var capitals map[string]string
// capitals是一个键类型为string，值类型为string的Map。
//初始值为nil，需要使用make函数初始化。
```

##### 10.2Map初始化

```go
capitals = make(map[string]string)
```

##### 10.3声明并初始化Map

```go
capitals := map[string]string{
    "中国": "北京",
    "美国": "华盛顿",
    "日本": "东京",
}
```

##### 10.4Map操作

###### 10.41添加和访问元素

```go
capitals["德国"] = "柏林" // 添加元素
capital := capitals["美国"] // 访问元素
fmt.Println("美国的首都是:", capital) // 输出: 美国的首都是: 华盛顿
```

###### 10.42删除元素

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "中国": "北京",
        "美国": "华盛顿",
        "日本": "东京",
    }

    delete(capitals, "美国")
    fmt.Println("删除美国后的Map:", capitals)
}

/*输出：
删除美国后的Map: map[中国:北京 日本:东京]
*/
```

###### 10.43修改元素

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "中国": "北京",
        "美国": "华盛顿",
        "日本": "东京",
    }

    // 修改所有首都名称
    for country := range capitals {
        capitals[country] = "首都-" + capitals[country]
    }

    fmt.Println("修改后的Map:", capitals)
}
/*输出：
修改后的Map: map[中国:首都-北京 美国:首都-华盛顿 日本:首都-东京]
*/
```

完整举例：

```go
package main

import "fmt"

func main() {
    // 声明并初始化Map
    capitals := map[string]string{
        "中国": "北京",
        "美国": "华盛顿",
        "日本": "东京",
    }
    fmt.Println("原始Map:", capitals)

    // 添加元素
    capitals["德国"] = "柏林"
    fmt.Println("添加德国后的Map:", capitals)

    // 访问元素
    capital := capitals["美国"]
    fmt.Println("美国的首都是:", capital)

    // 修改元素
    capitals["日本"] = "大阪"
    fmt.Println("修改日本后的Map:", capitals)

    // 删除元素
    delete(capitals, "德国")
    fmt.Println("删除德国后的Map:", capitals)
}
/*输出：
原始Map: map[中国:北京 美国:华盛顿 日本:东京]
添加德国后的Map: map[中国:北京 美国:华盛顿 德国:柏林 日本:东京]
美国的首都是: 华盛顿
修改日本后的Map: map[中国:北京 美国:华盛顿 德国:柏林 日本:大阪]
删除德国后的Map: map[中国:北京 美国:华盛顿 日本:大阪]
*/
```

###### 10.44遍历Map

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "中国": "北京",
        "美国": "华盛顿",
        "日本": "东京",
    }

    for country, capital := range capitals {
        fmt.Printf("%s 的首都是 %s\n", country, capital)
    }
}

/*输出：
中国 的首都是 北京
美国 的首都是 华盛顿
日本 的首都是 东京
*/
```

* 仅遍历键或值

  如果只需要键或值，可以使用`_`忽略不需要的部分。

  * 仅遍历键：

    ```go
    for country := range capitals {
        fmt.Println("国家:", country)
    }
    ```

    

  * 仅遍历值：

    ```go
    for _, capital := range capitals {
        fmt.Println("首都:", capital)
    }
    ```

#### 11.函数

##### 11.1普通函数

###### 11.11定义格式

```go
func 函数名(参数列表) 返回值类型 {
    // 函数体
}
```
* 有参数有返回值

  ```go
  func add(a int, b int) int {
      return a + b
  }
  ```

* 无参数无返回值

  ```go
  func sayHello() {
      fmt.Println("Hello, Go!")
  }
  ```

###### 11.12参数

- 普通参数 
  
  ```go
  func greet(name string) {
      fmt.Println("Hello,", name)
  }
  ```
  
- 多参数  
  
  ```go
  func sum(a, b int) int {
      return a + b
  }
  ```
  

###### 11.13返回值

* 单返回值 

```go
func square(x int) int {
    return x * x
}
```

* 多返回值

```go
func test(x, y int, s string) (int, string) {
    // 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
    n := x + y          
    return n, fmt.Sprintf(s, n)
}
```

###### 11.14函数调用

* 普通调用

```go
result := add(3, 5)
```

- 忽略返回值

  ```go
  _, err := divide(10, 0)
  if err != nil {
      fmt.Println("错误:", err)
  }
  ```

##### 11.2匿名函数

> 匿名函数是指不需要定义函数名的一种函数实现方式。
>
> 匿名函数由一个不带函数名的函数声明和函数体组成。
>
> 匿名函数的优越性在于可以直接使用函数内的变量，不必声明。

###### 11.21定义格式

```go
func(参数列表)(返回参数列表){
    函数体
}
```

###### 11.22赋值给变量 

```go
square := func(x int) int {
    return x * x
}
fmt.Println(square(5))  // 输出 25
```

###### 11.23定义时调用  

```go
func() {
    fmt.Println("这是一个匿名函数")
}()  // 立即执行
```

```go
func(data int) {
    fmt.Println("hello", data)
}(100) //(100)，表示对匿名函数进行调用，传递参数为 100。
```

##### 11.3高阶函数

###### 11.31函数作为参数  

```go
func apply(f func(int) int, x int) int {
    return f(x)
}

func double(x int) int {
    return x * 2
}

result := apply(double, 10)  // 输出 20
```

###### 11.32函数作为返回值

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := multiplier(2)
fmt.Println(double(5))  // 输出 10
```

##### 11.4闭包

###### 11.41定义

闭包是函数与其引用环境的组合，可以访问外部函数的变量。  

它可以理解为一个 **“自带记忆功能的函数”**。

它就像一个小背包，函数在执行时会把外部的变量“背”在身上，

即使离开了原来的环境，也能继续使用这些变量。

###### 11.42特点

* **记住外部变量**：闭包函数可以修改外部函数的变量，这些变量会一直存在。

* **隔离作用域**：闭包内的变量是私有的，外部无法直接访问。

###### 11.43举例

假设你有一个计数器，每次调用函数就让数字+1：

- **普通函数**：就像没有记忆的人，每次都会从0开始计数。
- **闭包**：就像带着小本本的人，每次都会记住上次的数字，接着计数。

**普通函数（无法记住状态）**

```go
func counter() int {
    count := 0
    count++
    return count
}

func main() {
    fmt.Println(counter()) // 每次都是 1
    fmt.Println(counter()) // 还是 1
}
```

**闭包（记住状态）**

```go
func createCounter() func() int {
    count := 0       // 这个变量被闭包"背"在身上
    return func() int {
        count++      // 修改外部变量
        return count
    }
}

func main() {
    myCounter := createCounter()
    fmt.Println(myCounter()) // 1
    fmt.Println(myCounter()) // 2
    fmt.Println(myCounter()) // 3
}
```

```go
package main

import "fmt"

// 外部函数
func outerFunction() func(int) int {
    // 外部函数的局部变量
    sum := 0

    // 内部函数（闭包）
    innerFunction := func(x int) int {
        // 内部函数可以访问和修改外部函数的变量
        sum += x
        return sum
    }

    // 返回内部函数
    return innerFunction
}

func main() {
    // 调用外部函数，返回一个闭包
    myClosure := outerFunction()

    // 使用闭包
    fmt.Println(myClosure(1)) // 输出: 1
    fmt.Println(myClosure(2)) // 输出: 3
    fmt.Println(myClosure(3)) // 输出: 6
}

//外部函数：outerFunction 是一个外部函数，它定义了一个局部变量 sum。
//内部函数：innerFunction 是一个内部函数，它捕获了外部函数的变量 sum。
//返回内部函数：outerFunction 返回 innerFunction，使得 innerFunction 可以在外部函数执行完毕后继续访问和修改 sum。
```

#### 12.指针

###### 12.1声明指针

```go
var p *int //p是一个指向int类型的指针，初始值为nil。
```

###### 12.2获取变量的地址

```go
a := 10
p := &a
fmt.Println("a的地址:", p) // 输出: a的地址: 0xc0000140b0
```

###### 12.3解引用指针

```go
fmt.Println("p指向的值:", *p) // 输出: p指向的值: 10
```

###### 12.4修改指针指向的值

```go
*p = 20
fmt.Println("修改后的a:", a) // 输出: 修改后的a: 20
```

```go
package main

import "fmt"

func main() {
    var a int = 10
    var p *int = &a

    fmt.Println("变量a的值:", a)         
    fmt.Println("指针p的地址:", p)       
    fmt.Println("指针p指向的值:", *p)     

    *p = 30
    fmt.Println("修改后的a:", a)         
}

/*输出
变量a的值: 10
指针p的地址: 0xc0000140b0
指针p指向的值: 10
修改后的a: 30
*/
```

#### 13.包管理

>Go语言是使用包来组织源代码的，包（package）是多个 Go 源码的集合。Go语言中为我们提供了很多内置包，如 fmt、os、io 等。

> 任何源代码文件必须属于某个包，同时源码文件的第一行有效代码必须是`package pacakgeName `语句，通过该语句声明自己所在的包。

> 一般包的名称就是其源文件所在目录的名称。

> 包可以定义在很深的目录中，包名的定义是不包括目录路径的，但是包在引用时一般使用全路径引用。

##### 13.1包的导入

* 单行导入

  ```go
  import "包 1 的路径"
  import "包 2 的路径"
  ```

* 多行导入

  ```go
  import (
      "包 1 的路径"
      "包 2 的路径"
  )
  ```

##### 13.2包的导入路径

包的绝对路径就是`GOROOT/src/`或`GOPATH`后面包的存放路径，如下所示：

```go
import "lab/test" // test 包是自定义的包，其源码位于GOPATH/lab/test目录下
import "database/sql/driver"// driver 包的源码位于GOROOT/src/database/sql/driver目录下
import "database/sql"// sql 包的源码位于GOROOT/src/database/sql 目录下
```

13.3包的引用格式

###### 13.1标准引用格式

```go
package main
import "fmt"
func main() {
    fmt.Println("Hello")
}
```

###### 13.2自定义别名引用格式

```go
package main
import F "fmt"
func main() {
    F.Println("Hello") // F 为 fmt 包的别名
}
```

###### 13.3省略引用格式

```go
package main
import . "fmt"
func main() {
    Println("Hello")  //不需要加前缀 fmt.
}
```

###### 13.4匿名引用格式

> 在引用某个包时，如果只是希望执行包初始化的 init 函数，而不使用包内部的数据时，可以使用匿名引用格式，如下所示：

```go
package main
import (
    _ "database/sql"
    "fmt"
)
func main() {
    fmt.Println("Hello")
}
```

> 使用标准格式引用包，但是代码中却没有使用包，编译器会报错。如果包中有 init 初始化函数，则通过`import _ "包的路径" `这种方式引用包，仅执行包的初始化函数，即使包没有 init 初始化函数，也不会引发编译器报错。

##### 13.3Go Modules核心概念

> Go Modules为Go语言默认的依赖管理工具。

###### 13.31模块（Module）

* 一组相关包的集合，通过 `go.mod` 文件定义。  
* 模块路径（如 `github.com/user/project`）是唯一标识符。 

* `go.mod` 文件  

  - 模块的配置文件，包含依赖和版本信息。  

  - 结构示例：  
    ```go
    module github.com/yourname/myapp  // 模块路径
    
    go 1.21                          // Go 版本
    
    require (
        github.com/gorilla/mux v1.8.0  // 依赖包及版本
        golang.org/x/text v0.3.8
    )
    
    replace github.com/old/pkg => ../local/pkg  // 替换本地路径
    ```

###### 13.32常用命令

* 初始化模块 

```bash
go mod init <module-path>  # 初始化模块（生成 go.mod）
# 示例：go mod init github.com/yourname/myapp
```

* 添加/更新依赖

```bash
go get <package>@<version>  # 安装指定版本
# 示例：
go get github.com/gin-gonic/gin@v1.9.1  # 安装 v1.9.1
go get -u github.com/gorilla/mux        # 更新到最新版本
go get gorm.io/gorm						# 安装gorm
go get gorm.io/driver/mysql				# 安装mysql驱动
```

* 整理依赖  

```go
go mod tidy  # 自动添加缺失的依赖，移除未使用的依赖
```

* 查看依赖关系  

```go
go list -m all          # 列出所有直接和间接依赖
go list -u -m all       # 查看可升级的依赖
go mod graph            # 显示依赖关系图
```

#### 14.结构体

##### 14.1定义结构体

```go
type StructName struct {
    Field1 Type1
    Field2 Type2
    // ...
}

/*eg:
type Person struct {
    Name string
    Age  int
}
*/
```

##### 14.2嵌入结构体

```go
type Address struct {
    City    string
    ZipCode string
}

type Employee struct {
    Person
    Address
    Position string
}
```

##### 14.3结构体实例化

>结构体就像是一个“盒子”，用来装不同类型的东西。
>
>结构体的实例化，就是根据结构体的定义，创建一个具体的“盒子”，并且给这个盒子里面的东西赋值。

###### 14.31基本形式

```go
var ins T //T为结构体类型，ins为结构体的实例
```

```go
package main

import "fmt"

type Point struct {
	X int
	Y int
}
func main() {
    //使用.来访问结构体的成员变量,结构体成员变量的赋值方法与普通变量一致。
	var p Point
	p.X = 1
	p.Y = 2
	fmt.Printf("%v,x=%d,y=%d",p,p.X,p.Y )
}
```

```go
package main

import "fmt"

type Point struct {
	X int
	Y int
}
func main() {

	var p Point
	//p.X = 1
	//p.Y = 2
    //未初始化的字段会使用类型的零值。
	fmt.Printf("%v,x=%d,y=%d",p,p.X,p.Y )
}
```

###### 14.32new关键字

new函数返回指向新分配的零值的指针。

```go
type Player struct{
    Name string
    HealthPoint int
    MagicPoint int
}
tank := new(Player)
tank.Name = "小明"
tank.HealthPoint = 300
```

完整示例：

```go
package main

import "fmt"

// 定义结构体
type Person struct {
    Name string
    Age  int
}

func main() {
    // 使用字面量初始化
    p1 := Person{Name: "Alice", Age: 30}
    fmt.Println("p1:", p1)

    // 不指定字段名
    p2 := Person{"Bob", 25}
    fmt.Println("p2:", p2)

    // 使用new关键字
    p3 := new(Person)
    p3.Name = "Charlie"
    p3.Age = 28
    fmt.Println("p3:", *p3)

    // 部分初始化
    p4 := Person{Name: "Diana"}
    fmt.Println("p4:", p4)
}

/*
p1: {Alice 30}
p2: {Bob 25}
p3: {Charlie 28}
p4: {Diana 0}
*/
```

##### 14.4匿名结构体

匿名结构体没有类型名称，无须通过 type 关键字定义就可以直接使用。

```go
ins := struct {
    // 匿名结构体字段定义
    字段1 字段类型1
    字段2 字段类型2
    …
}{
    // 字段值初始化
    初始化字段1: 字段1的值,
    初始化字段2: 字段2的值,
    …
}
```

```go
package main
import (
	"fmt"
)
// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n, msg:%v", msg,msg)
}
func main() {
	// 实例化一个匿名结构体
	msg := &struct {  // 定义部分
		id   int
		data string
	}{  // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}
```

#### 15.方法 and 接口

##### 15.1方法

> 方法就像是一个“指令”，告诉某个东西（比如一个变量或者一个结构体）要“做什么”。在编程里，我们通过定义方法，可以让数据或者对象完成一些特定的动作，比如打印信息、计算结果或者修改数据。

###### 15.11方法的定义与调用

> 方法的定义与函数类似，不同之处在于方法有一个接收器（Receiver），用于指定该方法属于哪个类型。

```go
func (接收器变量 接收器类型) 方法名(参数列表) (返回值类型) {
    方法体
}
```

假设你有一个 **“汽车”** 类型：

- **属性**：颜色、品牌、速度（这些是汽车的“状态”）。
- **方法**：启动、加速、刹车（这些是汽车的“行为”）。

方法就是让汽车这个类型能够执行某些操作，比如启动、加速等。

举例：

##### 1. 定义结构体（类型）

```go
type Car struct {
    Brand string
    Speed int
}
```

这里定义了一个 `Car` 类型，它有 `Brand`（品牌）和 `Speed`（速度）两个属性。

**2. 定义方法（行为）**

```go
// 定义一个方法：启动汽车
func (c Car) Start() {
    fmt.Println(c.Brand, "启动了！")
}

// 定义一个方法：加速汽车
func (c *Car) Accelerate(amount int) {
    c.Speed += amount
    fmt.Println(c.Brand, "加速到", c.Speed, "km/h")
}
```

- `Start()` 是 `Car` 的一个方法，表示启动汽车。
- `Accelerate(amount int)` 是另一个方法，表示加速汽车，并且可以修改汽车的速度。

#### 3. **使用方法和结构体**

```go
func main() {
    myCar := Car{Brand: "Tesla", Speed: 0} // 创建一辆汽车
    myCar.Start()                          // 调用方法：启动
    myCar.Accelerate(50)                   // 调用方法：加速
}
```

#### 4.输出结果

```go
Tesla 启动了！
Tesla 加速到 50 km/h
```



```go
package main

import "fmt"

// 定义结构体
type Circle struct {
    Radius float64
}

// 定义方法计算面积
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// 定义方法计算周长
func (c Circle) Circumference() float64 {
    return 2 * 3.14159 * c.Radius
}

func main() {
    circle := Circle{Radius: 5}
    fmt.Printf("面积: %.2f\n", circle.Area())               // 输出: 面积: 78.54
    fmt.Printf("周长: %.2f\n", circle.Circumference())     // 输出: 周长: 31.42
}

//Circle结构体定义了圆的半径。
//Area和Circumference方法分别计算圆的面积和周长。
```

###### 15.12方法的接收器

方法的接收者可以是值类型或指针类型，不同类型的接收者具有不同的行为和性能影响。

* 值接收器（Value Receiver）：

  - 接收器为值类型时，方法接收的是类型的副本。

  - 在方法中对接收者的修改不会影响原始变量。

  - 适用于不需要修改接收者数据的方法。

* 指针接收器（Pointer Receiver）：

  - 接收者为指针类型时，方法接收的是类型的地址。

  - 可以在方法中修改接收者的字段，影响原始变量。
  
  * 避免大结构体的复制，提高性能。
  
  * 适用于需要修改接收者数据的方法。


  **示例：**

  ```go
  package main
  
  import "fmt"
  
  // 定义结构体
  type BankAccount struct {
      Owner string
      Balance float64
  }
  
  // 值接收器方法
  func (ba BankAccount) DepositValue(amount float64) {
      ba.Balance += amount
      fmt.Printf("[DepositValue] 新余额: %.2f\n", ba.Balance)
  }
  
  // 指针接收器方法
  func (ba *BankAccount) DepositPointer(amount float64) {
      ba.Balance += amount
      fmt.Printf("[DepositPointer] 新余额: %.2f\n", ba.Balance)
  }
  
  func main() {
      account := BankAccount{Owner: "John Doe", Balance: 1000}
  
      account.DepositValue(500)     // DepositValue方法接收器为值类型，仅修改方法内部的副本，原始余额不变。
      fmt.Printf("余额 after DepositValue: %.2f\n", account.Balance) // 输出: 1000.00
  
      account.DepositPointer(500)   // DepositPointer方法接收器为指针类型，修改的是原始余额修改的是原始变量
      fmt.Printf("余额 after DepositPointer: %.2f\n", account.Balance) // 输出: 1500.00
  }
  
  /*输出：
  [DepositValue] 新余额: 1500.00
  余额 after DepositValue: 1000.00
  [DepositPointer] 新余额: 1500.00
  余额 after DepositPointer: 1500.00
  */
  ```

##### 15.2接口

###### 15.21接口定义

> 接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么。

> 接口也是一个“任务清单”。它告诉一个对象（比如一个结构体）：“你必须实现这些功能，但具体怎么实现，你自己决定。”

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
//接口类型名：使用 type 将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫 Writer，有字符串功能的接口叫 Stringer，有关闭功能的接口叫 Closer 等。
//方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
//参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以被忽略
```

```go
type Animal interface {
    MakeSound() string  // 动物必须会叫
    Run() string        // 动物必须会跑
}
//Animal 是一个接口，它要求任何实现它的类型都必须实现 MakeSound() 和 Run() 这两个方法。
```



###### 15.22接口实现条件

* 接口的方法与实现接口的类型方法格式一致

  > 在类型中添加与接口签名一致的方法就可以实现该方法。
  >
  > 签名包括方法中的名称、参数列表、返回参数列表。
  >
  > 也就是说，只要实现接口类型中的方法的名称、参数列表、返回参数列表中的任意一项与接口要实现的方法不一致，那么接口的这个方法就不会被实现。

* 接口中所有方法均被实现

  |当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用。

######  15.23类型与接口的关系

> 在Go语言中类型和接口之间有一对多和多对一的关系

* 一个类型可以实现多个接口

一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。

例如，狗可以叫，也可以动。

我们就分别定义Sayer接口和Mover接口，如下： 

~~~go
// 定义 Sayer 接口
type Sayer interface {
    say()
}

// 定义 Mover 接口
type Mover interface {
    move()
}
~~~

dog既可以实现Sayer接口，也可以实现Mover接口。

~~~go
// 定义 dog 结构体
type dog struct {
    name string
}

// 实现Sayer接口
func (d dog) say() {
    fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
    fmt.Printf("%s会动\n", d.name)
}

// 主函数逻辑
func main() {
    // 声明接口变量
    var x Sayer
    var y Mover

    // 创建一个 dog 实例
    var a = dog{name: "旺财"}

    // 将 dog 实例赋值给接口变量
    x = a  // 因为 dog 实现了 Sayer 接口
    y = a  // 因为 dog 实现了 Mover 接口

    // 通过接口变量调用方法
    x.say()  // 输出：旺财会叫汪汪汪
    y.move() // 输出：旺财会动
}
~~~

* 多个类型实现同一接口

Go语言中不同的类型还可以实现同一接口 首先我们定义一个Mover接口，它要求必须有一个move方法。

~~~go
// Mover 接口
type Mover interface {
    move()
}
~~~

例如狗可以动，汽车也可以动，可以使用如下代码实现这个关系：

~~~go
type dog struct {
    name string
}

type car struct {
    brand string
}

// dog类型实现Mover接口
func (d dog) move() {
    fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
    fmt.Printf("%s速度70迈\n", c.brand)
}
~~~

这个时候我们在代码中就可以把狗和汽车当成一个会动的物体来处理了，不再需要关注它们具体是什么，只需要调用它们的move方法就可以了。

~~~go
func main() {
    var x Mover
    var a = dog{name: "旺财"}
    var b = car{brand: "保时捷"}
    x = a
    x.move()
    x = b
    x.move()
}
~~~

并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

~~~go
// WashingMachine 洗衣机
type WashingMachine interface {
    wash()
    dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
    fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
    dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
    fmt.Println("洗刷刷")
}
~~~

###### 15.24接口嵌套

接口与接口间可以通过嵌套创造出新的接口

```go
// Sayer 接口
type Sayer interface {
    say()
}

// Mover 接口
type Mover interface {
    move()
}

// 接口嵌套
type animal interface {
    Sayer
    Mover
}
```

嵌套得到的接口的使用与普通接口一样，这里我们让cat实现animal接口：

~~~go
type cat struct {
    name string
}

func (c cat) say() {
    fmt.Println("喵喵喵")
}

func (c cat) move() {
    fmt.Println("猫会动")
}

func main() {
    var x animal
    x = cat{name: "花花"}
    x.move()
    x.say()
}
~~~

###### 15.25空接口

空接口是指没有定义任何方法的接口,因此任何类型都实现了空接口。

空接口类型的变量可以存储任意类型的变量。

~~~go
func main() {
    // 定义一个空接口x
    var x interface{}
    s := "你好"
    x = s
    fmt.Printf("type:%T value:%v\n", x, x) //type:string value:你好
    i := 100
    x = i
    fmt.Printf("type:%T value:%v\n", x, x) //type:int value:100
    b := true
    x = b
    fmt.Printf("type:%T value:%v\n", x, x) //type:bool value:true
}
~~~

###### 15.26类型断言

空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

**接口值**

一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。

这两部分分别称为`接口的动态类型`和`动态值`。

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

~~~go
x.(T)
//x：表示类型为interface{}的变量
//T：表示断言x可能是的类型。
~~~

该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

##### 类型断言

类型断言类似于 “拆盲盒”。假设你有一个接口类型的变量，但不确定它实际装的是什么具体类型（比如可能是 `Dog`、`Cat` 或其他），类型断言就是用来安全地拆开检查，确认具体类型并取出值

   ~~~go
func main() {
    var x interface{}
    x = "你好"
    v, ok := x.(string)
    if ok {
        fmt.Println(v)
    } else {
        fmt.Println("类型断言失败")
    }
}
   ~~~

上面的示例中如果要断言多次就需要写多个if判断，这个时候我们可以使用switch语句来实现：

~~~go
func justifyType(x interface{}) {
    switch v := x.(type) {
    case string:
        fmt.Printf("x is a string，value is %v\n", v)
    case int:
        fmt.Printf("x is a int is %v\n", v)
    case bool:
        fmt.Printf("x is a bool is %v\n", v)
    default:
        fmt.Println("unsupport type！")
    }
}
~~~

#### 16.错误处理

##### 16.1错误类型

在Go语言中，错误被视为一种普通的值。错误处理的核心是`error`接口，它定义了错误的基本行为。此外，Go还支持创建自定义错误类型，以满足更复杂的错误处理需求。

###### 16.11error 接口

`error`是Go语言内置的接口，用于表示错误状态。它定义了一个单一的方法：

```go
type error interface {
    Error() string
}
```

任何实现了`Error() string`方法的类型都满足`error`接口，可以被视为一个错误。

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("错误:", err)
    } else {
        fmt.Println("结果:", result) // 输出: 结果: 5
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("错误:", err) // 输出: 错误: 除数不能为零
    } else {
        fmt.Println("结果:", result)
    }
}

/*输出:
结果: 5
错误: 除数不能为零
*/
```

- `divide`函数尝试执行除法运算。如果除数为零，返回一个错误；否则，返回结果和`nil`错误。
- 主函数调用`divide`并根据返回的错误值进行相应的处理。

###### 16.12自定义错误

在某些情况下，内置的`errors.New`可能不足以满足复杂的错误处理需求。Go允许开发者创建自定义错误类型，以包含更多的上下文信息或实现特定的行为。

```go
package main

import (
    "fmt"
)

// 定义自定义错误类型
type MyError struct {
    Code    int
    Message string
}

// 实现 error 接口
func (e *MyError) Error() string {
    return fmt.Sprintf("错误代码 %d: %s", e.Code, e.Message)
}

func doSomething(flag bool) error {
    if flag {
        return &MyError{Code: 404, Message: "资源未找到"}
    }
    return nil
}

func main() {
    err := doSomething(true)
    if err != nil {
        fmt.Println("发生错误:", err) // 输出: 发生错误: 错误代码 404: 资源未找到
    } else {
        fmt.Println("操作成功")
    }
}
/*输出:
发生错误: 错误代码 404: 资源未找到
*/

//MyError结构体包含了错误代码和错误消息。
//实现了Error() string方法，使MyError满足error接口。
//doSomething函数根据传入的标志返回不同类型的错误。
```

* 使用fmt.Errorf创建带格式的错误：

Go 1.13引入了`fmt.Errorf`，支持创建带格式的错误消息，并且可以使用`%w`动词包装错误，以实现错误的链式包装。

```go
package main

import (
    "errors"
    "fmt"
)

func readFile(filename string) error {
    // 模拟一个错误
    return fmt.Errorf("无法读取文件 %s: %w", filename, errors.New("文件不存在"))
}

func main() {
    err := readFile("data.txt")
    if err != nil {
        fmt.Println("错误:", err) // 输出: 错误: 无法读取文件 data.txt: 文件不存在

        // 使用 errors.Is 检查错误
        if errors.Is(err, errors.New("文件不存在")) {
            fmt.Println("确认错误是文件不存在")
        } else {
            fmt.Println("错误类型未知")
        }
    }
}
/*输出:
错误: 无法读取文件 data.txt: 文件不存在
确认错误是文件不存在
*/
//fmt.Errorf使用%w动词包装了一个基础错误，使得错误可以被进一步检查和处理。
//errors.Is用于检查错误链中是否包含特定的错误。
```

## 参考资料

[《Go语言圣经》](https://golang-china.github.io/gopl-zh/index.html)

[《The Way To Go》](https://github.com/unknwon/the-way-to-go_ZH_CN)

[《GO语言基础 一.为什么我要学习Golang以及GO语言入门普及》](https://blog.csdn.net/Eastmount/article/details/111150449?ops_request_misc=%257B%2522request%255Fid%2522%253A%252252ba98e08a1deaa6a21e52c0cdb2b4f7%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=52ba98e08a1deaa6a21e52c0cdb2b4f7&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-1-111150449-null-null.142^v101^pc_search_result_base1&utm_term=go%E8%AF%AD%E8%A8%80&spm=1018.2226.3001.4187)

[《Go语言的前世今生》](https://zhuanlan.zhihu.com/p/322560760?utm_psn=1869192976024080384)

[安装虚拟机（VMware）保姆级教程（附安装包)](https://blog.csdn.net/weixin_74195551/article/details/127288338?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522f37757fca3048225234a7812192f203d%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=f37757fca3048225234a7812192f203d&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_positive~default-1-127288338-null-null.142^v101^pc_search_result_base1&utm_term=%E8%99%9A%E6%8B%9F%E6%9C%BA&spm=1018.2226.3001.4187)

[基于VMware虚拟机的Ubuntu22.04系统安装和配置（新手保姆级教程)](https://blog.csdn.net/qq_42417071/article/details/136327674?ops_request_misc=%257B%2522request%255Fid%2522%253A%25226c1443d902dd59a3aa4ac5e27b55e230%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=6c1443d902dd59a3aa4ac5e27b55e230&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-2-136327674-null-null.142^v101^pc_search_result_base1&utm_term=%E8%99%9A%E6%8B%9F%E6%9C%BA&spm=1018.2226.3001.4187)

[（新）VMware虚拟机安装Linux教程(超详细)](https://blog.csdn.net/weixin_52799373/article/details/140797573?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522b1a40439f9c655444f3861aedb62729d%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=b1a40439f9c655444f3861aedb62729d&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-2-140797573-null-null.142^v101^pc_search_result_base1&utm_term=%E8%99%9A%E6%8B%9F%E6%9C%BA%E5%AE%89%E8%A3%85linux&spm=1018.2226.3001.4187)

[超详细的VMware虚拟机安装Linux图文教程保姆级](https://blog.csdn.net/weixin_61536532/article/details/129778310?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522b1a40439f9c655444f3861aedb62729d%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=b1a40439f9c655444f3861aedb62729d&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-3-129778310-null-null.142^v101^pc_search_result_base1&utm_term=%E8%99%9A%E6%8B%9F%E6%9C%BA%E5%AE%89%E8%A3%85linux&spm=1018.2226.3001.4187)

 [【Go】Go语言介绍与开发环境搭建](https://blog.csdn.net/littlefun591/article/details/142093555?ops_request_misc=%257B%2522request%255Fid%2522%253A%25226b1b51e832639ed3537ad03b98c9751e%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=6b1b51e832639ed3537ad03b98c9751e&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-1-142093555-null-null.142^v101^pc_search_result_base1&utm_term=Go%20%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA&spm=1018.2226.3001.4187)

 [Go语言开发环境搭建详细教程+go常见bug合集【Go语言教程】](https://blog.csdn.net/weixin_45565886/article/details/130175889?ops_request_misc=%257B%2522request%255Fid%2522%253A%25226b1b51e832639ed3537ad03b98c9751e%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=6b1b51e832639ed3537ad03b98c9751e&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-3-130175889-null-null.142^v101^pc_search_result_base1&utm_term=Go%20%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA&spm=1018.2226.3001.4187)

  [Linux下搭建Go开发环境](https://blog.csdn.net/weixin_44617651/article/details/130280299?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522cb1fb53d717ff288441ba27bee96ddac%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=cb1fb53d717ff288441ba27bee96ddac&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-5-130280299-null-null.142^v101^pc_search_result_base1&utm_term=linux%E4%B8%8Bgo%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA&spm=1018.2226.3001.4187)

 [2.3 在 Linux 上安装 Go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/02.3.md)

 [Go语言执行原理以及Go命令](https://blog.csdn.net/qq_37003559/article/details/120766452?ops_request_misc=%257B%2522request%255Fid%2522%253A%252204534a63e7b7eae2f4dd4df4a759bb8e%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=04534a63e7b7eae2f4dd4df4a759bb8e&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-1-120766452-null-null.142^v101^pc_search_result_base1&utm_term=Go%20%E7%9A%84%E6%89%A7%E8%A1%8C%E5%8E%9F%E7%90%86&spm=1018.2226.3001.4187)

[Go的执行原理以及Go的命令](https://juejin.cn/post/7268947120799039542)

 [Go的执行原理以及Go的命令](https://www.tsyvps.com/helparticle/3214.html)

[Go语言基础语法](https://blog.csdn.net/qq_30614345/article/details/131214626?ops_request_misc=%257B%2522request%255Fid%2522%253A%25223cd73ef63ab7c9ea71b87c1ed052ebd5%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=3cd73ef63ab7c9ea71b87c1ed052ebd5&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_click~default-1-131214626-null-null.142^v101^pc_search_result_base1&utm_term=go%E8%AF%AD%E8%A8%80%E5%9F%BA%E7%A1%80%E8%AF%AD%E6%B3%95&spm=1018.2226.3001.4187)

[Go语言学习笔记—golang基础语法](https://blog.csdn.net/qq_39280718/article/details/125333519?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522ca35d8804ad9571179a4eb46c5e3aa94%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=ca35d8804ad9571179a4eb46c5e3aa94&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-7-125333519-null-null.142^v101^pc_search_result_base1&utm_term=go%E8%AF%AD%E8%A8%80%E5%9F%BA%E7%A1%80%E8%AF%AD%E6%B3%95&spm=1018.2226.3001.4187)

[前景-Go 语言中文文档](https://www.topgoer.com/)