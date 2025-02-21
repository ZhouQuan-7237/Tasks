# Gin入门学习笔记——周全

## 学习目录

一.MVC 架构分层

1.MVC模式

2.工作流程

3.三层架构

4.二者区别

二.Go 项目结构

1.目录结构

2.结构说明

三.Gin 基础学习

1.基础介绍

2.基础语法

3.基础操作

四.RESTful API

1.API介绍

2.RESTful API

五.验证、权限与加密

1.Cookie

2.Session

3.Token

4.JWT

## 学习内容

### 一.MVC 架构分层

#### 1.MVC模式

![739ba8787c0eb73782091de919c6141d](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/739ba8787c0eb73782091de919c6141d.png)

MVC（Model-View-Controller）架构是一种将应用程序分层的设计模式，它把软件系统分为模型、视图和控制器三个基本部分。

* **Model（模型）**：

  - 负责数据处理和业务逻辑，直接与数据库交互。

  - 包括对数据的增、删、改、查等操作。

* **View（视图）**：

  - 负责呈现数据给用户，即用户界面部分。

  - 视图接收用户的输入并将其传递给控制器，不直接处理数据。

* **Controller（控制器）**：

  * 充当模型和视图之间的协调者，负责接收来自视图的输入，并决定使用哪种模型来处理数据。

  - 控制器通过调用模型的业务逻辑方法，将处理后的结果传回给视图进行展示。

#### 2.工作流程

1. 用户通过视图层发起请求（如点击按钮或提交表单）。
2. 控制器接收请求，解析并调用模型层处理业务逻辑。
3. 模型层完成数据处理后，将结果返回给控制器。
4. 控制器根据结果选择合适的视图层进行数据展示。

#### 3.三层架构

![image-20250222005842643](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250222005842643.png)

三层架构是一种经典的分层设计方法，它将应用程序分为三层：

* **表现层（UI层，View）**：

  - 直接与用户交互，接收用户请求并返回展示数据。

  - 比如，前端的网页界面，通过 Ajax 请求与后端交互。

* **业务逻辑层（Service层，BLL）**：

  - 负责业务逻辑的处理。

  - 接收来自表现层的请求，处理后调用数据访问层获取数据并返回给表现层。

* **数据访问层（DAO层，DAL）**：

  - 负责直接操作数据库，进行 CRUD（增、删、改、查）操作。

  - 通过对数据库的抽象封装，提供统一的数据访问接口。

#### 3.两者关系

![2058524bf698f0eddff2ca7b4040b568](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/2058524bf698f0eddff2ca7b4040b568.png)

- MVC是三层架构的具体实现：

  - 视图层和控制层对应三层架构的表现层。
  - 模型层对应三层架构的业务逻辑层和数据访问层。

- 二者都强调高内聚、低耦合，

  但MVC更侧重于前端的视图和业务逻辑的分离，

  而三层架构更关注后端的层次划分。

### 二.Go项目结构

* #### **目录结构**

```
/projectname
├── cmd/                # 存放各个应用的入口
│   └── myapp/          # 每个应用都可以在这个文件夹下有一个独立的目录
│       └── main.go     # 应用入口文件
├── internal/           # 存放不对外部公开的代码
│   └── mypackage/      # 只在本项目内部使用的包
│       └── logic.go
├── pkg/                # 公共库和可复用模块
│   └── mypackage/      # 可被其他项目引用的公共包
│       └── utils.go
├── api/                # 存放 API 定义，通常是 Protobuf 文件
│   └── myservice.proto # 定义服务接口
├── scripts/            # 存放自动化脚本，如部署、数据库迁移等
│   └── deploy.sh
├── web/                # 存放 web 相关的前端、模板文件等
│   └── static/
│   └── templates/
├── test/               # 单元测试、集成测试等
│   └── mypackage_test.go
├── go.mod              # Go 模块文件
├── go.sum              # Go 校验文件
└── README.md           # 项目说明文件

```

