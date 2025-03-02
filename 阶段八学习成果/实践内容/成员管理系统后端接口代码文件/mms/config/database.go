// 声明当前文件属于main包
package config

// 导入其他包或模块
import (
	"fmt" //fmt包，用于格式化输入输出，比如打印信息到控制台
	"gorm.io/driver/mysql" //GROM的MySQL驱动包，用于连接MySQL数据库
	"gorm.io/gorm"  //GROM的核心包，提供ORM(对象关系映射)功能，简化Go结构体与数据库表的交互
                    //将​数据库表​映射为程序中的结构体（对象）​，数据库中的一行数据​对应一个结构体实例，表中的​字段​对应结构体的属性
	"mms/models" //项目的模型包，定义数据库模型
)

//定义了一个全局变量DB,它的类型是*gorm.DB(或是一个*gorm.DB实例或是一个指向gorm.DB结构体的指针),也就是GORM的数据库连接对象。通过DB，我们可以操作数据库。
//*gorm.DB是一个指向GORM数据库连接的指针，它表示我们可以对数据库进行增、删、改、查等操作                         
//DB是程序与数据库之间的“快捷操作入口”
//1.记忆数据库连接信息：程序与数据库的连接信息仅需设置一次，之后可通过快速连接操作访问数据库
//2.全局变量共享：数据库连接作为全局变量，程序中的所有函数均可直接访问
//3.操作数据库：通过 DB 对象，使用如 DB.Create()、DB.Find() 等方法，可以操作数据库
var DB *gorm.DB

// 定义InitDB函数，初始化数据库连接
func InitDB() {

    //定义数据库连接字符串dsn（Data Source Name）
    //dsn包含了数据库的连接信息
	//dsn格式：用户名:密码@tcp(主机:端口)/数据库名称?参数
	//用户名和密码用来连接数据库的身份验证
	//tcp(192.168.220.129:3306)包含了数据库的地址和端口，192.168.220.129是虚拟机的IP地址，3306是MySQL的默认端口号
	//mms是数据库的名称，charset=utf8mb4表示指定字符集为utf8mb4，确保支持存储所有Unicode字符，parseTime=True表示启用时间解析，loc=Local设置时区为本地时区
    dsn := "ZhouQuan:@ZhouQuan1134@tcp(192.168.220.129:3306)/mms?charset=utf8mb4&parseTime=True&loc=Local"
    
    //通过GORM连接MySQL数据库
    //调用GROM库的Open方法，使用MySQL数据库驱动打开数据库连接，返回两个对象：db和err
    //gorm.Open() 是GORM库里打开数据库连接的方法，接收两个参数：mysql.Open(dsn)和&gorm.Config{}
    //mysql.Open(dsn)表示使用MySQL数据库驱动，并传入数据库连接字符串dsn
    //&gorm.Config{}表示GORM的配置选项，{}是一个空的配置对象，表示使用默认设置
    //gorm.DB是GORM库中与数据库交互的核心对象
	//db是一个*gorm.DB实例(或类型是*gorm.DB或是一个指向gorm.DB结构体的指针)
    //db代表了与数据库的连接，通过db，我们可以操作数据库
	//error存储可能的错误信息。如果发生错误，错误信息会存储在err变量中;如果没有发生错误，err是nil
	//如果连接成功，db代表数据库连接，err为nil;如果连接失败，db不代表数据库连接，err存储错误信息
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    
    //检查数据库连接是否成功
    //如果err不为nil，说明数据库连接失败
    //panic("数据库连接失败: " + err.Error())：如果连接失败，程序会通过panic函数抛出错误并终止程序执行
    //err.Error()函数打印出具体的错误信息
    if err != nil {
        panic("数据库连接失败: " + err.Error())
    }

    //将成功建立的数据库连接实例db赋值给全局变量DB
    //通过这一步，整个项目可以通过DB来执行数据库操作
    DB = db
    fmt.Println("数据库连接成功")//打印数据库连接成功的信息
    
    //自动迁移User模型
    //调用db.AutoMigrate()方法，对models.User模型进行自动迁移
    //AutoMigrate方法会自动同步数据库表结构，使其与代码中的模型定义一致
	//&models.User{}传递User模型的指针，表示对User模型进行迁移
    db.AutoMigrate(&models.User{})
}