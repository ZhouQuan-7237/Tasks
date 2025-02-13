package main //package定义Go程序的包,main是Go语言中程序的入口包。每个Go程序必须有一个main包，且程序执行时会从main函数开始执行

import (	 //import导入其他包,在当前文件中使用
    "bufio"  //提供缓冲读写功能，用于高效的读取输入和输出
    "fmt"    //用于格式化输入输出，打印信息到控制台
    "net"    //用于实现网络通信功能，建立连接和传输数据
)

func main() {

    //使用net.Listen创建一个TCP监听器，监听本地的 8080 端口，等待客户端的连接请求
	// listener是监听器对象，err是错误信息
    listener, err := net.Listen("tcp", ":8080")
    
	//错误处理
	if err != nil { //如果监听端口时发生错误
        fmt.Println("监听端口的错误信息:", err) //打印错误信息
        return	//退出程序
    }
    
	//使用defer确保在函数结束时关闭响应体
    //defer会延迟执行listener.Close()，直到main函数返回时才执行，确保资源在函数结束时被正确关闭
	defer listener.Close()  
	
	//打印服务器启动成功的信息
    fmt.Println("Server is running on :8080...")
	
	//启动一个无限循环，服务器会持续等待和处理客户端的连接请求
	//这样，服务器可以一直运行，不会自动退出
    for {  

		// listener.Accept() 会阻塞，直到有客户端连接到服务器
		// 当有客户端连接时，返回一个连接对象conn和一个可能的错误err
        conn, err := listener.Accept()  
        
		//错误处理
		if err != nil {	//如果客户端连接服务器时发生错误
            fmt.Println("客户端连接服务器的错误信息:", err)	//打印错误信息
            continue  // 继续等待下一个客户端连接
        }

		//成功处理
		//如果客户端连接成功，打印"服务端已成功连接客户端"
        fmt.Println("服务端已成功连接客户端")  

        //启动一个新的 goroutine 来并发处理当前连接的客户端请求
		//handleClient(conn) 函数将在一个单独的 goroutine 中执行，允许服务器同时处理多个客户端请求
        go handleClient(conn)
    }
}

//handleClient 是一个客户端请求的函数
//conn是与客户端的连接对象，代表客户端和服务器之间的通信通道
func handleClient(conn net.Conn) {
	
	//defer会延迟执行conn.Close()，直到main函数返回时才执行，确保资源在函数结束时被正确关闭
    defer conn.Close()  // 确保在函数结束时关闭连接
	
	//bufio.NewReader(conn) 创建一个新的bufio.Reader，用于从客户端连接中读取数据
	//reader存储从客户端连接中读取的数据
    reader := bufio.NewReader(conn)
    
	//启动一个无限循环，不断接收并处理客户端发送的消息
	for {

		//ReadString('\n') 会从客户端连接中读取数据，直到遇到换行符（即客户端发送的消息结束）为止
		// 返回的 message 是读取的消息内容，err 是可能发生的错误
        message, err := reader.ReadString('\n')  // 读取客户端发送的消息
        
		//错误处理
		if err != nil {	//如果从客户端连接中读取数据时发生错误
            fmt.Println("从客户端连接中读取数据的错误信息:", err) //打印错误信息
            return	//退出程序
        }

		//打印客户端发送的消息
        fmt.Printf("客户端发送的信息: %s", message)  

        //服务端生成响应并返回给客户端
        response := "服务端已成功接受到客户端发送的消息"
        
		//conn.Write将服务端的响应消息转换为字节数组并发送给客户端
		// "\n" 是确保消息后面有换行符，使得消息完整且能够在客户端正确解析
		conn.Write([]byte(response + "\n"))  
    }
}