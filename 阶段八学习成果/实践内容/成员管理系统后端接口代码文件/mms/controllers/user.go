// 声明当前文件属于controllers包
package controllers

// 导入其他包或模块
import (
	"github.com/gin-gonic/gin"	//Gin框架，构建HTTP Web服务
	"github.com/golang-jwt/jwt/v4" //JWT库，生成和验证JSON Web Token
	"mms/config" //项目的配置包，配置数据库连接
	"mms/models" //项目的模型包，定义数据库模型
	"mms/utils" //项目的工具包，包含密码加密和验证等功能
	"net/http" //HTTP包，处理HTTP请求和响应
	"time" //时间包，处理时间相关功能
)

// 用户注册

// RegisterRequest结构体定义用户注册时的请求结构
//json:"..." 是结构体标签，指定将数据从JSON格式绑定到Go结构体时的字段名称
//binding:"required" 表示这个字段是必填项
type RegisterRequest struct {
	UserID   string `json:"user_id" binding:"required"` //用户名，必填
	Password string `json:"password" binding:"required"` //密码，必填
	Nickname string `json:"nickname"` //昵称，可选
	Phone    string `json:"phone"` //手机号，可选
	Email    string `json:"email"` //邮箱，可选
}

// Register函数处理用户注册请求
//Register函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)用于处理请求和生成响应
//gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法，是一个容器，装满了处理请求和生成响应的工具
//gin.Context作用：
//1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
//2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
//3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
func Register(c *gin.Context) {
	
	// 声明变量req，类型是RegisterRequest，存储客户端传递的用户注册请求的数据
	var req RegisterRequest
	
	// 绑定请求数据
	//c.ShouldBindJSON(&req)：调用Gin框架的ShouldBindJSON方法,将请求体中的JSON数据绑定到req结构体
	//if err := c.ShouldBindJSON(&req); err != nil：
	//将ShouldBindJSON的返回值赋给err，并检查err是否为nil,如果err不为nil，表示绑定失败,err会包含具体的错误信息
	if err := c.ShouldBindJSON(&req); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})：返回一个JSON响应，状态码为400 Bad Request，响应体包含错误信息
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusBadRequest：HTTP状态码,表示请求错误(400 Bad Request)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		//err.Error()：调用error接口的Error方法，获取具体的错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return //结束当前函数的执行，返回调用者
	}

	// 使用bcrypt算法加密密码
	//确保密码在数据库中以加密形式存储
	//utils.HashPassword(req.Password):调用utils包中的HashPassword函数，对req.Password进行加密
	//hashedPassword：存储加密后的密码
	//err：存储可能的错误信息
	hashedPassword, err := utils.HashPassword(req.Password)
	//检查err是否为nil，如果err不为nil，表示加密过程中发生了错误
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})：返回一个JSON响应，状态码为500 Internal Server Error，响应体包含错误信息(密码加密失败)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusInternalServerError：HTTP状态码，表示服务器内部错误(500 Internal Server Error)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return //结束当前函数的执行，返回调用者
	}

	// 创建用户对象
	//创建一个User模型实例(用户对象)user，用来存储用户信息
	user := models.User{
		UserID:    req.UserID, //用户名，从请求数据中获取
		Password:  hashedPassword,//密码，为加密后的密码
		Nickname:  req.Nickname, //昵称，从请求数据中获取
		Phone:     req.Phone, //手机号，从请求数据中获取
		Email:     req.Email, //邮箱，从请求数据中获取
		IsApproved: false, //默认设置为 false，表示用户未审核
		IsAdmin:    false, //默认设置为 false，表示用户不是管理员
	}

	// 将用户对象保存到数据库
	//result := config.DB.Create(&user)：调用GORM提供的Crerate方法将user(用户对象)保存到数据库，返回一个*gorm.Result对象，存储在result中
	//config.DB是一个config包里的全局变量DB，类型是 *gorm.DB，表示GORM的数据库连接实例
	//Create(&user)是GORM提供的一个方法，表示将user对象保存到数据库中
	//result是Create方法的返回值，类型是*gorm.Result，*gorm.Result 是GORM提供的一个结构体，存储查询结果和错误信息
	//result.Error是*gorm.Result结构体中的一个字段，类型是error，存储可能的错误信息
	//if result.Error != nil：检查保存是否成功，如果result.Error不为nil，表示保存过程中发生了错误
	if result := config.DB.Create(&user); result.Error != nil {
		//c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})：返回一个JSON响应，状态码为409 Conflict，响应体包含错误信息(用户已存在)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusConflict：HTTP状态码，表示冲突(409 Conflict)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
		return //结束当前函数的执行，返回调用者
	}
	//c.JSON(http.StatusCreated, gin.H{"message": "注册成功，等待管理员审核"}):返回一个JSON响应，状态码为201 Created，响应体包含资源创建成功信息(注册成功，等待管理员审核)
	//c.JSON：Gin框架的方法，用于返回JSON响应
	//http.StatusCreated:HTTP状态码,表示资源创建成功(201 Created)
	//gin.H{}：Gin的哈希映射，用于生成JSON响应体
	c.JSON(http.StatusCreated, gin.H{"message": "注册成功，等待管理员审核"})
}

