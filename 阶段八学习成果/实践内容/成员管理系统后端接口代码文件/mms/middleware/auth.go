// 声明当前文件属于middleware包
package middleware

// 导入其他包或模块
import (
	"github.com/gin-gonic/gin" //Gin框架，构建HTTP Web服务
	"github.com/golang-jwt/jwt/v4" //JWT库，生成和验证JSON Web Token
	"mms/config" //项目的配置包，配置数据库连接
	"mms/models" //项目的模型包，定义数据库模型
	"net/http" //HTTP包，处理HTTP请求和响应
	"strings" //字符串操作包
	
)
// 定义一个中间件函数，用于处理用户认证
//AuthMiddleware函数返回一个gin.HandlerFunc类型的中间件
//gin.HandlerFunc是Gin框架中的中间件函数类型，用于定义中间件的逻辑
func AuthMiddleware() gin.HandlerFunc {
    
    // 返回一个匿名函数，该函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)，用于处理请求和生成响应
    //gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法
    //gin.Context作用：
    //1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
    //2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
    //3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
    return func(c *gin.Context) {
        //从Gin框架的上下文对象c中获取HTTP请求头中的Authorization字段的值，存储在变量tokenString中
        //c：Gin框架的上下文对象，用于处理请求和生成响应
        //"Authorization"：HTTP请求头中的字段名，通常用于携带用户认证信息，如JWT令牌
        //比如tokenString为 "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
        tokenString := c.GetHeader("Authorization")
        
        //如果tokenString为空，表示客户端未提供认证令牌，执行后续代码
        if tokenString == "" {
            //c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})：返回一个JSON响应，状态码为401 Unauthorized，响应体包含错误信息(未提供认证令牌)
		    //c.AbortWithStatusJSON：Gin框架的方法，用于终止请求处理并返回 JSON 响应
		    //http.StatusUnauthorized：HTTP状态码，表示未授权(401 Unauthorized)
		    //gin.H{}：Gin的哈希映射，用于生成JSON响应体
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
            return //结束当前函数的执行，返回调用者
        }

        // 去掉 "Bearer " 前缀
        //调用strings包的TrimPrefix方法，去掉tokenString中的前缀 "Bearer "，并将去掉前缀后的令牌字符串存储在tokenString中
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // 解析令牌并验证有效性
        //jwt.Parse方法解析令牌字符串并验证其令牌签名的有效性，如果签名有效，token对象将包含解析后的令牌信息，如果签名无效，err将包含错误信息
        //jwt.Parse：调用jwt包中的Parse方法，解析JWT令牌
        //tokenString：要解析的JWT令牌字符串
        //func(token *jwt.Token) (interface{}, error)：解析函数，用于验证令牌的签名
        //token *jwt.Token：解析后的JWT令牌对象
        //interface{}：返回一个接口类型，用于存储签名验证的密钥
        //error：返回一个错误值，表示解析过程中是否发生错误
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            //[]byte("your-secret-key")：将字符串"your-secret-key"转换为字节切片，作为签名验证的密钥
            //nil：表示没有错误发生
            return []byte("your-secret-key"), nil 
        })
        
        //err != nil表明解析令牌时发生错误，!token.Valid表明令牌无效
        //如果解析令牌时发生错误或令牌无效，执行后续的代码
        if err != nil || !token.Valid {
            //c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效令牌"})：返回一个JSON响应，状态码为401 Unauthorized，响应体包含错误信息(无效令牌)
		    //c.AbortWithStatusJSON：Gin框架的方法，用于终止请求处理并返回 JSON 响应
		    //http.StatusUnauthorized：HTTP状态码，表示未授权(401 Unauthorized)
		    //gin.H{}：Gin的哈希映射，用于生成JSON响应体
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效令牌"})
        
            return //结束当前函数的执行，返回调用者
        }

        // 提取声明并验证user_id类型
        //token.Claims：token对象中的Claims字段，表示JWT令牌中的声明(claims)
        //Claims是一个接口类型，可以是jwt.MapClaims或其他类型
        //jwt.MapClaims：jwt包中的一个类型，表示一个映射表形式的声明
        //MapClaims 是一个 map[string]interface{} 类型的映射表，用于存储JWT令牌中的声明
        //claims, ok := token.Claims.(jwt.MapClaims)：
        //获取token对象中的Claims字段断言为jwt.MapClaims类型
        //如果断言成功，ok 为 true，claims被赋予token.Claims的值，类型为jwt.MapClaims
        //如果断言失败，ok 为 false，claims没有被赋予token.Claims的值，类型不为jwt.MapClaims
        claims, ok := token.Claims.(jwt.MapClaims)
        //如果断言失败，ok 为 false，!ok为 true，执行后续的代码
        if !ok {
            //c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "令牌声明格式错误"})：返回一个JSON响应，状态码为401 Unauthorized，响应体包含错误信息(令牌声明格式错误)
		    //c.AbortWithStatusJSON：Gin框架的方法，用于终止请求处理并返回 JSON 响应
		    //http.StatusUnauthorized：HTTP状态码，表示未授权(401 Unauthorized)
		    //gin.H{}：Gin的哈希映射，用于生成JSON响应体
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "令牌声明格式错误"})
            return //结束当前函数的执行，返回调用者
        }

        //claims["user_id"]：从声明中获取user_id字段
        //userID, ok := claims["user_id"].(string)：
        //从声明中获取user_id字段，断言为string类型
        //如果断言成功，ok 为 true，userID被赋予user_id字段的值，类型为string
        //如果断言失败，ok 为 false，userID没有被赋予user_id字段的值，类型不为string
        userID, ok := claims["user_id"].(string)
        //如果断言失败，ok 为 false，!ok为 true，或user_id 为空，执行后续的代码
        if !ok || userID == "" {
            //c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "令牌中缺少用户ID"})：返回一个JSON响应，状态码为401 Unauthorized，响应体包含错误信息(令牌中缺少用户ID)
		    //c.AbortWithStatusJSON：Gin框架的方法，用于终止请求处理并返回 JSON 响应
		    //http.StatusUnauthorized：HTTP状态码，表示未授权(401 Unauthorized)
		    //gin.H{}：Gin的哈希映射，用于生成JSON响应体
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "令牌中缺少用户ID"})
            return //结束当前函数的执行，返回调用者
        }

        // 查询用户信息
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
            //c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})：返回一个JSON响应，状态码为404 Not Found，响应体包含错误信息(用户不存在)
		    //c.AbortWithStatusJSON：Gin框架的方法，用于终止请求处理并返回 JSON 响应
		    //http.StatusUnauthorized：HTTP状态码，表示未找到资源(404 Not Found)
		    //gin.H{}：Gin的哈希映射，用于生成JSON响应体
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
            return //结束当前函数的执行，返回调用者
        }

        // 存储用户对象到上下文
        //将用户信息存储到上下文对象c中，键名为"user"，值为user对象
        //c：Gin框架的上下文对象，用于处理请求和生成响应
        //c.Set：Gin框架的方法，用于将数据存储到上下文对象c中
        //"user"：键名，表示存储的数据的名称
        //user：值，表示要存储的数据，即用户信息
        c.Set("user", user)
        //c：Gin框架的上下文对象，用于处理请求和生成响应
        //c.Next()：Gin框架的方法，用于继续执行后续的中间件或处理函数
        c.Next()
    }
}