* #### **结构说明**

  * **cmd**

    - 存放项目中各个应用程序的入口文件（如 `main.go`）。

    - 每个应用程序有一个独立的子目录。比如：`cmd/myapp/main.go`。

  * **internal**

    - 存放项目内部使用的代码。

    - `internal` 文件夹的内容仅能在同一项目内使用，不能被外部项目引用，从确保项目的封装性和安全性。

  * **pkg**

    - 存放可以被其他项目引用的公共库代码。

    - 比如：通用工具函数、服务组件、数据库操作等。

  * **api**

    - 存放 API 的定义文件。

    - 方便统一管理和版本控制。

  * **scripts**

    * 存放自动化脚本。

    -  比如：`deploy.sh`、`migrate.sh`、`build.sh` 等。

  * **web**

    - 如果项目涉及 Web 服务，该目录将包含静态文件、模板文件以及前端资源等。

    - 比如：HTML文件、CSS文件、JavaScript文件等。

  * **test**

    - 存放项目的测试代码。

    - 通常与生产代码并列，文件名以 `_test.go` 为后缀。

  * **go.mod 和 go.sum**

    - 用于管理 Go 项目的依赖关系。

    - `go.mod` 文件记录项目依赖的模块及版本信息， `go.sum` 文件确保依赖模块的版本一致性和安全性。

  * **README.md**

    - 项目的文档说明文件。

    - 包含项目概述、安装、运行说明及贡献指南等。

### 三.Gin基础学习

#### 1.基础介绍

Gin是一个用Go语言编写的高性能Web框架，设计简洁，专注于快速构建 Web 应用程序和 API 服务，适合高并发场景。

#### 2.基础语法

##### 2.1Web服务

```go
r := gin.Default()  // 创建Gin路由器,初始化Web服务
r.Run(":8080")       // 启动Web服务，监听8080端口
```

##### 2.2路由

* **路由定义**：路由是根据请求的HTTP方法和URL路径，将客户端请求分发到对应处理函数的机制。

* **路由结构**

  ```go
  r.Method("PATH", handler {
  	//处理函数执行逻辑
  })
  ```

  访问`URL路径`会触发`Method`请求，并执行`handle`函数。

  * **HTTP 方法**

    * **GET **: 获取资源或数据。

    * **POST **: 提交数据或创建资源。

    * **PUT **: 更新资源。

    * **DELETE **: 删除资源。

    * **PATCH **: 部分更新资源。

      ```go
      // GET 请求
      r.GET("/path", func(c *gin.Context) {
      	// 处理 GET 请求逻辑
      	c.JSON(200, gin.H{"message": "This is a GET request"})
      })
      
      // POST 请求
      r.POST("/path", func(c *gin.Context) {
      	// 处理 POST 请求逻辑
      	c.JSON(200, gin.H{"message": "This is a POST request"})
      })
      
      // PUT 请求
      r.PUT("/path", func(c *gin.Context) {
      	// 处理 PUT 请求逻辑
      	c.JSON(200, gin.H{"message": "This is a PUT request"})
      })
      
      // DELETE 请求
      r.DELETE("/path", func(c *gin.Context) {
      	// 处理 DELETE 请求逻辑
          c.JSON(200, gin.H{"message": "This is a DELETE request"})
      })
      
      // PATCH 请求
      r.PATCH("/path", func(c *gin.Context) {
      	// 处理 PATCH 请求逻辑
      	c.JSON(200, gin.H{"message": "This is a PATCH request"})
      })
      
      //'/path'为请求路径，func(c *gin.Context)是路由处理函数，c是上下文对象。
      ```

  * **URL 路径**

    比如：

    - `/users`：表示用户资源。
    - `/user/:id`：表示动态路径，`:id` 是路径参数。

  * **处理函数**

    处理函数是在请求匹配路由时执行的逻辑，通常包括：

    - 获取请求参数。
    - 执行业务逻辑。
    - 返回响应。

