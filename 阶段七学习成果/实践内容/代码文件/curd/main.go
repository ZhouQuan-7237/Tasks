package main				// 声明当前文件属于main包

import (					//导入所需的依赖包
	"fmt" 					//fmt包用于格式化输入输出，比如打印信息到控制台
	"gorm.io/driver/mysql" 	//MySQL数据库驱动包，让GORM能操作MySQL
	"gorm.io/gorm" 			//GORM的核心库,提供ORM(对象关系映射) 的功能，用来简化数据库操作
)

//定义User模型，对应数据库中的users表格
type User struct {
	ID     uint   `gorm:"primaryKey"` //主键列（类似身份证号），自动递增（类似 MySQL 中的 AUTO_INCREMENT）
	Name   string `gorm:"not null"`   //姓名列，不能为空
	Gender string `gorm:"not null"`   //性别列，不能为空
	MBTI   string `gorm:"size:4"`     //MBTI列，最多4个字符（如INTP）
}

//设置数据库连接
func SetupDatabase() (*gorm.DB, error) {
	
	// 数据库连接信息
	//dsn是数据库连接字符串（Data Source Name），包含了连接数据库所需的所有信息：
	//dsn格式：用户名:密码@tcp(主机:端口)/数据库名称?参数
	//用户名和密码用来连接数据库的身份验证
	//tcp(192.168.220.129:3306)：192.168.220.129是虚拟机的IP地址，3306是MySQL默认的端口号，意思是我们要连接虚拟机的MySQL数据库。
	//curd是数据库的名字。后面有一些其他的配置项：charset=utf8mb4表示字符集，确保支持 Unicode，parseTime=True表示解析时间，loc=Local设置时区为本地时区
	dsn := "ZhouQuan:@ZhouQuan1134@tcp(192.168.220.129:3306)/curd?charset=utf8mb4&parseTime=True&loc=Local"
	
	// 通过GORM连接到MySQL数据库
	//db:是指向gorm.DB结构体的指针，gorm.DB是GORM库中与数据库交互的核心对象，通过db，我们可以操作数据库
	//error：存储可能错误信息的变量。如果发生错误，err将被赋值为错误信息。如果没有发生错误，err会是nil
	//gorm.Open 是GORM库的一个函数，来打开与数据库的连接。它接收两个参数：1.mysql.Open(dsn)：用 MySQL 驱动来打开数据库，dsn 是上面定义的连接字符串。2.&gorm.Config{}：这是 GORM 的配置选项，{} 是一个空的配置对象，表示使用默认设置。
	//如果连接成功，db会代表数据库连接，err为nil;如果连接失败，db不代表数据库连接，err存储错误信息
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//如果err不为nil，代表连接失败，则db不存储数据库连接，err存储错误信息，并返回db和err
	if err != nil {
		return nil, err
	}

	// 自动迁移：根据User结构体来自动创建或更新数据库表
	//db.AutoMigrate(&User{}):利用GORM的AutoMigrate方法，根据User结构体的定义来自动创建或更新数据库中的表结构
	//err:AutoMigrate方法的返回值是一个错误对象err，如果迁移过程中出现任何问题（例如数据库连接失败、权限问题等），err将会捕获该错误
	err = db.AutoMigrate(&User{})
	//如果自动迁移失败，函数会返回nil和错误信息
	if err != nil {
		return nil, err
	}
	//如果自动迁移成功，函数会返回数据库连接对象db和nil错误
	return db, nil
}

// 创建用户（Create）
func createUser(db *gorm.DB, name string, gender string, mbti string) {
	
	//创建一个新的User实例，并赋予user变量
	user := User{
		Name:   name,    //将赋值的name参数赋值给User结构体中的Name字段，即将函数参数name保存到user实例的Name字段
		Gender: gender,  //将赋值的gender参数赋值给User结构体中的Gender字段，即将函数参数gender保存到user实例的Gender字段
		MBTI:   mbti,    //将赋值的mbti参数赋值给User结构体中的MBTI字段，即将函数参数mbti保存到user实例的MBTI字段
	}
	
	//使用GORM的Create方法将用户数据插入到数据库
	//db.Create()是GORM提供将数据插入到数据库中的对应表的方法
	//&user是user的指针，因为GORM需要修改传入的user对象，插入完成后它会将数据库中生成的ID等信息填充回user中
	//result存储Create方法的返回值，result 是一个包含操作结果的结构体，里面包含了Error字段，用来记录插入过程中发生的错误信息
	result := db.Create(&user)
	//如果插入失败，打印错误信息
	if result.Error != nil {
		fmt.Println("创建失败:", result.Error)
	} else {
	//如果插入成功，打印创建的用户信息
		fmt.Println("用户已创建:", user)
	}
}

