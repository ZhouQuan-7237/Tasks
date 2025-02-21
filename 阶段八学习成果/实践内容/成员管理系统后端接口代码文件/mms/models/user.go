// 声明当前文件属于models包
package models

// 定义User结构体
//User结构体用于表示数据库中的用户表，包含了用户的基本信息
//通过定义结构体，可以将数据库表中的字段与Go语言的结构体字段对应起来，方便进行数据库操作
type User struct {
	ID         uint  `gorm:"primaryKey"` //指定ID字段为主键，GORM会自动将ID字段当作数据库表中的主键
	UserID    string `gorm:"unique;not null"` //指定UserID字段唯一且非空
	Password  string `gorm:"not null"` //指定Password字段非空
	Nickname  string 
	Phone     string 
	Email     string `gorm:"unique"` //指定Email字段唯一
	IsAdmin   bool   `gorm:"default:false"` //指定IsAdmin字段默认值为 `false`，即新创建的用户默认不是管理员
	IsApproved bool  `gorm:"default:false"`	//指定IsApproved字段默认值为 `false`，即新创建的用户默认未通过审核
}