##### 2.3路由分组

路由分组是将相关的路由组织在一起，便于统一管理。

```go
// 创建一个以 "/user" 为路由前缀路径的路由分组
userGroup := r.Group("/user")
{
	//在分组中定义路由
	
    // 定义 GET 请求，路径为 "/user/:id"，用于根据 ID 获取用户信息
    userGroup.GET("/:id", getUsers)  // 示例路径：/user/123
    
    // 定义 POST 请求，路径为 "/user/"，用于创建新用户
    userGroup.POST("/", createUsers)    // 示例路径：/user/
}
```

**2.4请求上下文**

请求上下文（`gin.Context`）是用于封装请求和响应相关信息的结构体，简化了请求参数获取、响应返回以及中间件使用等操作。

* **请求参数获取**：请求参数可以通过 `gin.Context` 提供的多种方法进行获取：

  * **路径参数**：`c.Param("key")`,获取动态路由中的参数。

    ```go
    id := c.Param("id")  // 获取路径参数
    ```

  - **查询参数**：获取 URL 查询字符串中的参数。`c.Query("key")`获取指定参数的值，`c.DefaultQuery("key", "default")`，如果不存在该参数，则返回默认值。

    ```go
    value := c.Query("key")              // 获取查询参数
    value := c.DefaultQuery("key", "default")  // 获取查询参数，默认值为 "default"
    ```

  - **表单参数**：获取 POST 请求中的表单数据,`c.PostForm("key")`获取指定参数的值，`c.DefaultPostForm("key", "default")`，如果不存在该参数，则返回默认值。

    ```go
    value := c.PostForm("key")              // 获取表单参数
    value := c.DefaultPostForm("key", "default")  // 获取表单参数，默认值为 "default"
    ```

  - **JSON请求体**：用于从请求体中解析 JSON 数据，`c.ShouldBindJSON()` 将请求体绑定到结构体。

    ```go
    var data MyStruct
    if err := c.ShouldBindJSON(&data); err != nil {
        // 处理错误
    }
    ```

* **中间件使用**:中间件是位于请求处理流程中的函数，用于在请求到达路由处理函数之前或响应返回客户端之前执行共享逻辑，如身份验证、日志记录、错误处理等。

  * **注册中间件**：通过`r.Use()`注册中间件。

    ```go
    r.Use(middleware)
    ```

  * **自定义中间件**：可以创建自定义中间件函数，如日志记录、身份验证等。

    ```go
    // 自定义中间件示例：日志记录
    r.Use(func(c *gin.Context) {
        log.Println("Request started") // 请求开始时的日志
        c.Next()                       // 调用下一个处理函数
        log.Println("Request completed") // 请求完成时的日志
    })
    ```

* **响应返回**：

  - **JSON响应**：返回 JSON 格式的数据，常用于 API 开发。

    ```go
    c.JSON(200, gin.H{"key": "value"})
    ```

  - **字符串响应**：返回纯文本数据。

    ```go
    c.String(200, "Hello")
    ```

  - **HTML响应**：返回 HTML 页面。

    ```go
    c.HTML(200, "template.html", data)
    ```

  - **文件响应**：返回文件。

    ```go
    c.File("path/to/file")
    ```

##### **2.5参数绑定**

参数绑定是将请求中的数据（如路径、查询、表单或JSON等）自动解析并映射到Go语言结构体或变量中。在Gin中，使用 `ShouldBind()`系列方法将请求数据绑定到结构体。

