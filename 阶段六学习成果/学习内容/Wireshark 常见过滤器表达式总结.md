### **Wireshark 常见过滤器表达式总结**

------

#### **1. 捕获过滤器（Capture Filters）**

捕获过滤器在数据包捕获阶段进行筛选，使用 Berkeley Packet Filter（BPF）语法。![image-20250318075026269](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250318075026269.png)

- **协议过滤**：
  - 捕获所有 TCP 包：`tcp`
  - 捕获所有 UDP 包：`udp`
  - 捕获所有 ICMP 包：`icmp`
- **IP 地址过滤**：
  - 捕获特定 IP 地址：`host 192.168.1.1`
  - 捕获特定源 IP 地址：`src host 192.168.1.1`
  - 捕获特定目标 IP 地址：`dst host 192.168.1.1`
- **端口过滤**：
  - 捕获特定端口的数据包：`port 80`
  - 捕获指定源端口的数据包：`src port 80`
  - 捕获指定目标端口的数据包：`dst port 80`
- **源或目标地址组合**：
  - 捕获特定 IP 对的数据包：`host 192.168.1.1 and host 192.168.1.2`
  - 排除特定 IP 地址的数据包：`not net 172.16.0.52`
- **组合过滤**：
  - 捕获特定源地址和目标端口的数据包：`src host 192.168.1.1 and port 80`
  - 捕获特定协议的数据包：`tcp or udp`

------

#### **2. 显示过滤器（Display Filters）**

显示过滤器用于捕获后对数据包进行分析和过滤，语法更为灵活。

- **协议过滤**：
  - 过滤所有 TCP 包：`tcp`
  - 过滤 HTTP 请求：`http.request`
  - 过滤 HTTP 响应：`http.response`
  - 过滤 ICMP 包：`icmp`
- **IP 地址过滤**：
  - 过滤特定源 IP 地址：`ip.src == 192.168.1.1`
  - 过滤特定目标 IP 地址：`ip.dst == 192.168.1.1`
  - 过滤特定源/目标 IP 地址 ：`ip.addr == 192.168.1.1`
- **端口过滤**：
  - 过滤特定 TCP 端口：`tcp.port == 80`
  - 过滤特定 UDP 端口：`udp.port == 53`
  - 过滤特定源端口：`tcp.srcport == 8080`
- **复杂逻辑过滤**：
  - 过滤 TCP 三次握手的数据包：`tcp.flags.syn == 1 && tcp.flags.ack == 0 || tcp.flags.syn == 1 && tcp.flags.ack == 1 || tcp.flags.ack == 1 && tcp.flags.syn == 0 && tcp.flags.fin == 0`
  - 过滤 TCP 四次挥手的数据包：`tcp.flags.fin == 1 || (tcp.flags.ack == 1 && tcp.flags.fin == 1)`

---

#### 3. 逻辑操作符：

*  **AND**：`tcp && ip.src == 192.168.1.1`

- **OR**：`http || dns`

- **NOT**：`!arp`

------

#### **4. TCP 三次握手和四次挥手的过滤表达式**

##### **TCP 三次握手过滤**

三次握手用于建立 TCP 连接，涉及 SYN 和 ACK 标志。常用的过滤表达式如下：

- **第一次握手**（SYN 包）：`tcp.flags.syn == 1 && tcp.flags.ack == 0`
- **第二次握手**（SYN+ACK 包）：`tcp.flags.syn == 1 && tcp.flags.ack == 1`
- **第三次握手**（ACK 包）：`tcp.flags.ack == 1 && tcp.flags.syn == 0 && tcp.flags.fin == 0`

**综合表达式**：

```plaintext
(tcp.flags.syn == 1 && tcp.flags.ack == 0) || (tcp.flags.syn == 1 && tcp.flags.ack == 1) || (tcp.flags.ack == 1 && tcp.flags.syn == 0 && tcp.flags.fin == 0)
```

##### **TCP 四次挥手过滤**

四次挥手用于断开 TCP 连接，涉及 FIN 和 ACK 标志。常用的过滤表达式如下：

- **第一次挥手**（FIN 包）：`tcp.flags.fin == 1 && tcp.flags.ack == 0`
- **第二次挥手**（ACK 包）：`tcp.flags.fin == 1 && tcp.flags.ack == 1`
- **第三次挥手**（FIN 包）：`tcp.flags.ack == 1 && tcp.flags.fin == 0`
- **第四次挥手**（ACK 包）：`tcp.flags.ack == 1 && tcp.flags.fin == 1`

**综合表达式**：

```plaintext
(tcp.flags.fin == 1 && tcp.flags.ack == 0) || (tcp.flags.fin == 1 && tcp.flags.ack == 1) || (tcp.flags.ack == 1 && tcp.flags.fin == 0)
```

##### **同时捕获三次握手和四次挥手**

```plaintext
((tcp.flags.syn == 1 && tcp.flags.ack == 0) || (tcp.flags.syn == 1 && tcp.flags.ack == 1) || (tcp.flags.ack == 1 && tcp.flags.syn == 0 && tcp.flags.fin == 0)) || ((tcp.flags.fin == 1 && tcp.flags.ack == 0) || (tcp.flags.fin == 1 && tcp.flags.ack == 1) || (tcp.flags.ack == 1 && tcp.flags.fin == 0))
```