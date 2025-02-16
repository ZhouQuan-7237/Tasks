# GROM学习笔记——周全

## 学习目录

一.ORM基础知识

1.ORM概念

2.ORM作用

3.ORM原理

二.GORM基础知识

1.GORM概念

2.GORM优势

3.安装配置

4.基础操作

## 学习内容

### 一.ORM基础知识

#### 1.ORM概念

**ORM（Object-Relational Mapping，对象关系映射）**是一种编程技术，用于在面向对象编程语言（如Java、Python、C#等）和关系型数据库（如MySQL、PostgreSQL、Oracle等）之间建立映射关系。它通过将数据库中的表映射为编程语言中的类，表中的行映射为类的实例，表的字段映射为对象的属性，从而使开发者能够用面向对象的方式来操作数据库，而无需直接编写SQL语句。

#### 2.ORM作用

ORM就像是一个“翻译官”，在面向对象的编程语言和数据库之间进行“翻译”，可以让我们用面向对象的方式操作数据库，而无需写复杂的SQL语句。

**3.ORM原理**

![2550000-20220720102530912-199990477](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/2550000-20220720102530912-199990477.png)

```
对象操作 → ORM 翻译 → SQL 执行 → 数据库交互
```

ORM原理：通过定义对象与数据库表的映射关系，自动将面向对象的操作转换为SQL语句并执行，从而实现数据持久化。

### 二.GORM基础知识

#### 1.GORM概念

GORM是一个为Go语言设计的ORM库，通过对象关系映射（ORM）将数据库表与Go结构体绑定，实现以面向对象的方式操作数据库。

#### 2.GORM优势

**1）简化数据库操作**：通过链式API简化CRUD操作。

**2）自动迁移**：自动根据结构体定义同步数据库表结构。

**3）多数据库适配**：兼容多种数据库系统。 

#### 3.安装配置

* (1)打开 Linux 虚拟机

* (2)VSCode远程连接虚拟机

* (3)VSCode远程创建Go项目目录并进入

  * 法一：在文件浏览器中直接新建文件夹，右键“在集成终端中打开”。

  * 法二：在终端中输入命令来创建项目目录并进入。

    * 打开终端：

      ```
      Ctrl + ~
      ```

    * 输入指令：

      ```
      mkdir ~/go/myproject
      cd ~/go/myproject
      ```

* (4)初始化Go模块

  ```
  go mod init myproject
  ```

* (5)安装GORM和MySQL驱动

  ```
  go get gorm.io/gorm
  go get gorm.io/driver/mysql
  ```

* (6)检查 `go.mod` 文件( 整理和检查所有依赖)

  ```
  go mod tidy
  ```

* (7)编写代码

  * 新建文件

    * 法一：直接在文件浏览器创建`main.go`文件

    * 法二：在终端中输入命令来创建文件

      ```
      touch main.go
      ```

  * 编写代码

    ```go
    package main
    
    import (
        "fmt"
        "log"
        "gorm.io/driver/mysql"
        "gorm.io/gorm"
    )
    
    var DB *gorm.DB
    
    // 设置数据库连接
    func SetupDatabase() {
        var err error
        // MySQL的连接字符串，替换为你的用户名、密码和数据库名称，数据库是MySQL上已经存在的
        dsn := "用户名:密码@tcp(127.0.0.1:3306)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"
        DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatal("连接数据库失败:", err)
        }
        fmt.Println("数据库连接成功")
    }
    
    func main() {
        SetupDatabase()
    }
    ```

* (8)运行程序

  ```
  go run main.go
  ```

* (9)验证数据库

  * 虚拟机上检查 MySQL 服务是否正常运行

    ```
    sudo service mysql status
    ```

  * 虚拟机登录到 MySQL数据库

    ```
    su'do mysql -u root -p
    ```

  * 虚拟机查看表是否成功创建，从而检查是否连接到MySQL 的数据库

    ```
    SHOW TABLES;
    ```

  * 如果Workbench已经连接到虚拟机的MySQL数据库，还可以在虚拟机上查看表是否成功创建，从而检查是否成功MySQL 的数据库

    ```
    SHOW TABLES;
    ```

#### 4.基础操作