// 定义一个中间件函数，用于处理用户认证
//AuthMiddleware函数返回一个gin.HandlerFunc类型的中间件
//gin.HandlerFunc是Gin框架中的中间件函数类型，用于定义中间件的逻辑
func AdminMiddleware() gin.HandlerFunc {
    // 返回一个匿名函数，该函数接收一个参数c，类型是*gin.Context(是一个指向gin.Context的指针或是一个*gin.Context实例)，用于处理请求和生成响应
    //gin.Context：Gin框架的核心结构体，包含请求和响应的所有信息，以及一些便捷的方法
    //gin.Context作用：
    //1.处理请求：gin.Context包含了HTTP请求的所有信息，比如请求头、请求体、路由参数等
    //2.生成响应：gin.Context提供了方法来生成 HTTP 响应，比如c.JSON、c.HTML、c.Redirect等
    //3.中间件支持：gin.Context支持中间件，可以在请求处理过程中执行一些通用逻辑，例如身份验证、日志记录等
    return func(c *gin.Context) {
        
        // 从上下文中获取用户信息
        //c.Get("user")：从 Gin 框架的上下文对象c中获取键为 "user" 的值
        //如果键"user"存在，exists为true，user存储获取到的值
        //如果键"user"不存在，exists为false，user不存储获取到的值
        user, exists := c.Get("user")
        
        //如果exists为false，表示键"user"不存在，那么!exists为true，执行后续的代码
        if !exists {
            //c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})：返回一个JSON响应，状态码为401 Unauthorized，响应体包含错误信息(用户未认证)
		    //c.AbortWithStatusJSON：Gin框架的方法，用于终止请求处理并返回 JSON 响应
		    //http.StatusUnauthorized：HTTP状态码，表示未授权(401 Unauthorized)
		    //gin.H{}：Gin的哈希映射，用于生成JSON响应体
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
            return //结束当前函数的执行，返回调用者
        }

        // 类型断言并检查管理员权限
        //user：变量名，表示从上下文对象c中获取的用户信息
        //models.User：models包中的User结构体，表示用户模型
        //user.(models.User)：将user断言为models.User类型，以便访问其字段和方法
        //IsAdmin：models.User结构体中的一个字段，true表示用户是管理员，false表示用户不是管理员
        //将user断言为models.User类型，访问里面的IsAdmin字段，如果IsAdmin字段为false，即用户不是管理员，执行后续的代码
        if !user.(models.User).IsAdmin {
            //c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})：返回一个JSON响应，状态码为403 Forbidden，响应体包含错误信息(需要管理员权限)
		    //c.AbortWithStatusJSON：Gin框架的方法，用于终止请求处理并返回 JSON 响应
		    //http.StatusUnauthorized：HTTP状态码，表示服务器理解请求但拒绝执行(403 Forbidden)
		    //gin.H{}：Gin的哈希映射，用于生成JSON响应体
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "需要管理员权限"})
            return //结束当前函数的执行，返回调用者
        }

        //c：Gin框架的上下文对象，用于处理请求和生成响应
        //c.Next()：Gin框架的方法，用于继续执行后续的中间件或处理函数
        c.Next()
    }
}
