# HTML入门学习笔记-周全

## 学习目录

### 一.HTML简介

####   1. HTML基础介绍

####   2. HTML解析原理

####   3. HTML文件结构

### 二.HTML常用标签

####   1. 文本标签

* ##### 标题标签（`<h1>` 到 `<h6>`）

* ##### 段落标签（`<p>`）

* ##### 换行标签（`<br>`）

* ##### 注释标签（`<!-- ... -->`）

* ##### 水平线标签（`<hr>`）

* ##### 文本格式化标签

  * ###### 加粗（`<b>` 或 `<strong>`）

  * ###### 斜体（`<i>` 或 `<em>`）
  * ###### 下划线（`<u>`）

####   2. 列表标签

* ##### 有序列表标签（`<ol>` 和 `<li>`）
* ##### 无序列表标签（`<ul>` 和 `<li>`）

####   3. 表格标签

* ##### 表格容器（`<table>`）
* ##### 表格标题（`<caption>`）

* ##### 表格行（`<tr>`）

* ##### 表格数据单元（`<td>` 和 `<th>`）

####   4. 图片标签（`<img>`）

####   5. 链接标签

* ##### 跳转链接标签（`<a href="URL">`）
* ##### 锚点链接标签（`<a href="#id">`）

####   6. 表单标签

* ##### 输入标签（`<input>` 和其多种类型）
* ##### 按钮标签（`<button>`）
* ##### 单选框标签（`<input type="radio">`）
* ##### 多选框标签（`<input type="checkbox">`）
* ##### 提交标签（`<input type="submit">`）

####   7. 单标签

####   8. 无语义容器标签

* ##### `<div>`
* ##### `<span>`

### 三.行内元素与块级元素

####  1. 比较

####  2. 转换

####  3. 行内块元素（补充）

### 四.HTML标签属性

#### 1.属性介绍

#### 2.基本语法

#### 3.基本规则

#### 4.通用属性

---

## 学习内容



#### 一.HTML简介



####  1. HTML基础介绍

* **定义：HTML（超文本标记语言）是一种用于构建网页内容的标记语言。**

  它不像编程语言那样包含逻辑操作，而是通过标签来定义网页中的元素。

* **作用：HTML为页面提供基本结构。**

  它告诉浏览器哪些内容是标题、段落、列表、图片等，使网页内容能按预期显示。

* **工作方式：**HTML通过标签（也称为元素）来标记内容的类型，并通过属性为标签提供更多的信息，从而帮助浏览器正确解析和呈现网页。

  

#### 2. HTML解析原理

* 浏览器可以看作一个函数，HTML文件是输入的数据。浏览器函数接收并解析HTML文件数据，将其转化为网页画面。

* 在解析过程中，浏览器会从头到尾读取 HTML 文件，找到其中的标签，并将每个标签看作一个“盒子”，这些盒子可以包含文本、图片或其他盒子等。然后，浏览器为每个盒子添加样式信息，来确定盒子的大小、颜色和位置。最后，浏览器把这些盒子像俄罗斯方块一样组合搭建就形成了最后我们能看到的网页画面。

  

####  3. HTML文件结构

```html
<!DOCTYPE html><!--声明文档类型，告诉浏览器这是一个HTML5文件-->
<html lang="en"><!--是HTML文档的最外层标签,包含整个文档的结构，包括“后台信息”（放在 <head> 标签中）和“显示内容”（放在 <body> 标签中）。lang="en"说明文档的默认语言是英文-->
<head><!--存放文档的“后台信息”,不直接展示在网页上，但会影响浏览器和搜索引擎对网页的理解-->
    <meta charset="UTF-8"><!--设置网页的字符编码为UTF-8,避免乱码-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0"><!--让网页宽度自动适应设备屏幕,优化移动端显示效果-->
    <title>Document</title><!--定义网页标题，显示在浏览器标签栏-->
</head>
<body><!--存放页面的“显示信息”,包含所有在网页中显示的内容。此处为空，说明网页上没有内容-->
    
</body>
</html>
```



### 二.HTML常用标签



####  1. 文本标签

* ##### 标题标签（`<h1>` 到 `<h6>`）

  ![image-20241116130803492](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116130803492.png)

```html
<h1>一级标题</h1>
<h2>二级标题</h2>
<h3>三级标题</h3>
<h4>四级标题</h4>
<h5>五级标题</h5>
<h6>六级标题</h6>
```

* ##### 段落标签（`<p>`）

  ![image-20241116191305069](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116191305069.png)

```html
<p>这是一个段落</p>
```

* ##### 换行标签（`<br>`）

  ![image-20241116191322588](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116191322588.png)

```html
<p>这是一个段落</p><br><p>这是另一个段落</p>
```

* ##### 注释标签（`<!-- ... -->`）

  ![image-20241116191440289](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116191440289.png)

```html
<!-- 注释 -->
```

* ##### 分隔线标签（`<hr>`）

  ![image-20241116191340501](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116191340501.png)

```html
<p>这是一个段落</p>
<hr>
<p>这是另一个段落</p>
```

