# 使用Wireshark分析网络包——周全

## 1.Wireshark观察TCP的报文结构

![image-20250208202101257](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250208202101257.png)

* **源端口号**: `55493`

* **目的端口号**: `443`
* **序列号**: `3934`
* **确认号**: `1075`
* **数据偏移**: `5 (20 bytes)`
* **保留位**: `0`
* **标志位**: `0x010 (ACK)`
* **窗口大小**: `251`
* **校验和**: `0x74c4`
* **紧急指针**: `0`
* **选项**: `无`
* **填充**: `无`
* **数据**: `无` 

## 2.Wireshark观察HTTP的报文结构

#### （1）HTTP请求报文结构

![image-20250216214433413](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250216214433413.png)

- **方法**: `POST`
- **HTTP版本**: `HTTP/1.1`
- **请求头**:
  - `Connection`: `Keep-Alive`
  - `Content-Type`: `application/json`
  - `dw-protocol`: `1.0`
  - `Content-Length`: `1083`
  - `Host`: `shuc-pc-hunt.ksord.com`
- **请求体**: `File Data: 1083 bytes`

#### （2）HTTP响应报文结构

![image-20250216214500104](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250216214500104.png)

- **HTTP版本**: `HTTP/1.1`
- **状态码**：`200`
- **状态描述**：`ok`
- **响应头**:
  - **Date**: `Sat, 08 Feb 2025 12:26:16 GMT`
  - **Content-Type**: `application/json;charset=utf-8`
  - **Content-Length**: `168`
  - **Connection**: `keep-alive`
  - **x-host**: `mics-qing-5d6bd5c98c-t8xrw`
  - **x-ver**: `qing|07-05 22:12:58|42ba5cc`
  - **x-envoy-upstream-service-time**: `20`
  - **X-KLB**: `2`
  - **x-request-id**: `2330344919164fe1a65b8f71dccaeb3`
  - **Expires**: `Sun, 09 Feb 2025 12:26:16 GMT`
  - **Cache-Control**: `max-age=86400`
  - **Nginx-Cache**: `HIT`
- **响应体**:`File Data: 168 bytes`

## 3.Wireshark观察TCP的三次握手

![image-20250208213400454](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250208213400454.png)

* **第一次握手**: `[SYN]`:客户端向服务器发送一个SYN报文，表示请求建立连接。

* **第二次握手**: `[SYN, ACK]`:服务器收到SYN报文后，回复一个SYN, ACK报文，表示同意建立连接，并发送自己的 SYN 请求。
* **第三次握手**: `[ACK]`:客户端收到SYN, ACK报文后，发送一个ACK报文，表示连接建立完成。

## 4.Wireshark观察TCP的四次挥手

![image-20250208214258858](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250208214258858.png)

* **第一次挥手**: `[FIN, ACK]`：客户端向服务器发送一个FIN包，表示请求断开连接。

* **第二次挥手**: `[ACK]`：服务器收到FIN包后，回复一个ACK包，表示同意断开连接。

* **第三次挥手**: `[FIN, ACK]`：服务器向客户端发送一个FIN包，表示服务器也准备断开连接。

* **第四次挥手**: `[ACK]`：客户端收到FIN包后，回复一个ACK包，表示连接断开完成。