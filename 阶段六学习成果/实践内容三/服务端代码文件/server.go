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