* ##### 文本格式化标签

  ![image-20241116191359514](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116191359514.png)

  * ###### 加粗（`<b>` 或 `<strong>`）

  ```html
  <b>加粗</b>
  ```

  * ###### 斜体（`<i>` 或 `<em>`）

  ```html
  <i>斜体</i>
  ```

  * ###### 下划线（`<u>`）

  ```html
  <u>下划线</u>
  ```

  * ###### 删除线（`<s>`）

  ```html
  <s>删除线</s>
  ```

####  2. 列表标签

* ##### 有序列表标签（`<ol>` 和 `<li>`）

  ![image-20241116130655645](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116130655645.png)

``` html
<ol>
    <li>有序列表元素1</li>
    <li>有序列表元素2</li>
    <li>有序列表元素3</li>
</ol>
```

* ##### 无序列表标签（`<ul>` 和 `<li>`）

  ![image-20241116130718522](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116130718522.png)

```html
<ul>
    <li>无序列表元素1</li>
    <li>无序列表元素2</li>
    <li>无序列表元素3</li>
</ul>
```

####  3. 表格标签（table）

* ##### 列标题（`<th>`）

* ##### 表格行（`<tr>`）

* ##### 表格数据（`<td>`）

  ![image-20241116130604218](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116130604218.png)

```html
<table border="1"><!-- 表格边框粗细为1像素 -->
	<tr>
		<th>列标题1</th>
         <th>列标题2</th>
         <th>列标题3</th>
     </tr>
     <tr>
         <td>元素1</td>
         <td>元素2</td>
         <td>元素3</td>
     </tr>
     <tr>
         <td>元素4</td>
         <td>元素5</td>
         <td>元素6</td>
     </tr>
</table>
```

#### 4. 图片标签（`<img>`）

![碎梦渐全](D:\QQ头像\碎梦渐全.jpg)

```html
<img src="./碎梦渐全.jpg" 
     alt="QQ头像"
     width="100" 
     height="100">
<!--src属性定义了图像文件的路径(相对路径或绝对路径）orURL,
src="./碎梦渐全.jpg" 表示图片在当前 HTML 文件所在目录的相对路径下
-alt属性定义了图像的替代文本,当图像无法加载时,浏览器就会显示出替代文本-->
```

####  5. 链接标签（`<a href="URL">`）

* ##### 不保留原页面

```html
 <a href="http://baidu.com">百度</a>
 <!-- 不保留原页面：在原页面中跳转到指定页面或网站
   herf属性定义了链接地址 --> 
```

* ##### 保留原页面

```html
<a href="http://baidu.com" target="_blank">百度</a>
<!-- 保留原页面：在新标签页中跳转到指定页面或网站
  herf属性定义了链接地址，
  target属性定义了链接打开的方式，_blank 表示在新标签页中打开页面 -->
```

####  6. 表单标签（form）

* ##### 输入标签（`<input>` )

* ##### 按钮标签（`<button>`）
* ##### 单选框标签（`<input type="radio">`）
* ##### 多选框标签（`<input type="checkbox">`）
* ##### 提交标签（`<input type="submit">`）

![image-20241116170352074](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116170352074.png)

```html
<form><!--表单的容器-->
    <label>用户名：</label>
    <input type="text" placeholder="请输入用户名"><!--type属性规定了input元素的类型，placeholder属性的属性值内容可作为用户填写时的提示-->
    <br>
    <br>
    <label>密码：</label>
    <input type="text" value="请输入密码"><!--value属性的属性值内容可作为替用户提前输入的一些内容-->
    <br>
    <br>
    <button>确认</button>
    <br>
    <br>
    <label for="username">用户名：</label><!--for属性一般是与input标签中的id所对应的，id对于每个标签来说都是唯一的-->
    <input type="text" id="username" placeholder="请输入用户名">
    <br>
    <br>
    <label for="pwd">密码：</label>
    <input type="password" id="pwd" placeholder="请输入密码"><!--type的属性值password可以实现输入密码时隐藏密码-->
    <br>
    <br>       
    <label>性别</label>
    <input type="radio" name="gender">男<!--设置单选框时可以让type的属性值为“radio”,要让性别只能选其一，需为同组的单选框设置相同的 name 值-->
    <input type="radio" name="gender">女<!--设置单选框时可以让type的属性值为“radio”,要让性别只能选其一，需为同组的单选框设置相同的 name 值-->
    <br>
    <br>
    <label>爱好：</label>
    <input type="checkbox" name="hobby">唱歌<!--设置多选框时可以让type的属性值为“checkbox”-->
    <input type="checkbox" name="hobby">跳舞<!--设置多选框时可以让type的属性值为“checkbox”-->
    <input type="checkbox" name="hobby">RAP<!--设置多选框时可以让type的属性值为“checkbox”-->
    <input type="checkbox" name="hobby">打篮球<!--设置多选框时可以让type的属性值为“checkbox”-->
    <br>
    <br>
    <input type="submit"><!--提交按钮的作用是把我们这个表单中填写的这些数据提交到创建form标签时自带的action指向的URL-->
    <form action="#"></form><!--action的属性值表示我们提交标签的时，也就是点击这个提交按钮时，向何处发送我们这个填写的数据-->
</form>
```

