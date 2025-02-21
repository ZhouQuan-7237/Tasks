// 声明当前文件属于routes包
package routes

// 导入其他包或模块
import (
	"mms/controllers" //项目的控制器包，处理与业务逻辑相关的操作
	"mms/middleware" //项目的中间件包，设置中间件
	"github.com/gin-gonic/gin" //Gin框架，构建HTTP Web服务
)

// 设置项目的路由
//SetupRouter设置项目的路由
//*gin.Engine 是Gin框架的路由引擎类型，用于定义和管理路由
//*SetupRouter() *gin.Engine会执行以下操作：
//1.创建一个 gin.Engine 实例（Gin 框架的核心对象）。
//2.在这个实例上定义所有的路由规则（例如 /register、/login 等接口）。
//3.返回这个配置好的 gin.Engine 实例
func SetupRouter() *gin.Engine {
    //gin.Default()：创建一个默认的Gin路由引擎实例,并赋值给r，用于定义和管理路由
    r := gin.Default()

    // 公共接口
    //定义一个POST请求的路由，路径为/register，处理函数为controllers.Register
    //r.POST：定义一个POST请求的路由
    //"/register"：路由路径，表示注册接口
    //controllers.Register：controllers包里的Register函数，用于处理注册请求
    r.POST("/register", controllers.Register)

    //定义一个POST请求的路由，路径为/login，处理函数为controllers包里的Login函数
    //r.POST：定义一个POST请求的路由
    //"/login"：路由路径，表示登录接口
    //controllers.Login：controllers包里的Login函数，用于处理登录请求
    r.POST("/login", controllers.Login)

    // 需要认证的接口
    //创建一个路由组，用于定义需要认证的接口
    //创建一个路由组，路径为 /，并赋值给auth
    //r：Gin路由引擎实例
    //Group：Gin提供的方法，用于创建路由组
    //auth：路由组实例
    auth := r.Group("/")
    
    //auth.Use为路由组添加中间件
    //auth：路由组实例
    //Use：Gin提供的方法，用于添加中间件
    //middleware.AuthMiddleware()：
    //调用middleware包中的AuthMiddleware函数，获取认证中间件
    //作用：为路由组中的所有接口添加认证中间件，确保只有认证通过的用户才能访问这些接口
    auth.Use(middleware.AuthMiddleware())
    {
        //定义一个PUT请求的路由，路径为/user/profile，处理函数为controllers包里的UpdateUser函数
        //auth.PUT：定义一个PUT请求的路由
        //auth：路由组实例
        //PUT：HTTP方法，表示更新操作
        //"/user/profile"：路由路径，表示更新用户信息的接口
        //controllers.UpdateUser：controllers包里的UpdateUser函数，用于处理更新用户信息的请求
        auth.PUT("/user/profile", controllers.UpdateUser) // 用户只能更新自己的信息
    }

    // 管理员接口
    //创建一个子路由组，用于定义管理员接口
    //在auth路由组中创建一个子路由组，路径为 /admin，并赋值给admin
    //auth：父路由组实例
    //Group：Gin提供的方法，用于创建子路由组
    //"/admin"：子路由组的路径
    admin := auth.Group("/admin")

    //auth.Use为路由组添加中间件
    //auth：路由组实例
    //Use：Gin提供的方法，用于添加中间件
    //middleware.AuthMiddleware()：
    //调用middleware包中的AuthMiddleware函数，获取认证中间件
    //作用：为路由组中的所有接口添加认证中间件，确保只有认证通过的用户才能访问这些接口
    admin.Use(middleware.AdminMiddleware())
    {
        //定义一个DELETE请求的路由，路径为 /admin/users/:id，处理函数为controllers包的DeleteUser函数
        //admin：子路由组实例
        //DELETE：HTTP方法，表示删除操作
        //"/users/:id"：路由路径，表示删除用户的接口
        //controllers.DeleteUser：controllers包的DeleteUser函数，用于处理删除用户的请求
        admin.DELETE("/users/:id", controllers.DeleteUser)
        
        //定义一个GET请求的路由，路径为/users/:id，处理函数为controllers包的GetUser函数
        //admin：子路由组实例
        //GET：HTTP方法，表示获取操作
        //"/users/:id"：路由路径，表示获取单个用户的接口
        //controllers.GetUser：controllers包的GetUser函数，用于处理获取单个用户的请求
        admin.GET("/users/:id", controllers.GetUser)

        //定义一个GET请求的路由，路径为/users，处理函数为controllers包的GetAllUsers函数
        //admin：子路由组实例
        //GET：HTTP方法，表示获取操作
        //"/users"：路由路径，表示获取所有用户的接口
        //controllers.GetAllUsers：controllers包的GetAllUsers函数，用于处理获取所有用户的请求
        admin.GET("/users", controllers.GetAllUsers)

        //定义一个PUT请求的路由，路径为/users/:id/approve:id，处理函数为controllers包的ApproveUser函数
        //admin：子路由组实例
        //PUT：HTTP方法，表示更新操作
        //"/users/:id/approve"：路由路径，表示更新用户的接口
        //controllers.ApproveUser：controllers包的ApproveUser函数，用于处理更新用户的请求
        admin.PUT("/users/:id/approve", controllers.ApproveUser)
    }

    //返回配置好的Gin路由引擎实例r
    return r
}