* 绑定方法

  * **`ShouldBindJSON()`**：绑定 JSON 请求体到结构体。

    ```go
    var user struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }
    
    if err := c.ShouldBindJSON(&user); err == nil {
        c.JSON(200, gin.H{"message": "User created", "user": user})
    } else {
        c.JSON(400, gin.H{"error": err.Error()})
    }
    ```

  * **`ShouldBindUri()`**：绑定 URI 路径参数到结构体。

    ```go
    var user struct {
        ID string `uri:"id"`
    }
    
    if err := c.ShouldBindUri(&user); err == nil {
        c.JSON(200, gin.H{"message": "User retrieved", "id": user.ID})
    } else {
        c.JSON(400, gin.H{"error": err.Error()})
    }
    ```

  * **`ShouldBindQuery()`**：绑定查询参数到结构体。

    ```go
    var user struct {
        Name string `query:"name"`
        Age  int    `query:"age"`
    }
    
    if err := c.ShouldBindQuery(&user); err == nil {
        c.JSON(200, gin.H{"message": "User retrieved", "user": user})
    } else {
        c.JSON(400, gin.H{"error": err.Error()})
    }
    ```

    

  * **`ShouldBindForm()`**：绑定表单数据到结构体。

    ```go
    var user struct {
        Name string `form:"name"`
        Age  int    `form:"age"`
    }
    
    if err := c.ShouldBindForm(&user); err == nil {
        c.JSON(200, gin.H{"message": "User created", "user": user})
    } else {
        c.JSON(400, gin.H{"error": err.Error()})
    }
    ```

##### **2.6错误处理**

* **返回错误信息**：可以通过`c.JSON()`或`c.String()`返回错误信息

```go
c.JSON(400, gin.H{"error": "Bad request"})
```

* **自定义错误处理** :可以根据需求自定义错误响应，控制返回的状态码、错误信息以及格式。

```
c.String(500, "Internal Server Error")
```

* **全局错误处理** :可以使用`gin.DefaultErrorWriter`来处理全局错误。

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// 全局错误处理中间件
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				c.JSON(500, gin.H{"error": "Internal server error"})
			}
		}()
		c.Next()
	})

	r.GET("/panic", func(c *gin.Context) {
		// 模拟 panic
		panic("Simulated panic")
	})

	r.Run(":8080")
}
```

##### **2.7静态文件服务**

静态文件服务是指通过Gin的`Static()`方法将本地文件夹中的文件暴露为Web资源，

供客户端通过指定路径直接访问。

```go
r.Static("/assets", "./static")
//把本地目录 ./static 下的文件映射到路由 /assets，使得用户可以通过访问 http://localhost/assets/filename 来获取静态文件
```

#### 3.基础操作

##### 3.1**环境搭建**

* **（1）下载安装**

  ```go
  go get -u github.com/gin-gonic/gin
  ```

  若无法下载，执行以下指令

  ```go
  go env -w GO111MODULE=on
  go env -w GOPROXY=https://goproxy.cn,direct
  ```

* **（2）项目导入**

  ```go
  import "github.com/gin-gonic/gin"
  ```

* **（3）简单示例**

  ```go
  package main
  
  import (
  	"github.com/gin-gonic/gin"
  )
  
  func main() {
      // 创建 Gin 引擎
  	r := gin.Default()
      //定义路由
  	r.GET("/", func(c *gin.Context) {
  		c.JSON(200, gin.H{
  			"message": "Hello,Gin!",
  		})
  	})
      //启动服务器，监听在 8080 端口
      r.Run(":8080") 
  }
  ```

##### 3.2处理请求

```go
package main

// 导入 Gin 框架
import (
	"github.com/gin-gonic/gin"
)

