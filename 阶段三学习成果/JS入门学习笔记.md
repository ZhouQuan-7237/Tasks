# JS入门学习笔记-周全

## 学习目录

#### 一.JS简介

#####   1. JS基础介绍

#####   2. JS基本用途

#####   3. JS导入方式

#### 二.JS基本语法

#####   1. 变量常量

#####   2. 控制语句

* ##### 条件语句

* ##### 循环语句

#### 三.JS数据类型

* ##### undefined和null

* ##### Number和Boolean

* ##### 字符串

  * ##### 字符串内置方法

* ##### 函数

* ##### 数组

  * ##### 简单介绍

  * ##### 内置方法

* ##### JS对象

* ##### 对象数组

#### 四.JS事件

* ##### 定义

* ##### 绑定方式

#### 五.HTML、CSS和JS关系

##### 1.三者职能

##### 2.三者关系

---

## 学习内容



### 一.JS简介



#### 1.JS基础介绍

**JS（JavaScript）是一种轻量级、解释型、面向对象的脚本语言，同时也是一种编程语言**。



#### 2.JS基本用途

* **客户端脚本：**用于在用户浏览器中执行，实现动态效果和用户交互。

* **网页开发：**与HTML和CSS协同工作，使网页具有更强的交互性和动态性。

* **后端开发：**使用Node.js,JavaScript也可以在服务器端运行，实现服务器端应用的开发。

  

#### 3.JS导入方式

* ##### 内联式

在HTML文档中直接嵌入JS代码，CSS代码是放在style标签内，而JS代码是放在head部分或body部分的script标签内。

![屏幕截图 2024-11-19 181046](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202024-11-19%20181046.png)

![屏幕截图 2024-11-19 181007](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202024-11-19%20181007.png)



```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script>
        console.log('Hello World,body标签的内联样式');//console.log表示要在控制台里面打印一些日志内容
        alert("你好，内联样式弹窗");//alert函数是让网页在刷新时弹出来一个弹窗
    </script>
</head>
<body>
    <script>
        console.log('Hello World,body标签的内联样式');//console.log表示要在控制台里面打印一些日志内容
        alert("你好，内联样式弹窗");//alert函数是让网页在刷新时弹出来一个弹窗
    </script>
</body>
</html>
```

* ##### 外部引入

把JS代码，单独保存外部文件中，通过script标签的src属性引入HTML文档。

![image-20241117101245337](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20241117101245337.png)

![image-20241117101326755](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20241117101326755.png)



### 二.JS基本语法



#### 1.变量常量

* var声明具有函数作用域的变量（变量只在函数内有效）。

* let也声明具有块级作用域的变量（变量只在当前块内有效）。

* const声明常量，值不可更改（但如果const声明的是数组或者对象时，其内部的元素或属性是可以被修改，只是不能将整个数组或对象完全更改）。

let比var更安全，更灵活，避免了var可能出现的一些问题，推荐用let和const。



#### 2.控制语句

* ##### 条件语句

`if`语句的语法如下：

```js
if(condition){
    //如果条件为真，执行这里的代码
}
```

`else`语句的语法如下：

```
if(condition){
    //如果条件为真，执行这里的代码
}else{
	//如果条件为假，执行这里的代码
}
```

`if else`语句的语法如下：

```
if(condition1){
    //如果条件1为真，执行这里的代码
}else if(condition2){
	//如果条件2为假，执行这里的代码
}else{
	//如果以上条件都为假，执行这里的代码
}
```

* ##### 循环语句

`for`循环语句的语法如下：

```js
for(初始化表达式;循环条件;迭代器){
    //循环体，执行这里的代码
}
```

`while`循环语句的语法如下:

```js
while(循环条件){
	//循环体，执行这里的代码
}
```

循环的关键字：break和continue。

`break`用于跳出循环，结束剩余循环的执行。

`continue`用于跳过当前循环中的剩余代码，继续下一次循环。

![屏幕截图 2024-11-19 180813](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202024-11-19%20180813.png)

