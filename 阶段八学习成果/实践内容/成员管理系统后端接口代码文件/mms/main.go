// 声明当前文件属于main包
package main 

// 导入其他包或模块
import ( 
	"mms/config" //项目的配置包，提供数据库初始化功能
	"mms/routes" //项目的路由包，提供路由配置功能
)

// main()函数，Go程序的入口
func main() {
	
	//调用config包中的InitDB()函数
	//作用：初始化数据库连接：1.连接到MySQL数据库 2.自动迁移数据库表结构
	config.InitDB() 

	//调用routes包中的SetupRouter()函数，返回给一个r对象
	//作用：初始化路由配置，返回一个 *gin.Engine 实例
	
	//SetupRouter()执行逻辑：
	//1.创建一个gin.Engine实例(Gin框架的核心对象)
	//2.在这个实例上定义所有接口的路由规则
	//3.返回配置好的gin.Engine实例

	//gin.Engine：Gin框架的核心结构体，包含了：
	//1.路由表：存储所有的路由规则(例如POST/register、POST/login等)
	//2.中间件：存储全局和分组的中间件（例如身份验证、日志记录等）
	//3.处理器：存储处理函数，用于响应HTTP请求
	//4.其他功能：如静态文件服务、模板渲染等
	
	//配置好的*gin.Engine是Gin框架的路由引擎实例，包含了所有的路由规则和中间件配置
	//r对象是一个*gin.Engine实例，是一个指针类型，通过它我们可以调用Gin提供的各种方法
	//比如：GET、POST、PUT、DELETE等方法定义路由规则
    //Use方法添加全局中间件，Run方法启动Web服务
	//也就是说，r对象是Gin框架的核心对象，通过它我们可以管理和操作所有的路由和中间件
	r := routes.SetupRouter() 
	
	//启动Web服务器并监听8080端口，等待用户的请求
	//r.Run()表示调用gin框架的Run()方法，启动Web服务
	//":8080"表示监听8080端口，启动后，服务器会监听并处理来自客户端的请求
	//客户端可以通过http://localhost:8080访问服务器
	r.Run(":8080")
}

/*main.go文件作用：
1.初始化数据库连接
2.初始化路由配置
3.启动Web服务*/