func main() {
    //创建Gin路由器,初始化Web服务
	r := gin.Default()

	// 定义一个 GET 路由。
	// 当用户访问根路径 `/` 时，Gin会执行下面的处理函数
	r.GET("/", func(c *gin.Context) {
		// c.String() 返回一个字符串响应。
		// 第一个参数200是 HTTP 状态码，表示请求成功
		// 第二个参数 "Hello, World!" 是我们返回的HTTP内容
		c.String(200, "Hello, World!")
	})

	// 启动Web服务，监听8080端口。
	// 当浏览器访问 http://localhost:8080 时，会触发相应的路由和处理函数
	// 注意检查8080端口是否被占用：ss -tuln | grep 8080 ,这条指令会列出所有监听的端口并查找 8080 端口的相关信息
	r.Run(":8080") 
}
```

![image-20250212181551494](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250212181551494.png)

### 四.RESTful API

#### 1.API介绍

API（应用程序编程接口）是不同软件系统之间互相通信和交互的接口，允许一个程序访问另一个程序的功能或数据。

##### 1.1API工作原理

**(1) 请求**：客户端向服务器发出请求，带上需要的参数和数据。

**(2) 处理**：服务器接收到请求后，根据请求内容进行处理，可能是查数据库、执行操作等。

**(3) 响应**：服务器处理完后，把结果返回给客户端，通常是数据和一个状态，比如“成功”或“失败”。

##### 1.2API分类

**(1) RESTful API**：基于HTTP协议，使用GET、POST、PUT、DELETE等方法来操作资源，适用于Web和移动应用，强调无状态和统一接口。

**(2) RPC API**：通过远程调用服务器端的函数或方法，通常用于内部服务之间的通信。

**(3) GraphQL API**：一种查询语言，客户端可以精确选择需要的数据，解决了传统RESTful API的“数据过多或过少”问题。

**(4) SOAP API**：基于XML协议，主要用于企业级应用之间的通信，强调安全性和可靠性。

#### 2.RESTful API

RESTful API 是基于HTTP 协议和 REST 架构风格的 API 设计方式，使用 URL 来操作资源，强调无状态性和统一接口。

##### 2.1设计原则

**(1) 统一接口**：使用标准的HTTP方法（GET、POST、PUT、DELETE等）来操作资源。

**(2) 无状态**：每个请求都包含完整的信息，相互独立，服务器不存储客户端状态。

**(3) 可缓存**：通过HTTP缓存机制提高性能。

**(4) 分层系统**：支持多层架构，便于扩展和维护。

##### 2.2URL设计规范

**(1) 资源导向**：URI 应该使用名词来表示资源，避免使用动词（如 `/users` 而非 `/getUsers`）。

**(2) 复数形式**：使用复数形式表示资源集合（如 `/users` 表示多个用户）。

**(3) 版本号**：可以在 URL 中添加版本号（如 `/v1/users`）以便后期版本升级时不破坏兼容性。

**(4) 过滤参数**：支持分页、排序、筛选等操作，通过查询参数（如 `?limit=10&offset=20`）来实现。

##### 2.3资源标识

**资源标识**是通过**URI**（统一资源标识符）来唯一标识和定位RESTful API中的资源。每个资源都有一个清晰的URI，用于表示该资源的位置。URI设计通常采用名词表示资源，并遵循层次化结构，确保简洁、一致和易于理解。例如，`/users`表示所有用户资源，`/users/123`表示特定ID为123的用户。

##### **2.4媒体类型**

**（1）JSON** ：最常用的媒体类型，可读性较好，易于解析。

```go
{
  "id": 123,
  "name": "John Doe",
  "email": "john@example.com"
}
```

**(2)XML** ：另一种常用的媒体类型，尤其在企业级应用中广泛使用。

```go
<user>
  <id>123</id>
  <name>John Doe</name>
  <email>john@example.com</email>
