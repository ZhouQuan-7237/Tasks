## 实践任务三：Go 基于 net/http 开发 C/S 架构的 demo——周全

### **一.任务目标**

开发一个简单的基于HTTP的C/S架构的Go语言Demo，实现以下功能：

1. **服务器端**：监听客户端的HTTP请求，接收客户端发送的数据，并返回响应。
2. **客户端**：向服务器发送HTTP请求，接收服务器返回的响应。

---

### 二.操作步骤

#### **1.准备工作**

##### 1.1打开 Linux 虚拟机

##### 1.2VScode远程连接虚拟机

##### 1.3宿主机Git Bash创建项目目录并进入

```bash
mkdir -p /d/Go/gopath/HTTP_demo
 cd /d/Go/gopath/HTTP_demo
```

##### **1.4Git Bash初始化Go模块**

```bash
go mod init HTTP_demo
```

##### **1.5Git Bash创建客户端代码文件**

```go
touch client.go
```

##### **1.6Git Bash编辑客户端代码文件**

```go
nano client.go
```

##### 1.7VScode创建服务器端代码目录并进入

法一：在文件浏览器中直接新建文件夹，右键“在集成终端中打开”。

法二：在终端中输入命令来创建项目目录并进入。

* 打开终端

  ```
  Ctrl + ~
  ```

* 输入指令

  ```go
  mkdir ~/go/HTTP_demo
  cd ~/go/HTTP_demo
  ```

##### 1.8VScode创建服务器端代码文件

法一：直接在文件浏览器创建`server.go`文件

法二：在终端中输入命令来创建文件

```go
touch server.go
```

#### **2.VScode编写服务器端代码**

```go
package main	 //package定义Go程序的包,main是Go语言中程序的入口包。每个Go程序必须有一个main包，且程序执行时会从main函数开始执行

import (		 //import导入其他包,在当前文件中使用
    "fmt"        //用于格式化输入输出。它可以将信息打印到控制台或接收用户输入
    "net/http"   //用于处理HTTP请求和响应。可以用它来访问网页，发送HTTP请求等
    "io"         //用于读取文件、处理I/O流。在这里，它用来读取HTTP响应的内容
)

func main() {    //func main()定义了程序的主入口函数，程序从该位置开始执行
   
    //向服务器发送GET请求
    //http.Get函数发送一个HTTP GET请求到http://localhost:8080/
    //http.Get函数返回两个值：一个是response（包含HTTP响应的详细信息），另一个是 err（错误信息，若请求失败时会非 nil）
    //HTTP响应存储在response中，同时将可能的错误存储在err变量中
    response, err := http.Get("http://localhost:8080/")
    
    //错误处理
    if err != nil { //如果http.Get返回的err不为 nil，表示发送请求时发生了错误
        fmt.Println("错误信息:", err) //打印发送请求时的错误信息
        return //终止执行main函数
    }

    // 使用defer确保在函数结束时关闭响应体
    //defer会延迟执行response.Body.Close()，直到 main函数返回时才执行，确保资源在函数结束时被正确关闭
    defer response.Body.Close()

    // 读取响应的主体内容
    //ioutil.ReadAll用来读取response.Body中的内容，返回一个字节数组（body）和可能发生的错误（err）
    //响应体的内容存储在body中，可能发生的错误存储在err变量中
    body, err := io.ReadAll(response.Body)
    
    //错误处理
    if err != nil { //如果io.ReadAll返回的err不为nil，表示在读取响应体时发生了错误
        fmt.Println("客户端的错误信息:", err)// 打印读取响应体时的错误信息
        return //终止执行main函数
    }

    // 打印服务器返回的响应内容
    fmt.Printf("Server response: %s\n", body)
}
```

#### **3.nano编写客户端代码**

```go
package main //package定义Go程序的包,main是Go语言中程序的入口包。每个Go程序必须有一个main包，且程序执行时会从main函数开始执行

import (    ////import导入其他包,在当前文件中使用
    "fmt"   //用于格式化输入输出，打印信息
    "net/http" //提供了HTTP协议的相关功能，用来处理HTTP请求和响应
)

// 处理请求的函数
//handler函数：这是一个处理HTTP请求的函数，定义了如何处理当客户端访问指定路由时的请求
//w http.ResponseWriter：w 是响应写入器，用于向客户端发送HTTP响应
//r *http.Request：r 是请求对象，包含了客户端发送的请求的详细信息，如请求方法、请求头、请求体等
func handler(w http.ResponseWriter, r *http.Request) {
   
    //当接收到客户端请求时，服务端输出"服务端已成功接收到客户端发出的请求！"
    fmt.Println("服务端已成功接收到客户端发出的请求！")
    

    //向客户端发送一个响应，内容为 "客户端已成功接收到服务端返回的响应！"
    //通过 w 向客户端写入响应。
    fmt.Fprintf(w, "客户端已成功接收到服务端返回的响应！")
}

func main() {

    //设置路由，当客户端访问 http://localhost:8080/ 时，
    //Go的HTTP服务器会调用handler函数来处理请求
    http.HandleFunc("/", handler)

    
    //启动服务器之前，输出一条显示监听的地址和端口的提示信息
    //告诉用户服务器正在localhost上的8080端口运行
    fmt.Println("Server is running on http://localhost:8080")

    // 启动HTTP服务器，监听 8080 端口
    //当客户端发起请求时，服务器会根据已设置的路由调用对应的处理函数
    //nil：表示不使用自定义的ServeMux路由器，默认使用http.DefaultServeMux
    //如果启动服务器时发生错误，输出错误信息
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("服务端的错误信息:", err)
    }
}
```

按 `Ctrl+O` 保存，按 `Enter` 确认，然后按 `Ctrl+X` 退出。

#### 4.编译和运行

##### 4.1VScode编译并运行服务器端代码

法一：先编译，后运行

```go
go build -o server server.go
./server
```

法二：编译并运行

```go
go run server.go
```

服务器将开始监听 `:8080` 端口。

##### 4.2Git Bash编译并运行服务器端代码

法一：先编译，后运行

```go
go build -o client client.go
./client
```

法二 :  编译并运行

```go
go run client.go
```

客户端将连接到服务器，并发送请求。

#### **5.测试程序**

- 服务器将接收消息并返回响应，比如：

![image-20250213225407258](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250213225407258.png)

- 客户端将收到服务器的响应，比如：

![image-20250213225532808](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250213225532808.png)

#### **6.停止程序**

在客户端和服务端输入 `Ctrl+C` 停止程序。