![屏幕截图 2024-11-17 103150](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202024-11-17%20103150.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>JS基本语法</title>
</head>
<body>
    <script>
        var x;
        let y = 5;
        const PI = 3.14
        console.log(x,y,PI);
        let name = '周全';
        console.log(name);
        let empty_value = null;
        console.log(empty_value);

        let grade = 88;
        if(grade > 60){
            console.log('合格');
        }else{
            console.log('不合格');
        }

        let time = 22;
        if(time < 12){
            alert('上午好');
        }else if(time < 18){
            alert('下午好');
        }else{
            alert('晚上好');
        }

        console.log('for循环');
        for(let i = 1;i <= 10;i++){
            console.log(i);
        }

        console.log('while循环');
        let count = 1;
        while(count <= 10){
            console.log(count);
            count++;
        }

        console.log('循环关键字');
        for(var i = 0;i < 7;i++){
            if(i == 2){
                continue;
            }
            if(i == 5){
                break;
            }
            console.log(i);
        }
    </script>
</body>
</html>
```





### 三.JS数据类型

* #### undefined和null

  * `undefined`是一种 未定义的状态，表示变量声明但未赋值、对象属性不存在、函数无返回值等情况。

  * `null` 是一种明确的空值，表示变量被手动赋值为空，或用于占位表示没有对象值。

  通俗理解：

  - **`undefined`**：还没准备好（可能还会有值）。
  - **`null`**：明确地没有值（我就是空）。

* #### Number和Boolean

|              | **`Number`**                                   | **`Boolean`**                   |
| ------------ | ---------------------------------------------- | ------------------------------- |
| **定义**     | 表示数值类型，包括整数和浮点数。               | 表示逻辑值：`true` 和 `false`。 |
| **特殊值**   | `NaN`、`Infinity`、`-Infinity`。               | `false`、`true`。               |
| **常用方法** | `toFixed()`、`parseInt()`、`parseFloat()` 等。 | 逻辑运算符：`&&`、`||`、`!`。   |
| **应用场景** | 数学运算、格式化数值、数据转换等。             | 条件判断、逻辑控制。            |

![image-20241117104505895](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20241117104505895.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script>
    //String,Number,Boolean,null,undifined
        const username = "John";
        const age = 30;
        const rate = 4.5;
        const isCool = true;
        const x = null;
        const y = undefined;
        let z;
        
        console.log(typeof username);
        console.log(typeof age);
        console.log(typeof rate);
        console.log(typeof isCool);
        console.log(typeof x);
        console.log(typeof y);
        console.log(typeof z);
    </script>
</head>
<body>

</body>
</html>
```

* #### 字符串
  
  * ##### 字符串内置方法

![image-20241117104849532](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20241117104849532.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script>
    //字符串的内置方法
    const s ="Hello World!";

    console.log(s.length);//获取字符串的长度
    console.log(s.toUpperCase());//将字母全部大写
    console.log(s.toLowerCase());//将字母全部小写
    console.log(s.substring(0,5).toUpperCase());//将字符串全部分割开,传入的两个参数分别是起始和终止位置的索引
    console.log(s.split(""));//将字符串转化为数组

    const p = "technology, computers, it, code";
    console.log(p.split(", "));//使用特定分割内容来分割字符串
    </script>
</head>
<body>

</body>
</html>
```

* #### 函数

函数是一段可重复使用的代码块，它接受输入（参数）、执行特定任务，并返回输出

```js
function function_name(参数1，参数2，参数3，...){//参数可以不写，表示不传参
    //函数体，执行这里的代码
    return 返回值;//可选，返回值
}
```

作用域分为全局作用域和局部作用域，

在函数内部声明的变量就有局部作用域，

在函数外部声明的变量具有全局作用域。

![image-20241117105202018](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20241117105202018.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <script>
        function hello(){
            console.log('Hello,world!');
        }

        hello();

        function hello_with_return(){
            return 'Hello,World!-返回值'
        }

        let a = hello_with_return();
        console.log(a);
        console.log(hello_with_return());

        function hello_with_params(name){
            console.log('hello,'+name);
        }
        
        hello_with_params('周全');
        hello_with_params('ZhouQuan');
        
        //作用域
        let global_var = '全局变量';
        function local_var_function(){
            let local_var = '局部变量';
            console.log('函数内打印全局变量:'+ global_var);
            console.log('函数内打印局部变量:'+ global_var);
        }

        local_var_function();

        console.log('全局打印全局变量:'+ global_var);
        console.log('全局打印局部变量:'+ local_var);
    </script>
</body>
</html>
```

* #### 数组

  * ##### 简单介绍

    数组内部可以储存许多元素

    在同一个数组中，每个元素都可以是不同数据类型

    const声明的数组内部的元素是可以被部分修改的

  * ##### 内置方法

  ![image-20241117105848344](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20241117105848344.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script>
    //数组的内置方法
    const numbers = new Array(1,2,3,4,5);
    console.log(numbers);

    const fruits = ["apples","oranges","pears",10,true];
    fruits[5] = "grapes";
    fruits.push("mangos");//数组末尾添加一个元素
    fruits.unshift("strawberries");//数组头部添加一个元素
    fruits.pop();//删除数组末尾的元素
    console.log(Array.isArray(fruits))//判断一个变量是否是数组
    console.log(fruits.indexOf("oranges"));//得到特定元素的索引
    console.log(fruits);
    console.log(fruits[0]);//访问数组中的特定元素
    </script>
</head>
<body>

</body>
</html>
```

* #### JS对象

JavaScript 对象是由属性（键值对）和方法组成的数据结构

属性用来存储数据，方法用来定义行为

表示形式是键值对的组合，键是属性名，值是属性值



![image-20241117110128681](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20241117110128681.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script>
        //对象
    const person = {
        firstName: "John",
        lastName: "Doe",
        age: 30,
        hobbies: ["music","movies","sports"],
        address: {
        street: "50 main st",
        city: "Boston",
        state: "MA", 
        },
    };
    console.log(person);
    console.log(person.firstName,person.lastName);
    console.log(person.hobbies[1]);
    console.log(person.address.city);

    const{ firstName, 
    lastName,
    address:{ city }, 
    } = person;

    console.log(city);


    person.email = "john@gmail.com";

    console.log(person);
    </script>
</head>
<body>

</body>
</html>
```

* #### 对象数组

![image-20241117110632589](https://raw.githubusercontent.com/ZhouQuan-7237/image-bed/main/image-20241117110632589.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script>
        //对象数组
    const todos = [
        {
        id: 1,
        text: "Take out trash",
        isCompleted: true,
        },
        {
        id: 2,
        text: "Meeting with boss",
        isCompleted: true,
        },
        {
        id: 3,
        text: "Dentist appt",
        isCompleted: false,
        },
    ];

    console.log(todos[1].text);
    </script>
</head>
<body>

</body>
</html>
```



### 四.事件



#### **1.定义：**

事件是网页中用户交互的触发点，如点击、悬浮等操作。常见事件如下：

|    事件     |       描述       |
| :---------: | :--------------: |
|   onClick   |     点击事件     |
| onMouseOver |     鼠标经过     |
| onMouseOut  |     鼠标移出     |
|  onChange   | 文本内容改变事件 |
|  onSelect   |    文本框选中    |
|   onFocus   |     光标聚集     |
|   onBlur    |     移开光标     |

#### 2.绑定方式

* 使用`HTML`属性:在HTML标签中添加事件属性

* `DOM`属性

* `addEventListener`方法

  

### 五.HTML、CSS和JS关系



#### 1.三者职能

* ##### HTML：负责页面的结构（内容和框架）

* ##### CSS：负责页面的样式（装饰和布局）

* ##### JavaScript：负责页面的行为（交互和动态效果）

  

#### 2.三者关系

##### HTML 提供结构，CSS 为结构添彩，JavaScript 为页面赋予生命。

三者相辅相成，缺一不可：没有 HTML，页面无内容；没有 CSS，页面无样式；没有 JavaScript，页面无交互。



## 参考资料

1. [前端三剑客 HTML、CSS、JavaScript 入门到上手](https://zhuanlan.zhihu.com/p/526785618)

2. [四十分钟JavaScript快速入门 | 无废话且清晰流畅 | 手敲键盘 | WEB前端必备程序语言~](https://www.bilibili.com/video/BV15L4y1a7or?vd_source=c91cd0d85dd28e19f062edef99115834)

3. [3小时前端入门教程（HTML+CSS+JS）](https://www.bilibili.com/video/BV1BT4y1W7Aw?vd_source=c91cd0d85dd28e19f062edef99115834)