</user>
```

### 五.验证、权限与加密

**1.Cookie**：Cookies 是由服务器生成并存储在客户端浏览器中的小数据块，常用于会话跟踪和身份验证。浏览器首次访问网站时，服务器通过 HTTP 响应将 Cookies 发送给客户端并存储。此后，浏览器每次向同一网站发起请求时，都会自动将 Cookies 发送回服务器，帮助服务器识别用户身份。

* **工作流程**：

  ![619c97ef1b452faf27d23a81b23c1bf4](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/619c97ef1b452faf27d23a81b23c1bf4.png)

  **（1）客户端发起请求**：当用户首次访问网站时，浏览器会发起一个 HTTP 请求，向服务器请求所需的资源。

  **（2）服务器响应并设置 Cookies**：服务器接收到请求后，会通过 HTTP 响应头的 `Set-Cookie` 字段将 Cookie 发送到客户端浏览器。此 Cookie 通常包含用户的会话标识符或身份验证信息。（从打开浏览器访问一个网站到关闭浏览器结束此次访问称为一个会话）

  **（3）客户端存储 Cookies**：浏览器接收到服务器发送的 Cookie 后，会根据 Cookie 的设置（如有效期、路径、域名等）将其存储在本地，以便后续使用。

  **（4）客户端携带 Cookies 再次发送请求**：在用户后续访问相同网站时，浏览器会自动将存储的 Cookies 添加到 HTTP 请求头中的 `Cookie` 字段，并随请求一起发送给服务器。

  **（5）服务器通过 Cookie 识别用户**：服务器接收到请求后，解析请求头中的 Cookie 信息，进而识别用户身份、维持会话状态或为用户提供个性化服务。

  **（6）服务器更新或删除 Cookies**：如果需要，服务器可以在随后的响应中通过 `Set-Cookie` 字段更新现有的 Cookie（如修改值或设置新的有效期），或删除特定的 Cookie（通过设置过期时间为过去的时间）。

* **缺点**：

  - Cookies 会暴露在客户端，可能会被恶意用户窃取。

  - Cookies 有大小限制（通常 4KB）。

**2.Session**：Session 是一种存储在服务器端的会话管理机制，用于跟踪和管理用户的状态和会话信息。当用户访问网站时，服务器会为每个用户创建一个唯一的 Session ID，并将与该会话相关的数据（如登录状态、用户信息、购物车内容等）存储在服务器上。Session ID 通常通过 Cookie 或 URL 参数传递给客户端，客户端在后续请求中携带该 Session ID，服务器通过它查找并使用对应的会话数据。

* **工作流程**

  ![image-20250222005958077](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250222005958077.png)

  **（1）用户第一次请求服务器**： 当用户第一次访问网站时，浏览器会发起一个 HTTP 请求。服务器接收到该请求后，根据用户的相关信息（如登录状态、IP 地址等）创建一个新的 Session。

  **（2）服务器返回 SessionID**： 服务器生成一个唯一的 Session ID，并将其与会话数据（如用户信息、登录状态等）存储在服务器端。然后，服务器通过 HTTP 响应头中的 `Set-Cookie` 字段将生成的 Session ID 返回给客户端浏览器。

  **（3）浏览器存储 SessionID**： 客户端浏览器接收到服务器返回的 Session ID 后，将其存储在浏览器的 Cookie 中。Cookie 会记录这个 Session ID 所属的域名和相关的有效期等信息。

  **（4）后续请求携带 SessionID**： 当用户第二次访问服务器时，浏览器会自动将存储在 Cookie 中的 Session ID 作为请求的一部分，发送回服务器。此时，浏览器会将 Cookie 中的 SessionID 添加到 HTTP 请求头中的 `Cookie` 字段中。

  **（5）服务器根据 SessionID 验证用户身份**： 服务器接收到带有 Session ID 的请求后，会通过解析该 Session ID 来查找对应的会话数据。如果找不到相关的 Session 数据，说明用户未登录或者登录已过期；如果找到，服务器则验证该用户身份，并可以继续执行后续的操作。

* **优点**
  * **安全性较高**：用户数据存储在服务器端，客户端只能通过 Session ID 与服务器进行交互，减少了数据被篡改或盗用的风险。
  * **适用于大数据存储**：Session 可以存储比 Cookies 更大的数据量。
* **缺点**
  * **服务器负担**：Session 数据存储在服务器端，随着用户量增加，服务器需要存储更多的会话数据，这可能会导致服务器负担加重。
  * **会话丢失风险**：如果用户关闭浏览器或 Session 过期，服务器端的会话数据会丢失，用户需要重新登录或重新建立会话。
* **Session 与 Cookie 的区别**
  - **存储位置**：
    - **Session**：存储在服务器端。
    - **Cookie**：存储在客户端。
  - **安全性**：
    - **Session**：较高，数据不会暴露在客户端。
    - **Cookie**：较低，容易被窃取或篡改。
  - **生命周期**：
    - **Session**：通常与用户的会话相关，可以在用户关闭浏览器或会话超时后失效。
    - **Cookie**：可以设置过期时间，过期后自动删除。
  - **存储限制**：
    - **Session**：没有大小限制，可以存储更多的数据。
    - **Cookie**：通常限制在 4KB 左右。

**3.Token**：Token 是一种加密字符串，通常用于身份认证和授权。它由服务器在用户登录时生成，并返回给客户端。客户端将 Token 存储并在后续的请求中携带，以进行身份验证。与传统的会话管理不同，Token 不依赖服务器存储状态。每次请求时，客户端都会将 Token 作为认证凭证附加在请求头中，服务器通过验证 Token 来确认用户身份。这种无状态的身份验证方式非常适合用于 API 请求和跨域认证场景。

* **工作流程**

  ![image-20250222010045418](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250222010045418.png)

  **（1）用户登录**： 用户通过浏览器或移动应用向服务器发送包含用户名和密码的登录请求。

  **（2）服务器验证**： 服务器验证用户信息，验证成功后生成一个唯一的 Token，通常包含用户身份和权限信息，并通过加密和签名保证安全性和完整性。

  **（3）返回 Token**： 服务器将生成的 Token 返回给客户端。

  **（4）客户端存储 Token**： 客户端将 Token 存储在本地（如 `localStorage`、`sessionStorage` 或 `Cookie`）以便后续使用。

  **（5）后续请求携带 Token**： 在后续的 API 请求中，客户端将 Token 附加到请求头中的 `Authorization` 字段（格式为 `Authorization: Bearer <Token>`）。

  **（6）服务器验证 Token**： 服务器解析并验证 Token 的有效性。验证成功后，服务器根据 Token 中的信息（如用户身份和权限）判断是否允许访问资源；验证失败则拒绝请求并返回错误信息。

* **特点**

  * **无状态认证**：Token 不依赖于服务器的状态，它不需要服务器存储会话数据。每次请求都携带 Token，服务器仅通过 Token 的内容进行验证，无需存储用户会话。
  * **跨域支持**：由于 Token 存储在客户端，不受同源策略限制，适合跨域认证的场景。

* **应用场景**

  * **API 认证**：适用于无状态的身份验证，特别是在使用 RESTful API 时，Token 可以在不同的服务之间传递，验证用户身份。
  * **跨平台认证**：在不同的设备和平台（如移动端、Web端等）之间，Token 可以方便地在各个平台上进行身份认证。

**4.JWT**：JWT 是一种基于JSON格式的认证Token，包含头部、有效载荷和签名部分。最大优势在于自包含，即能够在 Token 中直接存储用户的身份信息，并通过签名确保数据的完整性和真实性。

* **JWT结构**

  * **头部** 

    通常包含两个字段

    - **`typ`**：指定 Token 类型，通常为 `JWT`。

    - **`alg`**：指定签名算法，常见的有 `HS256`（HMAC SHA-256）和 `RS256`（RSA SHA-256）。

  ```json
  {
    "alg": "HS256",
    "typ": "JWT"
  }
  ```

  * **有效载荷 **

    有效载荷包含实际要传递的声明（Claims）。声明通常是关于用户身份、权限等信息。有效载荷包括以下类型的声明：

    - 注册声明：JWT 规范定义的标准声明，例如：
      - `iat`（Issued At）：签发时间。
      - `exp`（Expiration Time）：过期时间。
      - `sub`（Subject）：主题（通常是用户 ID）。

    - **公共声明**：自定义声明，可以传递任何信息，需确保不会与其他系统冲突。

    - **私有声明**：双方约定使用的信息，通常用于不同系统间交换数据。

  ```json
  {
    "sub": "1234567890",
    "name": "John Doe",
    "iat": 1516239022,
    "exp": 1516239022 + 3600  // 1小时后过期
  }
  ```

  *  **签名** 

  签名用于验证 Token 的有效性，确保数据未被篡改并验证其来源。签名生成过程如下：

  1. 对头部和有效载荷进行 Base64Url 编码。
  2. 使用指定的签名算法（如 `HMACSHA256`）和密钥（`Secret`）生成签名。

  **生成签名的公式**：

  ```scss
  HMACSHA256(
    base64UrlEncode(header) + "." + base64UrlEncode(payload), 
    secret
  )
  ```

  签名确保服务器可以验证 JWT 的完整性和真实性，防止数据被篡改。

* **工作流程**

![image-20250222010204219](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250222010204219.png)

**（1）用户登录**：用户通过客户端向服务器发送 POST 请求，路径为 `/user/login`，请求体中包含用户名和密码。

**（2）服务器验证**：服务器验证用户的登录信息。如果验证成功，服务器生成一个包含用户身份（如用户 ID、用户名等）和权限信息的 JWT，并使用私钥对其进行签名。

**（3）返回 JWT**：服务器将生成的 JWT 返回给客户端，客户端可以将其存储在本地存储（如 `localStorage`、`sessionStorage` 或 `Cookie`）中。

**（4）客户端携带 JWT 请求资源**：在后续请求中，客户端将存储的 JWT 附加到 HTTP 请求头的 `Authorization` 字段中，格式为 `Authorization: Bearer <JWT>`。

**（5）服务器验证 JWT**：服务器接收到带有 JWT 的请求后，解析并验证其签名。若签名有效且有效载荷未被篡改，服务器根据 JWT 中的信息（如用户身份、权限等）判断是否允许访问资源。

**（6）返回响应**：服务器根据验证后的信息执行相关操作，并返回相应的结果或数据。

**5.四者比较**

| 特性           | Cookie                       | Session                            | Token                              | JWT                                    |
| :------------- | :--------------------------- | :--------------------------------- | :--------------------------------- | :------------------------------------- |
| **存储位置**   | 客户端                       | 服务器端                           | 客户端                             | 客户端                                 |
| **安全性**     | 较低                         | 较高                               | 较高                               | 较高                                   |
| **跨域支持**   | 默认不支持，可通过设置实现   | 不支持                             | 支持                               | 支持                                   |
| **大小限制**   | 约4KB                        | 无限制                             | 无限制                             | 受JSON大小限制                         |
| **生命周期**   | 可设置过期时间               | 通常在用户关闭浏览器或超时后失效   | 可设置过期时间，过期后需要重新获取 | 可设置过期时间，过期后需重新认证       |
| **无状态支持** | 不支持，依赖于Cookie         | 不支持，依赖于Session              | 支持，但Session需基于Cookie        | 支持，服务器端无状态                   |
| **优点**       | 容易在请求间使用             | 比 Cookies 安全                    | 无需服务器端存储，适合分布式系统   | 自包含，适合跨域认证                   |
| **缺点**       | 安全性低，容易被窃取或篡改   | 增加服务器负担，分布式系统扩展困难 | 如果存储不安全，容易被窃取         | 如果密钥泄露，存在安全风险             |
| **主要用途**   | 存储会话数据，如登录状态     | 存储服务器端的会话数据             | 用于身份验证和授权                 | 用于身份验证和授权，是 一种特殊的Token |
| **适用场景**   | 简单的会话跟踪、用户偏好设置 | 需要服务器记住用户状态的场景       | 移动应用、API身份验证、跨域请求    | Web应用、移动应用、单点登录            |

## 参考资料

请查阅阶段八任务计划思路