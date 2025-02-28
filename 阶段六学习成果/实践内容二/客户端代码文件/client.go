package main //package定义Go程序的包,main是Go语言中程序的入口包。每个Go程序必须有一个main包，且程序执行时会从main函数开始执行

import (	 //import导入其他包,在当前文件中使用
    "bufio"  //提供缓冲读写功能，用于高效的读取输入和输出
    "fmt"    //用于格式化输入输出，打印信息到控制台
    "net"    //用于实现网络通信功能，建立连接和传输数据
    "os"     //用于与操作系统交互，如读取控制台输入
)

func main() {   // 定义程序的主函数,程序从此处开始执行

    //通过net.Dial方法发起一个TCP连接，尝试连接到本地计算机（localhost）的 8080 端口
    //net.Dial函数返回两个值：连接对象存储在conn中，同时将可能的错误存储在err变量中
	conn, err := net.Dial("tcp", "localhost:8080")
    
	//错误处理
	if err != nil {  //如果net.Dial返回的err不为 nil，表示发起连接时发生了错误
        fmt.Println("客户端的错误信息:", err)  //打印发起连接时的错误信息
        return  //终止执行main函数
    }

    //defer会延迟执行conn.Close()，直到main函数返回时才执行，确保资源在函数结束时被正确关闭
    defer conn.Close()  // 使用defer确保程序结束时关闭连接，无论程序是否正常退出

	//如果连接成功，打印"客户端已成功连接服务端"
    fmt.Println("客户端已成功连接服务端") 

    //创建一个读取器bufio.NewReader，它用来从标准输入（即用户在命令行中的输入）读取数据
    //bufio.NewReader创建一个读取器, os.Stdin表示标准输入（Standard Input）,通常是指通过键盘输入的内容 
    //bufio.NewReader(os.Stdin)用标准输入作为读取器的输入源
    //reader := bufio.NewReader(os.Stdin)创建一个读取标准输入的读取器,可以读取用户通过键盘输入的文本
    reader := bufio.NewReader(os.Stdin)

	//启动一个无限循环
	//客户端将在这个循环中不断接收用户输入并与服务器进行交互，直到用户输入exit为止
    for {  
        
		//提示用户在客户端输入消息
		fmt.Print("请输入消息: ")  
        
		// 使用ReadString方法从标准输入读取字符串，直到遇到换行符（'\n'）为止
		//读取到的字符串存储在message中，可能发生的错误存储在err中
		message, err := reader.ReadString('\n')  
        //错误处理
		if err != nil {  //如果reader.ReadString返回的err不为 nil，表示读取字符串时发生了错误
            fmt.Println("读取字符串的错误信息:", err)  //打印读取字符串时的错误信息
            return  //终止执行main函数
        }

        // 将用户输入的消息发送到服务器，通过TCP连接
		//Write 方法将消息转换为字节流并通过连接发送
		// "\n" 用来确保消息在服务器端可以正确分隔
		//可能的错误存储在err变量中
        _, err = conn.Write([]byte(message + "\n")) 
        //错误处理
		if err != nil {  //如果Write方法返回的err不为 nil，表示发送消息时发生了错误
            fmt.Println("发送消息的错误信息:", err)  // 打印发送消息时错误信息
            return  //终止执行main函数
        }

        // 读取服务器的响应，直到遇到换行符
		//创建一个新的bufio.Reader，从连接中读取服务器返回的响应，直到遇到换行符为止
		//读取的服务器返回的响应返回给response，可能的错误返回给err
        response, err := bufio.NewReader(conn).ReadString('\n')  
        //错误处理
		if err != nil {  ///如果bufio.NewReader返回的err不为 nil，表示读取服务器响应时发生了错误
            fmt.Println("读取服务器响应的错误信息:", err)  // 打印读取服务器响应时的错误信息
            return  //终止执行main函数
        }

        // 打印服务器返回的响应内容
        fmt.Printf("%s\n", response)  
    }
}