## 实践任务二：Go 基于 TCP Socket 开发 C/S 架构的 demo——周全

### **一.任务目标**

开发一个简单的 C/S 架构的 TCP Socket 程序，实现以下功能：

1. **服务器端**：监听客户端的连接请求，接收客户端发送的消息，并返回响应。
2. **客户端**：连接到服务器，发送消息，并接收服务器的响应。

---

### 二.操作步骤

#### **1.准备工作**

##### 1.1打开 Linux 虚拟机

##### 1.2VScode远程连接虚拟机

##### 1.3宿主机Git Bash创建项目目录并进入

```bash
mkdir -p /d/Go/gopath/TCP_demo
 cd /d/Go/gopath/TCP_demo
```

##### **1.4Git Bash初始化Go模块**

```bash
go mod init TCP_demo
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
  mkdir ~/go/TCP_demo
  cd ~/go/TCP_demo
  ```

##### 1.8VScode创建服务器端代码文件

法一：直接在文件浏览器创建`server.go`文件

法二：在终端中输入命令来创建文件

```go
touch server.go
```

#### **2.VScode编写服务器端代码**

```go
package main

import (
    "bufio"      // 导入用于缓冲读取的包
    "fmt"        // 导入格式化输入输出的包
    "net"        // 导入网络功能包
)

func main() {
    // net.Listen 用于创建一个 TCP 监听器，指定监听地址和端口。这里监听的是 8080 端口。
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        // 如果监听失败，输出错误并结束程序。
        fmt.Println("Error listening:", err)
        return
    }
    // 确保程序结束时关闭监听器，释放资源
    defer listener.Close()

    // 输出服务器正在运行的提示信息
    fmt.Println("Server is running on :8080...")

    // 程序进入循环，等待并处理来自客户端的连接
    for {
        // 使用 Accept 方法等待并接受客户端连接
        conn, err := listener.Accept()
        if err != nil {
            // 如果连接出错，输出错误并继续等待新的连接
            fmt.Println("Error accepting connection:", err)
            continue
        }

        // 输出连接的客户端地址，帮助服务器管理员了解连接来源
        fmt.Println("Client connected:", conn.RemoteAddr())

        // 每当接收到一个新的客户端连接，启动一个新的 goroutine 来处理这个连接
        // 这样做可以让服务器同时处理多个客户端的请求
        go handleClient(conn)
    }
}

// 这个函数用于处理与客户端的通信
func handleClient(conn net.Conn) {
    // 确保函数结束时关闭连接
    defer conn.Close()

    // 创建一个 bufio.Reader 来读取客户端发来的数据
    reader := bufio.NewReader(conn)

    // 无限循环，保持与客户端的连接并接收数据
    for {
        // 使用 ReadString 方法读取客户端发送的数据，直到遇到换行符（\n）
        message, err := reader.ReadString('\n')
        if err != nil {
            // 如果读取出错（比如客户端关闭连接），输出错误并结束处理
            fmt.Println("Error reading message:", err)
            return
        }

        // 输出接收到的消息
        fmt.Printf("Received message from client: %s", message)

        // 构造一个响应消息，表示服务器成功接收到了客户端的信息
        response := "Message received by server: " + message

        // 使用 conn.Write 向客户端发送响应消息
        // 注意：这里通过将响应消息转换为字节数组发送
        conn.Write([]byte(response + "\n"))
    }
}
```

#### **3.nano编写客户端代码**

```go
package main

import (
    "bufio"      // 导入用于缓冲读取的包
    "fmt"        // 导入格式化输入输出的包
    "net"        // 导入网络功能包
    "os"         // 导入操作系统功能包（用于读取用户输入）
)

func main() {
    // 通过 net.Dial 函数连接到服务器，指定连接的类型是 TCP，地址是 localhost:8080。
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        // 如果连接服务器出错，输出错误信息并退出程序。
        fmt.Println("Error connecting to server:", err)
        return
    }
    // 确保在函数结束时关闭连接，释放资源
    defer conn.Close()

    // 输出成功连接到服务器的提示信息
    fmt.Println("Connected to server.")

    // 创建一个 bufio.Reader 来从标准输入（用户键盘输入）读取数据
    reader := bufio.NewReader(os.Stdin)

    // 无限循环，保持与服务器的交互
    for {
        // 提示用户输入消息
        fmt.Print("Enter message: ")
        
        // 读取用户输入的字符串，直到遇到换行符（\n）
        message, _ := reader.ReadString('\n')

        // 将用户输入的消息发送到服务器。注意：消息末尾加上换行符（\n）
        conn.Write([]byte(message + "\n"))

        // 读取服务器返回的响应消息，直到遇到换行符（\n）
        response, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            // 如果读取服务器响应出错（如服务器关闭连接），输出错误信息并退出
            fmt.Println("Error reading response:", err)
            return
        }
        // 输出服务器返回的消息
        fmt.Printf("Server response: %s", response)
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

客户端将连接到服务器，并可以输入消息与服务器交互。

#### **5.测试程序**

##### **5.1在客户端输入消息**

- 输入消息，比如：`Hello World!`
- 按 `Enter` 发送消息。

##### **5.2查看服务器响应**

- 服务器将接收消息并返回响应，比如：

  ![image-20250214004729460](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250214004729460.png)

- 客户端将收到服务器的响应，比如：

![image-20250214004703898](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250214004703898.png)

#### **6.停止程序**

在客户端和服务端输入 `Ctrl+C` 停止程序。