// 用户登录

// LoginRequest结构体定义用户登录时的请求结构
type LoginRequest struct {
	UserID   string `json:"user_id" binding:"required"` //用户名，必填
	Password string `json:"password" binding:"required"` //密码，必填
}

// Login函数处理用户登录请求
//Login函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)，用于处理请求和生成响应
//gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法，是一个容器，装满了处理请求和生成响应的工具
//gin.Context作用：
//1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
//2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
//3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
func Login(c *gin.Context) {

	// 声明变量req，类型是LoginRequest，存储客户端传递的用户登录请求的数据
	var req LoginRequest
	
	// 绑定请求数据
	//c.ShouldBindJSON(&req)：调用Gin框架的ShouldBindJSON方法,将请求体中的JSON数据绑定到req结构体
	//if err := c.ShouldBindJSON(&req); err != nil：
	//将ShouldBindJSON的返回值赋给err，并检查err是否为nil,如果err不为nil，表示绑定失败,err会包含具体的错误信息
	if err := c.ShouldBindJSON(&req); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})：返回一个JSON响应，状态码为400 Bad Request，响应体包含错误信息
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusBadRequest：HTTP状态码,表示请求错误(400 Bad Request)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		//err.Error()：调用error接口的Error方法，获取具体的错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return //结束当前函数的执行，返回调用者
	}

	// 创建用户对象
	//创建一个User模型实例(用户对象)user，用来存储用户信息(声明一个User类型的变量user，它是一个结构体实例，可以存储用户信息)
	var user models.User

	// 查询数据库，检查请求登录的用户是否存在
	//result := config.DB.Where("user_id = ?", req.UserID).First(&user):
	//调用GORM提供的First方法查询数据库，检查是否存在user_id等于req.UserID的用户，将查询结果存储到user变量，并将查询结果返回给result
	//config.DB是一个config包里的全局变量DB，类型是 *gorm.DB，表示GORM的数据库连接实例
	//Where("user_id = ?", req.UserID)是GORM的查询条件，表示查询user_id等于req.UserID的记录
	//user_id：数据库表中的字段名，req.UserID：从请求中获取的用户名
	//First(&user)：GORM提供的查询方法，查询第一条匹配的记录，并将结果存储到user变量中
	//result是First方法的返回值，类型是*gorm.Result，*gorm.Result 是GORM提供的一个结构体，存储查询结果和错误信息
	//result.Error是*gorm.Result结构体中的一个字段，类型是error，存储可能的错误信息
	//if result.Error != nil：检查保存是否成功，如果result.Error不为nil，表示查询过程中发生了错误
	if result := config.DB.Where("user_id = ?", req.UserID).First(&user); result.Error != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})：返回一个JSON响应，状态码为404 Not Found，响应体包含错误信息(用户不存在)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusNotFound：HTTP状态码，表示未找到资源(404 Not Found)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return //结束当前函数的执行，返回调用者
	}

	// 检查请求登录的用户是否已通过审核
	//!user.IsApproved是对user对象中的IsApproved字段取反，
	//如果user.IsApproved为false，则!user.IsApproved为true，执行后续代码
	if !user.IsApproved {
		//c.JSON(http.StatusForbidden, gin.H{"error": "账户未通过审核"})：返回一个JSON响应，状态码为403 Forbidden，响应体包含错误信息(账户未通过审核)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusConflict：HTTP状态码，表示禁止访问(403 Forbidden)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusForbidden, gin.H{"error": "账户未通过审核"})
		return //结束当前函数的执行，返回调用者
	}

	// 验证用户输入的密码是否正确
	//utils.CheckPassword()：调用utils包中的CheckPassword方法，验证用户输入的密码是否正确
	//req.Password：用户输入的密码，user.Password：数据库中存储的加密密码
	//如果密码不正确，!utils.CheckPassword(req.Password, user.Password) 为true，执行后续代码
	if !utils.CheckPassword(req.Password, user.Password) {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "账户未通过审核"})：返回一个JSON响应，状态码为401 Unauthorized，响应体包含错误信息(密码错误)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusUnauthorized：HTTP状态码，表示未授权(401 Unauthorized)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return //结束当前函数的执行，返回调用者
	}

	// 生成JWT token
	//jwt.NewWithClaims()是JWT库中的一个函数，用于创建一个新的JWT token，指定签名方法和声明
	//jwt.SigningMethodHS256：指定使用HMAC SHA256签名方法，jwt.MapClaims：创建一个包含声明的映射表，用于存储token中的用户信息
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//"user_id": user.UserID：从user对象中获取用户名，将用户名存储在JWT token中
		//user.UserID：从user对象中获取用户名，"user_id"：在JWT token中存储用户名
		//通过user_id字段，服务器可以在后续请求中识别用户
		"user_id": user.UserID,
		//"admin":   user.IsAdmin：从user对象中获取IsAdmin，将IsAdmin存储在JWT token中
		//user.IsAdmin：从user对象中获取IsAdmin，"admin"：在JWT token中存储IsAdmin
		//通过admin字段，服务器可以在后续请求中判断用户是否为管理员
		"admin":   user.IsAdmin, 
		//"exp":     time.Now().Add(time.Hour * 72).Unix()：设置JWT token的过期时间，确保token在72小时后失效
		//time.Now()：调用时间库里的Now方法，获取当前时间
		//Add(time.Hour * 72)：将当前时间加上 72 小时(3天)
		//Unix() 返回一个Unix时间戳(秒级)，表示到期时间
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	//token.SignedString():调用token中的SignedString方法，使用指定的密钥对JWT token进行签名，生成token字符串
	//token：是一个jwt.Token类型的对象，通过jwt.NewWithClaims()方法创建
	//SignedString：是jwt.Token类型的一个方法，用于将Token签名并生成一个字符串形式的Token
	//"your-secret-key" 是示例密钥，[]byte()：将密钥转换为字节切片，用于签名
	//tokenString, _：将生成的token字符串存储在tokenString变量中，忽略错误
	tokenString, _ := token.SignedString([]byte("your-secret-key")) // 使用合适的密钥
	//c.JSON(http.StatusOK, gin.H{"token": tokenString})：返回一个JSON响应，状态码为200 OK，响应体包含生成的token字符串(tokenString)
	//c.JSON：Gin框架的方法，用于返回JSON响应
	//http.StatusUnauthorized：HTTP状态码，表示请求成功(200 OK)
	//gin.H{}：Gin的哈希映射，用于生成JSON响应体
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