// 查询用户（Read）
func getUser(db *gorm.DB, userID uint) {
	
	//声明user变量，它是一个空的User结构体类型的实例，用来储查询结果
	//User结构体包含了数据库中与用户相关的字段（如 ID、姓名、性别、MBTI）
	//user用来接收查询结果，最终将存储查询到的用户信息
	var user User 
	
	//使用GORM的First方法，根据用户ID查找用户
	//db.First()是GORM提供的根据主键查询表中的数据的方法
	//&user：将user变量的地址传给db.First方法。GORM会将查询结果填充到user变量中
	//userID：作为查询条件，表示根据指定的用户ID查询对应的用户
	//result存储First方法的返回值，result 是一个包含操作结果的结构体，里面包含了Error字段，用来记录查询过程中发生的错误信息
	result := db.First(&user, userID)
	//如果查询失败，打印错误信息
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
	} else {
	// 如果查询成功，打印用户的信息
		fmt.Printf("查询结果: ID=%d, 姓名=%s, 性别=%s, MBTI=%s\n",
			user.ID, user.Name, user.Gender, user.MBTI)
	}
}

// 更新用户（Update）
func updateUser(db *gorm.DB, userID uint, newName string, newGender string, newMBTI string) {
	
	//声明user变量，它是一个空的User结构体类型的实例，用来储查询结果
	//User结构体包含了数据库中与用户相关的字段（如 ID、姓名、性别、MBTI）
	//user用来接收查询结果，最终将存储查询到的用户信息
	var user User // 创建一个空的 User 实例，用于存储查询结果
	
	//使用GORM的First方法，根据用户ID查找用户
	//db.First()是GORM提供的根据主键查询表中的数据的方法
	//&user：将user变量的地址传给db.First方法。GORM会将查询结果填充到user变量中
	//userID：作为查询条件，表示根据指定的用户ID查询对应的用户
	//result存储First方法的返回值，result 是一个包含操作结果的结构体，里面包含了Error字段，用来记录查询过程中发生的错误信息
	result := db.First(&user, userID)
	//如果查询用户失败，打印错误信息，并通过return语句提前结束程序
	if result.Error != nil {
		fmt.Println("用户不存在:", result.Error)
		return
	}

	//如果用户存在，更新用户的信息
	user.Name = newName     //将新的姓名newName赋值给user变量中的Name字段，更新用户的姓名
	user.Gender = newGender //将新的性别newGender赋值给user变量中的Gender字段，更新用户的性别
	user.MBTI = newMBTI     //将新的MBTI类型newMBTI赋值给user变量中的MBTI字段，更新用户的MBTI类型
	
	//使用GORM的Save方法将更新后的用户信息保存到数据库
	//db.Save（）是 GORM 提供的将修改后的user实例保存到数据库的一个方法
	//Save会检查该对象是否有主键值（即是否存在），如果该对象已经存在（根据 userID 查找到了该用户），则会更新数据库中的现有记录
	//如果该记录不存在（即 userID 不对应任何用户），则会插入新记录
	db.Save(&user)
	// 打印更新后的用户信息
	fmt.Println("用户已更新:", user)
}

// 删除用户（Delete）
func deleteUser(db *gorm.DB, userID uint) {
	
	//使用GORM的Delete方法根据用户ID删除用户
	//db.Delete()是GORM提供的从数据库中删除记录的方法
	//&User{}表示传入一个空的User结构体指针。Delete方法会根据结构体的定义删除数据库中对应表的记录，这里是User表
	//userID表示需要删除的用户的ID，Delete方法会根据该ID查找并删除对应的用户
	//db.Delete方法执行后，会返回一个result对象，该对象包含了操作的结果，包括是否成功删除和可能的错误信息
	result := db.Delete(&User{}, userID)
	
	//如果删除失败，打印错误信息
	if result.Error != nil {
		fmt.Println("删除失败:", result.Error)
	} else {
	//如果删除成功，打印被删除用户的ID
		fmt.Println("用户已删除,ID:", userID)
	}
}

// main函数，程序的入口
func main() {

	//调用前面定义的SetupDatabase函数来初始化数据库连接
	//SetupDatabase返回两个值：
	//db：数据库连接对象，用于后续的数据库操作
	//err：错误对象，如果数据库连接失败，err会存储错误信息
	db, err := SetupDatabase()
	// 如果数据库连接失败，打印错误信息并退出程序
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	//创建用户（姓名：小明，性别：男，MBTI：INFJ）
	//调用前面定义的createUser函数，传入数据库连接对象db和用户的基本信息
	createUser(db, "小明", "男", "INFJ")
	//创建另一个用户（姓名：小红，性别：女，MBTI：ENFP）
	createUser(db, "小红", "女", "ENFP")

	//查询用户（ID为1的用户）
	//调用前面定义的getUser函数，传入数据库连接对象db和用户ID
	getUser(db, 1)

	//更新用户（ID为2的用户，更新姓名为“小红”，性别为“女”，MBTI为“INFP”）
	//调用前面定义的updateUser函数，传入数据库连接对象db、更新用户的ID和基本信息
	updateUser(db, 2, "小红", "女", "INFP")

	//删除用户（ID为2的用户）
	//调用前面定义的deleteUser函数，传入数据库连接对象db和用户ID
	deleteUser(db, 2)
}