####  7. 单标签

单标签表示没有内容的元素，它们不需要结束标签。常见单标签如下：

```html
<input type="text">
<br>
<hr>
<img>
```

####  8. 无语义容器标签

无语义标签本身没有任何具体意义或默认样式，它们的作用主要是组织和分组内容，或者配合 CSS 和 JavaScript 实现样式和行为。最常见无语义容器标签如下：

```html
`<div>`
`<span>`
```

|              | **`<div>`**                              | **`<span>`**                           |
| ------------ | ---------------------------------------- | -------------------------------------- |
| **标签类型** | 块级元素                                 | 行内元素                               |
| **显示方式** | 独占一行，默认宽度占满父容器             | 不独占一行，可与文字或其他元素并排显示 |
| **主要用途** | 分组或组织页面布局                       | 标记一小段文字，方便设置样式或操作行为 |
| **嵌套规则** | 可以嵌套块级元素和行内元素               | 只能嵌套行内元素                       |
| **常见场景** | 用作容器划分页面的头部、侧边栏或主内容等 | 给文字加颜色、样式，标记部分内容       |

![image-20241116184853481](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116184853481.png)

```html
<div class="content">
	<h1>标题</h1>
     <p>内容</p>
     <p>内容</p>
     <p>内容</p>
</div>
    <span>span标签</span>
    <span>span标签</span>
    <span>span标签</span>
    <span>span标签</span>
```



### 三.块级元素与行内元素



####  1. 比较

|              | **块级元素**                                                 | **行内元素**                                                 |
| ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **显示方式** | 独占一行                                                     | 与其他内容并排显示                                           |
| **宽度**     | 默认占满父容器的宽度                                         | 宽度由内容决定                                               |
| **嵌套规则** | 可以包含块级元素和行内元素                                   | 只能包含行内元素                                             |
| **应用场景** | 组织页面布局和分隔内容区域                                   | 为局部内容添加样式或功能                                     |
| **常见标签** | `<div>`、`<p>`、`<h1>`～`<h6>`、`<ul>`、`<ol>`、`<li>`、`<table>`、`<tr>`、`<td>`、`<form>` | `<span>`、`<a>`、`<b>`、`<i>`、`<button>`、`<input>`、`<img>`、`<br>` |



####  2. 转换

- CSS中，块级元素的默认 `display` 值是 `block`,可以通过 `display: inline` 转换为行内元素。
- CSS中，行内元素的默认 `display` 值是 `inline`,可以通过 `display: block` 转换为块级元素。



####  3. 行内块元素(补充)

* 水平方向上排列，但可以设置宽度、高度、内外边距等块级元素的属性。

* 行内块元素可以包含其他行内元素或块级元素。

* 行内块元素和块级元素都可以设置宽度和高度，

  但是行内元素无法设置宽度和高度（因为行内元素的宽度和高度是由它包含的内容决定的）。

  

### 四.HTML标签属性



####  1. 属性介绍：

* HTML 属性是定义标签行为和样式的重要工具。

* 属性为元素提供更多的信息。

  

####  2. 基本语法

```html
<开始标签 属性名="属性值">内容</开始标签>
```



####  3. 基本规则

* **属性名不区分大小写，属性值对大小写敏感**。

```html
<img src="example.jpg" alt=	"样图">
<img SRC="example.jpg" alt=	"样图">
<img src="EXAMPLE.JPG" alt=	"样图">
<!--前两者相同，第三个与前两个不同-->
```

* **每个标签都可以有一个或多个属性。**

  ```html
  <img src="example.jpg" alt=	"样图">
  ```

* **属性永远被写在开始标签中。**



####  4. 适用于大多数HTML元素的属性

| 属性  |                         描述                         |
| :---: | :--------------------------------------------------: |
| class | 为HTML，元素定义一个或多个类名（类名从样式文件引入） |
|  id   |                   定义元素唯一的id                   |
| style |                  规定元素的行内样式                  |

比如：

```html
<h1 id="title"></h1>
<div class="nav-bar"></div>
<h2 class="nav-bar"></h2>
```

​      

# 参考资料

1. [Html和css入门（更新中）](https://zhuanlan.zhihu.com/p/268900362)

2. [史上最全的HTML、CSS知识点总结，浅显易懂。适合入门新手](https://zhuanlan.zhihu.com/p/62967499)

3. [前端三剑客 HTML、CSS、JavaScript 入门到上手](https://zhuanlan.zhihu.com/p/526785618)

4. [二十分钟HTML快速入门 | 无废话且清晰流畅 | 手敲键盘 | WEB前端必备语言~](https://www.bilibili.com/video/BV1jf4y1J7Ms?vd_source=c91cd0d85dd28e19f062edef99115834)

5. [3小时前端入门教程（HTML+CSS+JS）](https://www.bilibili.com/video/BV1BT4y1W7Aw?vd_source=c91cd0d85dd28e19f062edef99115834)