// UpdateUser函数处理更新用户信息的请求
//UpdateUser函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)，用于处理请求和生成响应
//gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法，是一个容器，装满了处理请求和生成响应的工具
//gin.Context作用：
//1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
//2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
//3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
func UpdateUser(c *gin.Context) {

    // 从上下文中获取当前用户的信息
	//c.Get("user")：从Gin框架的上下文对象c中获取键为"user"的值，
	//currentUser：存储获取到的值(存储用户信息)，类型为interface{}
	//exists：布尔值，表示该键是否存在(表示用户信息是否存在)
    currentUser, exists := c.Get("user")

	// 验证当前用户是否已通过认证
	//如果exists为false，表示当前用户未通过认证,!exists为true，执行后续代码
	if !exists {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "用户为认证"})：返回一个JSON响应，状态码为401 Unauthorized，响应体包含错误信息(用户未认证)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusUnauthorized：HTTP状态码，表示未授权(401 Unauthorized)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
        c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
        return //结束当前函数的执行，返回调用者
    }
	//将currentUser转换为models.User类型，并赋值给user，之后user表示当前用户的信息，类型为models.User
	//将currentUser转换为models.User类型，是为了方便后续操作，因为currentUser是一个interface{}类型的变量，需要通过类型断言转换为具体的类型
	//类型断言后，currentUser是models.User类型，我们便可以对user进行后续操作，比如更新用户信息等
    user := currentUser.(models.User)

    // 绑定请求数据到临时结构体（避免覆盖敏感字段）
	//updateData结构体定义更新新时的请求结构
	//json:"..." 是结构体标签，指定将数据从JSON格式绑定到Go结构体时的字段名称
    var updateData struct {
        Nickname *string `json:"nickname"` //昵称，可选
        Phone    *string `json:"phone"`	//手机号，可选
        Email    *string `json:"email"` //邮箱，可选
    }
	// 绑定请求数据
	//c.ShouldBindJSON(&updateData)：调用Gin框架的ShouldBindJSON方法,将请求体中的JSON数据绑定到updateData结构体
	//if err := c.ShouldBindJSON(&updateData); err != nil：
	//将ShouldBindJSON的返回值赋给err，并检查err是否为nil,如果err不为nil，表示绑定失败,err会包含具体的错误信息
    if err := c.ShouldBindJSON(&updateData); err != nil {
        //c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})：返回一个JSON响应，状态码为400 Bad Request，响应体包含错误信息
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusBadRequest：HTTP状态码,表示请求错误(400 Bad Request)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		//err.Error()：调用error接口的Error方法，获取具体的错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return //结束当前函数的执行，返回调用者
    }

    // 部分更新逻辑
	//检查请求中是否提供了Nickname字段，如果提供了，更新user.Nickname
	//这样可以确保只有请求中提供的字段被更新，避免覆盖其他字段
	//updateData.Nickname：从updateData结构体中获取Nickname字段
	//如果updateData.Nickname不为nil，说明请求中提供了Nickname字段,执行后续代码
    if updateData.Nickname != nil {
		//解引用updateData.Nickname，获取其值，赋给user.Nickname
		//作用：更新user对象的Nickname字段
        user.Nickname = *updateData.Nickname
    }
	//检查请求中是否提供了Phone字段，如果提供了，更新user.Phone
	//这样可以确保只有请求中提供的字段被更新，避免覆盖其他字段
	//updateData.Phone：从updateData结构体中获取Phone字段
	//如果updateData.Phone不为nil，说明请求中提供了Phone字段,执行后续代码
    if updateData.Phone != nil {
		//解引用updateData.Phone，获取其值，赋给user.Phone
		//作用：更新user对象的Phone字段
        user.Phone = *updateData.Phone
    }
	//检查请求中是否提供了Email字段，如果提供了，更新user.Email
	//这样可以确保只有请求中提供的字段被更新，避免覆盖其他字段
	//updateData.Email：从updateData结构体中获取Email字段
	//如果updateData.Email不为nil，说明请求中提供了Email字段,执行后续代码
    if updateData.Email != nil {
		//解引用updateData.Email，获取其值，赋给user.Email
		//作用：更新user对象的Email字段
        user.Email = *updateData.Email
    }

    // 保存到数据库
	//result := config.DB.Save(&user):
	//调用GORM提供的Save方法，将user对象保存到数据库，并将操作结果返回给result
	//config.DB是一个config包里的全局变量DB，类型是 *gorm.DB，表示GORM的数据库连接实例
	//First(&user)：GORM提供的保存方法，将更新后的user对象保存到数据库
	//result是Save方法的返回值，类型是*gorm.Result，*gorm.Result 是GORM提供的一个结构体，存储查询结果和错误信息
	//result.Error是*gorm.Result结构体中的一个字段，类型是error，存储可能的错误信息
	//if result.Error != nil：检查保存是否成功，如果result.Error不为nil，表示保存过程中发生了错误
    if result := config.DB.Save(&user); result.Error != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})：返回一个JSON响应，状态码为500 Internal Server Error，响应体包含错误信息(更新失败)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusInternalServerError：HTTP状态码，表示服务器内部错误(500 Internal Server Error)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
        c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
        return //结束当前函数的执行，返回调用者
    }
	//c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功", "user": user})：返回一个JSON响应，状态码为200 OK，响应体包含成功消息和更新后的用户信息
	//c.JSON：Gin框架的方法，用于返回JSON响应
	//http.StatusUnauthorized：HTTP状态码，表示请求成功(200 OK)
	//gin.H{}：Gin的哈希映射，用于生成JSON响应体
    c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功", "user": user})
}

