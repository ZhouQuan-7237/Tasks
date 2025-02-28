# MySQL入门学习笔记——周全

## 学习目录

一.MySQL简介

1. MySQL基础介绍
2. MySQL前世今生
3. MySQL remote连接

二.MySQL学习

1. 基础知识
2. 表格建联
3. CURD语句

## 学习内容

### 一.MySQL简介

#### 1.MySQL基础介绍

MySQL是一款开源的关系型数据库管理系统（RDBMS），使用SQL语言，以高性能、可靠性和可扩展性着称，采用典型的客户端—服务器架构(C/S架构)。

#### ![image-20250216215331354](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250216215331354.png)

![2d65cfb6aafc8aa48d6c3b0e0dcac2fc](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/2d65cfb6aafc8aa48d6c3b0e0dcac2fc.png)

#### 2.MySQL前世今生

![4f7b659bcf0734c377e654edd25948b4d41a4b7f](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/4f7b659bcf0734c377e654edd25948b4d41a4b7f.jpg)

* **MySQL的早期发展（1995年-2000年）**

  * **1995年**：MySQL AB在瑞典成立，发布仅供内部使用的MySQL 1.0版本。

  * **1996年**：MySQL发布3.11.1版本，跳过2.x版本，开始面向公众开放。

  * **1999年**：发布3.23版本，集成Berkeley DB存储引擎，为后续存储引擎架构奠定基础。

* **MySQL的发展与Sun的收购（2000年-2008年）**

  * **2000年**：ISAM升级为MyISAM，MySQL开始使用GPL协议开源。

  * **2001年**：MySQL 4.0发布，集成InnoDB存储引擎，支持事务和行级锁。

  * **2005年**：MySQL 5.0发布，支持游标、存储过程、触发器、视图和XA事务；Oracle收购InnoDB开发商Innobase。

  * **2008年**：Sun以10亿美元收购MySQL AB，发布MySQL 5.1，增加定时器、分区和基于行的复制等特性。

* **MySQL的创新与Oracle收购（2009年-2017年）**

  * **2009年**：Oracle收购Sun公司，间接控制MySQL。

  * **2010年**：MySQL 5.5发布，InnoDB成为默认存储引擎，增加多核扩展、半同步复制和utf8mb4字符集等特性。

  * **2013年**：MySQL 5.6发布，增加GTID复制、丢失复制和延迟复制等功能，成为里程碑版本。

  * **2015年**：MySQL 5.7发布，增加组复制、InnoDB Cluster和多源复制等功能，优化性能和安全性。

* **MySQL的现代化与创新（2018年至今）**

  * **2018年**：MySQL 8.0发布，基于InnoDB的数据字典，支持原子DDL、降序索引，默认字符集改为utf8mb4，进一步优化性能和安全性。

  * **2024年4月30日**：MySQL 8.4发布，成为8.x系列的长期支持版本。

#### 3.MySQL remote连接

![image-20250210135355234](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20250210135355234.png)

### 二.MySQL学习

#### 1.基础知识

##### 1.1相关概念

![07FAD8111A192A47BAF43A9B31C7A45C - 副本](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/07FAD8111A192A47BAF43A9B31C7A45C%20-%20%E5%89%AF%E6%9C%AC.jpg)

```
MySQL:相当于EXCEL
数据库:相当于EXCEL文件
表格：相当于EXCEL文件里的标签
```

```
DB：数据库，保存了一系列有组织的数据；
DBMS：数据库管理系统；
SQL：结构化查询语言，用来与数据库进行沟通；
```

##### 1.2语法规范

```
1. 不区分大小写，建议关键字大写，表名，列名小写
2. 每条命令最好分号结尾；
3. 每条命令根据需要进行缩进，换行；
4. 注释
   1,单行注释：# 注释内容；
   2,单行注释：-- 注释内容（-- 后面必须有空格）
   3,多行注释：/* 注释内容 */
```

##### 1.3SQL语言

```
1. DQL：数据查询语言；select
2. DML：数据操作语言 增删改
3. DDL：数据定义语言
4. TCL：事务控制语言
```

##### 1.4字段类型

* **数值类型** ：

  - `INT`：存储整数。

  - `FLOAT`：存储浮点数。

  - `DECIMAL`：存储精确的小数，可指定精度和小数位数，

    如 `DECIMAL(10, 2)` 表示总共 10 位数字，其中 2 位是小数。

- **字符串类型** ：

  - `CHAR`：固定长度的字符串，

    比如 `CHAR(10)` 会占用 10 个字符的空间，不足部分会用空格填充。

  - `VARCHAR`：可变长度的字符串，比如 `VARCHAR(50)`。

- **日期和时间类型** ：

  - `DATE`：存储日期，格式为 `YYYY-MM-DD`。
  - `TIME`：存储时间，格式为 `HH:MM:SS`。
  - `DATETIME`：存储日期和时间的组合，格式为 `YYYY-MM-DD HH:MM:SS`。

##### 1.5常见约束

* **非空约束**：`NOT NULL`，确保字段值不能为空。

