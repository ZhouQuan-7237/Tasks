// 声明当前文件属于utils包
package utils

// 导入其他包或模块
import "golang.org/x/crypto/bcrypt" //bcrypt 包，用于密码加密和验证

// HashPassword函数使用bcrypt算法加密密码
//func：关键字，用于定义函数
//HashPassword：函数名，表示该函数用于加密密码
//password string：函数参数，表示要加密的密码，类型为string
//(string, error)：函数返回值类型为string和error
func HashPassword(password string) (string, error) {
	
	// 使用bcrypt加密密码，生成加密后的密码字节切片
	//bcrypt.GenerateFromPassword：bcrypt包提供的函数，用于生成加密后的密码
	//[]byte(password)：将字符串password转换为字节切片
	//bcrypt.DefaultCost：表示bcrypt的默认加密难度
	//bytes：返回的加密后的密码，类型为字节切片
	//err：返回的错误信息，如果加密失败，err会包含具体的错误信息，如果加密成功，error为nil
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	// 返回加密后的密码字符串和可能的错误信息
	//string(bytes)：将加密后的密码字节切片转换为字符串
	return string(bytes), err
}

// CheckPassword函数验证密码是否匹配
//func：关键字，用于定义函数
//CheckPassword：函数名，表示该函数用于验证密码是否匹配
//password：函数参数，表示用户输入的密码，类型为string
//hash：函数参数，表示数据库中存储的加密密码，类型为string
//bool：函数返回值类型为bool，true：密码匹配，false：密码不匹配
func CheckPassword(password, hash string) bool {
	//bcrypt.CompareHashAndPassword：bcrypt包提供的方法，用于比较加密密码和原始密码
	//[]byte(hash)：将加密密码hash转换为字节切片
	//[]byte(password)：将用户输入的密码password转换为字节切片
	//err：返回的错误信息，如果密码不匹配，err会包含具体的错误信息，如果密码匹配，err为nil
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	
	//返回一个布尔值，表示密码是否匹配
	//如果err为nil，表示密码匹配，返回true
	//如果err不为nil，表示密码不匹配，返回false
	return err == nil
}