* **4.1模型定义**

  模型定义是将Go语言结构体与数据库表结构进行映射的过程。

  结构体的字段对应数据库表的列，GORM 提供了标签来指定字段属性和约束。

  * **4.11结构体映射数据库表**

    ```go
    type User struct {
        gorm.Model // 内置模型，包含 ID、CreatedAt、UpdatedAt 和 DeletedAt
        Name       string
        Age        int
        Email      string `gorm:"unique"` // 使用标签定义字段属性（唯一约束）
    }
    ```

  * **4.12字段标签**

    字段标签定义数据库列的特性，比如字段类型、长度、约束等，常见的标签包括：

    - **`gorm:"primaryKey"`**：指定字段为主键。
    - **`gorm:"autoIncrement"`**：指定字段自增长。
    - **`gorm:"unique"`**：指定字段唯一。
    - **`gorm:"not null"`**：指定字段不能为空。
    - **`gorm:"default:value"`**：指定字段的默认值。
    - **`gorm:"index"`**：为字段创建索引。
    - **`gorm:"size:length"`**：指定字段的长度。

    ```go
    type Product struct {
        ID       uint    `gorm:"primaryKey;autoIncrement"`
        Name     string  `gorm:"size:255;not null"`
        Price    float64 `gorm:"not null"`
        Category string  `gorm:"index"`
    }
    ```

  * **4.13`gorm.Model` 内置模型**

    `gorm.Model` 是 GORM 提供的一个结构体，包含了数据库中常见的字段，如：

    **ID**：主键

    **CreatedAt**、**UpdatedAt**、**DeletedAt**：时间戳，支持软删除

    通过嵌入 `gorm.Model`，可以自动为模型添加这些常用字段。

  * **4.14模型字段类型**

    - **普通字段**：直接定义字段类型，如 `string`、`int`、`float64` 等。
    - **可空字段**：使用指针类型（如 `*string`、`*time.Time`）表示可以为空的字段。
    - **零值类型**：使用 `sql.NullType` 处理数据库中可能包含 `NULL` 的字段。
    - **忽略字段**:   使用 `gorm:"-"` 忽略无需映射到数据库表的结构体字段。

    ```
    type User struct {
        ID       uint
        Name     string
        Email    *string      // 可空字段
        Age      uint8        // 无符号整数
        Birthday *time.Time   // 可空时间
        IgnoreMe int `gorm:"-"` // 忽略该字段
    }
    ```

  * **4.15自动迁移**

    GORM支持自动迁移功能，可以根据结构体定义自动生成对应的表格。

    ```go
    db.AutoMigrate(&User{})  
    ```

* **4.2CRUD操作**

  * **4.21创建（Create）**

    * **插入单条记录**

      ```go
      user := User{Name: "Alice", Age: 25}
      db.Create(&user) 
      ```

    * **插入多条记录**

      ```go
      users := []User{
          {Name: "Bob", Age: 20},
          {Name: "Charlie", Age: 30},
      }
      db.Create(&users) 
      ```

  * **4.22读取（Read）**

    * **简单查询**

    ```go
    var user User
    db.First(&user, 1) // 根据主键查询
    db.Take(&user)     // 查询第一条记录
    db.Last(&user)     // 查询最后一条记录
    db.Find(&user, "name = ?", "Alice") // 条件查询
    db.Find(&users) // 查询所有记录
    ```

    * **条件查询**

    ```go
    db.Where("age > ?", 18).Find(&users) // 查询年龄大于 18 的用户
    db.Where("name = ? AND age > ?", "Alice", 20).Find(&users) // 多条件查询
    ```

    * **排序和分页**

    ```go
    db.Order("age ASC").Find(&users) // 按年龄升序排序
    db.Order("age DESC").Find(&users) // 按年龄降序排序
    db.Limit(10).Offset(20).Find(&users) // 分页查询(跳过前20条记录，取接下来的10条记录)
    ```

  * **4.23更新（Update）**

    * **更新所有记录**

      ```go
      var user User
      db.First(&user, 1) // 查询记录
      user.Age = 26      // 修改字段
      db.Save(&user)     // 更新所有字段
      ```

    * **更新单条记录**

      ```go
      db.Model(&user).Update("age", 26) 
      ```

    * **更新多条记录**

      ```go
      db.Model(&user).Updates(User{Name: "Alice", Age: 26})
      ```

  * **4.24删除（Delete）**

    * **删除单条记录**

      ```go
      var user User
      db.First(&user, 1) // 查询记录
      db.Delete(&user)   // 删除记录
      ```

    * **删除多条记录**

      ```go
      db.Where("age > ?", 18).Delete(&User{}) // 删除所有年龄大于 18 的用户
      ```

## 参考资料

请查阅阶段七任务计划思路