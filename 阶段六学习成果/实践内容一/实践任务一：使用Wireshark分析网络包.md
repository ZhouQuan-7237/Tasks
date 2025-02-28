# 使用Wireshark分析网络包——周全

## 1.Wireshark观察TCP的报文结构

#### 步骤说明：

##### step1. 启动Wireshark，选择接口（Wlan）

双击网络连接接口开始捕获流量

##### step2. 获取目标IP地址（以百度为例）

终端输入以下指令：

* IPV4:

```bash
ping —4 www.baidu.com //获取百度的ipv4地址
```

![image-20250226183751781](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250226183751781.png)

* IPV6:

```bash
ping -6 www.baidu.com //获取百度的ipv6地址
```

![image-20250226073534849](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250226073534849.png)

##### step3. 设置Wireshark过滤器

* ##### IPV4:

```
ip.addr == 153.3.238.28 and tcp
```

表示只显示TCP协议且源主机IP或者目的主机IP为`153.3.238.28`的数据包

* IPV6:

```
ipv6.addr == 2408:873d:22:18cb:0:ff:b037:e6d8 and tcp
```

表示只显示TCP协议且源主机IP或者目的主机IP为`2408:873d:22:18cb:0:ff:b037:e6d8`的数据包。

##### step4. 触发TCP流量

访问目标网站（如浏览器输入 `www.baidu.com`）

或使用工具发送TCP请求：

```bash
curl http://www.baidu.com    # 发送HTTP请求（可能触发TCP连接）
```

##### step5. 查看TCP报文结构

![image-20250226184124218](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250226184124218.png)

![image-20250226184254174](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250226184254174.png)

* **源端口号**: `52426`

* **目的端口号**: `443`
* **序列号**: `1361`
* **确认号**: `1`
* **数据偏移**: `5 (20 bytes)`
* **保留位**: `0`
* **标志位**: `0x018 (PSH,ACK)`
* **窗口大小**: `255`
* **校验和**: `0x39c5`
* **紧急指针**: `0`
* **选项**: `无`
* **填充**: `无`
* **数据**: `395 bytes` 

## 2.Wireshark观察HTTP的报文结构

#### 步骤说明：

##### step1. 启动Wireshark，选择接口（Wlan）

双击网络连接接口开始捕获流量

##### step2. 设置Wireshark过滤器

Wireshark显示过滤器输入以下指令：

```bash
http.host == "baidu.com"    # 精确过滤目标域名
```

##### step3. 触发http流量

访问目标网站（如浏览器输入 `http://baidu.com`）

或使用工具发送TCP请求：

```bash
curl http://www.baidu.com/get    # 发送HTTP请求（可能触发http连接）
```

##### step5. 查看HTTP报文结构



以下图片为Wireshark显示过滤器输入`http`指令得到的结果：

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

#### 步骤说明：

##### step1. 启动Wireshark，选择接口（Wlan）

双击网络连接接口开始捕获流量

##### step2. 设置Wireshark过滤器

Wireshark显示过滤器输入以下指令：

```
(tcp.flags.syn == 1 && tcp.flags.ack == 0) || (tcp.flags.syn == 1 && tcp.flags.ack == 1) || (tcp.flags.ack == 1 && tcp.flags.syn == 0 && tcp.flags.fin == 0)
```

![image-20250208213400454](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250208213400454.png)

* **第一次握手**: `[SYN]`:客户端向服务器发送一个SYN报文，表示请求建立连接。

* **第二次握手**: `[SYN, ACK]`:服务器收到SYN报文后，回复一个SYN, ACK报文，表示同意建立连接，并发送自己的 SYN 请求。
* **第三次握手**: `[ACK]`:客户端收到SYN, ACK报文后，发送一个ACK报文，表示连接建立完成。

## 4.Wireshark观察TCP的四次挥手

#### 步骤说明：

##### step1. 启动Wireshark，选择接口（Wlan）

双击网络连接接口开始捕获流量

##### step2. 设置Wireshark过滤器

Wireshark显示过滤器输入以下指令：

```
(tcp.flags.fin == 1 && tcp.flags.ack == 0) || (tcp.flags.fin == 1 && tcp.flags.ack == 1) || (tcp.flags.ack == 1 && tcp.flags.fin == 0)
```

![image-20250208214258858](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250208214258858.png)

* **第一次挥手**: `[FIN, ACK]`：客户端向服务器发送一个FIN包，表示请求断开连接。

* **第二次挥手**: `[ACK]`：服务器收到FIN包后，回复一个ACK包，表示同意断开连接。

* **第三次挥手**: `[FIN, ACK]`：服务器向客户端发送一个FIN包，表示服务器也准备断开连接。

* **第四次挥手**: `[ACK]`：客户端收到FIN包后，回复一个ACK包，表示连接断开完成。