// DeleteUser函数处理删除用户请求
//DeleteUser函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)，用于处理请求和生成响应
//gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法，是一个容器，装满了处理请求和生成响应的工具
//gin.Context作用：
//1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
//2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
//3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
func DeleteUser(c *gin.Context) {
	// 从URL中获取用户ID
	//userID := c.Param("id")：从URL中获取用户名，存储在变量userID中
	//c：Gin框架的上下文对象(*gin.Context)，用于处理请求和生成响应
	//Param：*gin.Context对象的一个方法，用于从URL中获取参数值
	//"id"：参数的键名，表示从URL中获取名为id的参数值
	userID := c.Param("id") 

	// 强制删除用户记录（物理删除）
	//result := config.DB.Unscoped().Where("user_id = ?", userID).Delete(&models.User{}):
	//调用GORM提供的Unscoped方法取消 GORM 的默认软删除行为，物理删除user_id等于userID的用户，并将操作结果返回给result
	//config.DB是一个config包里的全局变量DB，类型是 *gorm.DB，表示GORM的数据库连接实例
	//config.DB.Unscoped()：调用GORM的Unscoped方法，取消GORM的默认软删除行为，进行物理删除
	//Where("user_id = ?", userID)：指定删除条件，表示删除user_id等于userID的用户
	//Delete(&models.User{})：执行删除操作，删除匹配条件的用户记录
	//result是Delete方法的返回值，类型是*gorm.Result，*gorm.Result 是GORM提供的一个结构体，存储操作结果和错误信息
	//result.Error是*gorm.Result结构体中的一个字段，类型是error，存储可能的错误信息
	//if result.Error != nil：检查保存是否成功，如果result.Error不为nil，表示删除过程中发生了错误
	if result := config.DB.Unscoped().Where("user_id = ?", userID).Delete(&models.User{}); result.Error != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})：返回一个JSON响应，状态码为404 Not Found，响应体包含错误信息(用户未找到)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusNotFound：HTTP状态码，表示未找到资源(404 Not Found)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return //结束当前函数的执行，返回调用者
	}
	//c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})：返回一个JSON响应，状态码为200 OK，响应体包含操作成功消息
	//c.JSON：Gin框架的方法，用于返回JSON响应
	//http.StatusOK：HTTP状态码，表示请求成功(200 OK)
	//gin.H{}：Gin的哈希映射，用于生成JSON响应体
	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}