* **默认值约束**：`DEFAULT`，为字段设置默认值。

* **主键约束**：`PRIMARY KEY`，确保字段值唯一且非空。

* **唯一约束**：`UNIQUE`，确保字段值唯一，但可以为空。

* **外键约束**：`FOREIGN KEY`，用于在一个表中建立与另一个表的关联，确保字段值来自关联表的主键或唯一键，以保证数据的一致性。

#### 主键和唯一键的区别：

* **主键**：用于唯一标识表中每行记录，像身份证号，不允NULL值，一个表仅一个主键。

* **唯一**：确保列值唯一，可含NULL值，一个表可设多个唯一约束。

##### 1.6数据表结构

数据表由行（记录）和列（字段）组成，每行存储一条记录，每列代表一个数据字段，数据字段定义了数据的类型和约束。

表格举例：

```
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,  -- 主键，自动递增
    username VARCHAR(50) NOT NULL,      -- 非空字段
    email VARCHAR(100) UNIQUE,          -- 唯一约束
    age INT DEFAULT 18                  -- 默认值
);
```

##### 1.7基础操作

* 创建数据库:

```sql
CREATE DATABASE 数据库名;
```

* 切换数据库:

```sql
USE 数据库名;
```

* 创建表格:

```sql
CREATE TABLE 表格名 (
    列名1 类型(长度) 约束,
    列名2 类型(长度) 约束,
    列名3 类型(长度) 约束
);
```

* 删除数据库：

```sql
DROP DATABASE 数据库名;
```

* 删除表格：

```sql
DROP TABLE 数据库名.表格名;
```

* 查看所有数据库：

```sql
SHOW DATABASES;
```

* 查看所有表格：

```sql
SHOW TABLES;
```

* 查看表结构：

```sql
DESC 表格名;
```

* 重命名表格

```sql
RENAME TABLE 旧表格名 TO 新表格名;
```

* 重置 `users` 表的 `AUTO_INCREMENT` 值

将 `AUTO_INCREMENT` 的起始值设置为 1，那么下一个插入的用户 ID 会从 1 开始

```sql
ALTER TABLE users AUTO_INCREMENT = 1;
```

#### 2.表格建联

**2.1一对一关联**：在从表中添加一个外键，引用主表的主键或唯一键。

eg.每个员工对应一个员工详细信息记录。

```sql
CREATE TABLE employee (
    employee_id INT PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE employee_detail (
    detail_id INT PRIMARY KEY,
    employee_id INT,
    address VARCHAR(100),
    FOREIGN KEY (employee_id) REFERENCES employee(employee_id)
);
```

`employee_detail` 表通过 `employee_id` 外键引用 `employee` 表的主键，从而建立关联，

确保每个员工对应一个员工详细信息记录。

**2.2一对多关联**：在“多”的表中添加一个外键，引用“一”的表的主键或唯一键。

eg.一个部门可以有多个员工，但一个员工只能属于一个部门。

```sql
CREATE TABLE department (
    department_id INT PRIMARY KEY,
    department_name VARCHAR(50)
);

CREATE TABLE employee (
    employee_id INT PRIMARY KEY,
    name VARCHAR(50),
    department_id INT,
    FOREIGN KEY (department_id) REFERENCES department(department_id)
);
```

`employee` 表通过 `department_id` 外键引用 `department` 表的主键，从而建立关联

实现一个部门对应多个员工的关系。

**2.3多对多关联**：通过创建一个中间表来实现，中间表包含两个表的外键。

eg.一个学生可以选修多门课程，一门课程也可以被多个学生选修。

```sql
CREATE TABLE student (
    student_id INT PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE course (
    course_id INT PRIMARY KEY,
    course_name VARCHAR(50)
);

CREATE TABLE student_course (
    student_id INT,
    course_id INT,
    FOREIGN KEY (student_id) REFERENCES student(student_id),
    FOREIGN KEY (course_id) REFERENCES course(course_id),
    PRIMARY KEY (student_id, course_id)
);
```

`student_course` 中间表通过 `student_id` 和 `course_id` 外键分别引用 `student` 和 `course` 表的主键，

实现学生和课程的多对多关系。

#### 3.CURD语句

* **Create（创建）**: 向表中添加新数据（`INSERT`）

  * 插入一条记录:

  ```sql
  INSERT INTO 数据库名.表格名 (列名1, 列名2, 列名3) VALUES 
  (数值1, 数值2, 数值3);
  ```

  * 插入多条记录:

  ```sql
  INSERT INTO 数据库名.表格名 (列名1, 列名2, 列名3) VALUES 
  (数值1, 数值2, 数值3),
  (数值4, 数值5, 数值6);
  ```

