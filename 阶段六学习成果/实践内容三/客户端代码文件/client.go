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