// GetUser函数处理获取单个用户信息请求
//GetUser函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)，用于处理请求和生成响应
//gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法，是一个容器，装满了处理请求和生成响应的工具
//gin.Context作用：
//1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
//2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
//3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
func GetUser(c *gin.Context) {
	
	// 从URL中获取用户ID
	//userID := c.Param("id")：从URL中获取用户名，存储在变量userID中
	//c：Gin框架的上下文对象(*gin.Context)，用于处理请求和生成响应
	//Param：*gin.Context对象的一个方法，用于从URL中获取参数值
	//"id"：参数的键名，表示从URL中获取名为id的参数值
	userID := c.Param("id")
	
	// 创建用户对象
	//创建一个User模型实例(用户对象)user，用来存储用户信息(声明一个User类型的变量user，它是一个结构体实例，可以存储用户信息)
	var user models.User

	// 查询数据库，检查请求获取的用户是否存在
	//result := config.DB.Where("user_id = ?", req.UserID).First(&user):
	//调用GORM提供的First方法查询数据库，检查是否存在user_id等于req.UserID的用户，将查询结果存储到user变量，并将查询结果返回给result
	//config.DB是一个config包里的全局变量DB，类型是 *gorm.DB，表示GORM的数据库连接实例
	//Where("user_id = ?", req.UserID)是GORM的查询条件，表示查询user_id等于req.UserID的记录
	//user_id：数据库表中的字段名，req.UserID：从请求中获取的用户名
	//First(&user)：GORM提供的查询方法，查询第一条匹配的记录，并将结果存储到user变量中
	//result是First方法的返回值，类型是*gorm.Result，*gorm.Result 是GORM提供的一个结构体，存储查询结果和错误信息
	//result.Error是*gorm.Result结构体中的一个字段，类型是error，存储可能的错误信息
	//if result.Error != nil：检查保存是否成功，如果result.Error不为nil，表示查询过程中发生了错误
	if result := config.DB.Where("user_id = ?", userID).First(&user); result.Error != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})：返回一个JSON响应，状态码为404 Not Found，响应体包含错误信息(用户未找到)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusNotFound：HTTP状态码，表示未找到资源(404 Not Found)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return //结束当前函数的执行，返回调用者
	}

	//c.JSON(http.StatusOK, gin.H{"user": user})：返回一个JSON响应，状态码为200 OK，响应体包含操作成功消息
	//c.JSON：Gin框架的方法，用于返回JSON响应
	//http.StatusOK：HTTP状态码，表示请求成功(200 OK)
	//gin.H{}：Gin的哈希映射，用于生成JSON响应体
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetAllUsers函数处理获取所有用户信息请求
//GetAllUsers函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)，用于处理请求和生成响应
//gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法，是一个容器，装满了处理请求和生成响应的工具
//gin.Context作用：
//1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
//2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
//3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
func GetAllUsers(c *gin.Context) {

	// 创建所有用户对象
	//声明一个models.User类型的切片变量users，用于存储查询到的所有用户信息
	var users []models.User

	// 查询数据库，检查请求获取的所有用户是否存在
	//result := config.DB.Find(&users)：
	//调用GORM提供的Find方法查询数据库，查询数据库中的所有用户信息，并将查询结果存储到users切片中，然后返回给result
	//config.DB是一个config包里的全局变量DB，类型是 *gorm.DB，表示GORM的数据库连接实例
	//First(&user)：GORM提供的查询方法，查询数据库中的所有用户信息，并将查询结果存储到users切片中
	//result是Find方法的返回值，类型是*gorm.Result，*gorm.Result 是GORM提供的一个结构体，存储查询结果和错误信息
	//result.Error是*gorm.Result结构体中的一个字段，类型是error，存储可能的错误信息
	//if result.Error != nil：检查保存是否成功，如果result.Error不为nil，表示查询过程中发生了错误
	if result := config.DB.Find(&users); result.Error != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户失败"})：返回一个JSON响应，状态码为500 Internal Server Error，响应体包含错误信息(获取用户失败)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusInternalServerError：HTTP状态码，表示服务器内部错误(500 Internal Server Error)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户失败"})
		return //结束当前函数的执行，返回调用者
	}

	//c.JSON(http.StatusOK, gin.H{"user": user})：返回一个JSON响应，状态码为200 OK，响应体包含操作成功消息
	//c.JSON：Gin框架的方法，用于返回JSON响应
	//http.StatusOK：HTTP状态码，表示请求成功(200 OK)
	//gin.H{}：Gin的哈希映射，用于生成JSON响应体
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// ApproveUser函数处理审核用户请求
//ApproveUser函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)，用于处理请求和生成响应
//gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法，是一个容器，装满了处理请求和生成响应的工具
//gin.Context作用：
//1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
//2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
//3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
func ApproveUser(c *gin.Context) {

	// 从URL中获取用户ID
	//userID := c.Param("id")：从URL中获取用户名，存储在变量userID中
	//c：Gin框架的上下文对象(*gin.Context)，用于处理请求和生成响应
	//Param：*gin.Context对象的一个方法，用于从URL中获取参数值
	//"id"：参数的键名，表示从URL中获取名为id的参数值
	userID := c.Param("id") 

	// 创建用户对象
	//创建一个User模型实例(用户对象)user，用来存储用户信息(声明一个User类型的变量user，它是一个结构体实例，可以存储用户信息)
	var user models.User

	
	// 查询数据库，检查请求获取的用户是否存在
	//result := config.DB.Where("user_id = ?", req.UserID).First(&user):
	//调用GORM提供的First方法查询数据库，检查是否存在user_id等于req.UserID的用户，将查询结果存储到user变量，并将查询结果返回给result
	//config.DB是一个config包里的全局变量DB，类型是 *gorm.DB，表示GORM的数据库连接实例
	//Where("user_id = ?", req.UserID)是GORM的查询条件，表示查询user_id等于req.UserID的记录
	//user_id：数据库表中的字段名，req.UserID：从请求中获取的用户名
	//First(&user)：GORM提供的查询方法，查询第一条匹配的记录，并将结果存储到user变量中
	//result是First方法的返回值，类型是*gorm.Result，*gorm.Result 是GORM提供的一个结构体，存储查询结果和错误信息
	//result.Error是*gorm.Result结构体中的一个字段，类型是error，存储可能的错误信息
	//if result.Error != nil：检查保存是否成功，如果result.Error不为nil，表示查询过程中发生了错误
	if result := config.DB.Where("user_id = ?", userID).First(&user); result.Error != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})：返回一个JSON响应，状态码为404 Not Found，响应体包含错误信息(用户未找到)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusNotFound：HTTP状态码，表示未找到资源(404 Not Found)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return //结束当前函数的执行，返回调用者
	}

	// 修改用户审核状态为通过
	//将user对象的IsApproved字段设置为true
	user.IsApproved = true
	//result := config.DB.Save(&user):
	//调用GORM提供的Save方法，将user对象保存到数据库，并将操作结果返回给result
	//config.DB是一个config包里的全局变量DB，类型是 *gorm.DB，表示GORM的数据库连接实例
	//First(&user)：GORM提供的保存方法，将更新后的user对象保存到数据库
	//result是Save方法的返回值，类型是*gorm.Result，*gorm.Result 是GORM提供的一个结构体，存储查询结果和错误信息
	//result.Error是*gorm.Result结构体中的一个字段，类型是error，存储可能的错误信息
	//if result.Error != nil：检查保存是否成功，如果result.Error不为nil，表示保存过程中发生了错误
	if result := config.DB.Save(&user); result.Error != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户失败"})：返回一个JSON响应，状态码为500 Internal Server Error，响应体包含错误信息(获取用户失败)
		//c.JSON：Gin框架的方法，用于返回JSON响应
		//http.StatusInternalServerError：HTTP状态码，表示服务器内部错误(500 Internal Server Error)
		//gin.H{}：Gin的哈希映射，用于生成JSON响应体
		c.JSON(http.StatusInternalServerError, gin.H{"error": "审核失败"})
		return //结束当前函数的执行，返回调用者
	}

	//c.JSON(http.StatusOK, gin.H{"message": "用户审核通过", "user": user})：返回一个JSON响应，状态码为200 OK，响应体包含操作成功消息
	//c.JSON：Gin框架的方法，用于返回JSON响应
	//http.StatusOK：HTTP状态码，表示请求成功(200 OK)
	//gin.H{}：Gin的哈希映射，用于生成JSON响应体
	c.JSON(http.StatusOK, gin.H{"message": "用户审核通过", "user": user})
}