* **Read（查询）**: 从表中读取数据（`SELECT`）

  * 查询所有数据

  ```sql
  SELECT * FROM 表格名;
  ```

  * 查询特定列

  ```sql
  SELECT 列名1, 列名2 
  FROM 表格名;
  ```

  * 去重查询

  ```sql
  SELECT DISTINCT 列名 
  FROM 表格名;
  ```

  * 排序查询（升序）

  ```sql
  SELECT * FROM 表格名 
  ORDER BY 列名 ASC;
  ```

  * 排序查询（降序）

  ```sql
  SELECT * FROM 表格名
  ORDER BY 列名 DESC;
  ```

  * 条件查询

  ```sql
  SELECT * FROM 表格名 WHERE 含有比较运算符或逻辑运算符的条件;
  ```

  ```sql
  SELECT * FROM 表格名 WHERE 列名1 = 数值1 AND 列名2 > 数值2;
  ```

  ![EE5DE6A7FC87D5EB71C795EF18F254D9 - 副本](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/EE5DE6A7FC87D5EB71C795EF18F254D9%20-%20%E5%89%AF%E6%9C%AC.jpg)

  * 范围查询

    * 数值范围

    ```sql
    SELECT * FROM 表格名 WHERE 列名 BETWEEN 数值1 AND 数值2;
    ```

    * 字符范围

    ```sql
    SELECT * FROM 表格名 WHERE 列名 IN ('字符1', '字符2');
    ```

    * 首字母模糊查询

    ```sql
    SELECT * FROM 表格名 WHERE 列名 LIKE '首字母%';
    ```

    * 尾字母模糊查询

    ```sql
    SELECT * FROM 表格名 WHERE 列名 LIKE '%尾字母';
    ```

  * 表连接（取交集）

    ```SQL
    SELECT * FROM 表格1名 INNER JOIN 表格2名 ON 表格1名.列名 = 表格2名.列名; 
    ```

    ```sql
    SELECT * FROM table1 INNER JOIN table2 ON table1.id = table2.id; 
    ```

    从 table1 和 table2 两张表中，找出那些在某个列上匹配的记录，并把它们结合在一起。

    table1在左，table2在右。

    | id   | name  |
    | ---- | ----- |
    | 1    | Alice |
    | 2    | Bob   |

    | id   | sorce |
    | ---- | ----- |
    | 1    | 95    |
    | 2    | 90    |

    | id   | name  | id   | score |
    | ---- | ----- | ---- | ----- |
    | 1    | Alice | 1    | 95    |
    | 2    | Bob   | 2    | 90    |

  * 表连接（左连接）

    ```SQL
    SELECT * FROM 表格1名 LEFT JOIN 表格2名 ON 表格1名.列名 = 表格2名.列名; 
    ```

    ```sql
    SELECT * FROM table1 LEFT JOIN table2 ON table1.id = table2.id; 
    ```

    从 table1 和 table2 两张表中，保留左边表格所有选取的记录，再和右边表格符合条件的记录结合在一起。

    table1在左，table2在右。

    | id   | name  |
    | ---- | ----- |
    | 1    | Alice |
    | 2    | Bob   |
    | 3    | John  |

    | id   | sorce |
    | ---- | ----- |
    | 1    | 95    |
    | 2    | 90    |

    | id   | name  | id   | score |
    | ---- | ----- | ---- | ----- |
    | 1    | Alice | 1    | 95    |
    | 2    | Bob   | 2    | 90    |
    | 3    | Jon   | null | null  |

  * 表连接（右连接）

    ```sql
    SELECT * FROM 表格1名 RIGHT JOIN 表格2名 ON 表格1名.列名 = 表格2名.列名; 
    ```

    ```sql
    SELECT * FROM table1 RIGHT JOIN table2 ON table1.id = table2.id; 
    ```

    从 table1 和 table2 两张表中，保留右边表格所有选取的记录，再和左边表格符合条件的记录结合在一起。

    table1在左，table2在右。

    | id   | name  |
    | ---- | ----- |
    | 1    | Alice |
    | 2    | Bob   |
    | 4    | John  |

    | id   | sorce |
    | ---- | ----- |
    | 1    | 95    |
    | 2    | 90    |
    | 3    | 80    |

    | id   | name  | id   | score |
    | ---- | ----- | ---- | ----- |
    | 1    | Alice | 1    | 95    |
    | 2    | Bob   | 2    | 90    |
    | null | null  | 3    | 80    |

  * 表合并（并集，不包含重复值）

    ```sql
    SELECT 列名 FROM 表格1名 
    UNION 
    SELECT 列名 FROM 表格2名;
    ```

  * 表合并（并集，不包含重复值）

    ```sql
    SELECT 列名 FROM 表格1名 
    UNION ALL
    SELECT 列名 FROM 表格2名;
    ```

* **Update（更新）**: 修改表中的现有数据（`UPDATE`）

  ```sql
  UPDATE 数据库名.表格名
  SET 值
  WHERE 条件;
  ```

  ```sql
  UPDATE students
  SET age = 18
  WHERE id = 1;
  ```

* **Delete（删除）**: 从表中删除数据（`DELETE`）

  ```sql
  DELETE FROM 数据库名.表格名
  WHERE 条件;
  ```

  ```sql
  DELETE FROM students
  WHERE id=1;
  ```

## 参考资料

请查阅阶段七任务计划思路