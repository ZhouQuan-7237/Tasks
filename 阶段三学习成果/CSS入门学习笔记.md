# CSS入门学习笔记-周全

## 学习目录

### 一.CSS简介

####   1. CSS基础介绍

####   2. CSS基本语法

####   3. CSS基本规则

####   4. CSS导入方式

####   5. CSS常用属性

### 二.CSS选择器

* #### 元素选择器

* #### 类选择器

* #### ID选择器

* #### 通用选择器

* #### 子元素选择器

* #### 后代选择器

* #### 兄弟选择器

  * ##### 相邻兄弟选择器

  * ##### 通用兄弟选择器

* #### 伪类选择器

### 三.CSS盒子模型

### 四.CSS常见样式

#### 1.颜色

#### 2.透明度

#### 3.文本样式

### 五.CSS常见布局

* #### 普通流布局

* #### 浮动布局

* #### 定位布局

* #### 弹性布局

* #### 网格布局

### 六.CSS定位

* #### 简单介绍

* #### 常见定位

* #### 通俗理解

* #### 定浮优缺

---

## 学习内容



### 一.CSS简介



####   1. CSS基础介绍

* **定义：**CSS （层叠样式表）是一种于用于定义网页外观和布局的样式语言。

* **作用：**CSS为页面提供丰富的样式，包括字体、颜色、布局等，让页面更美观、结构更清晰。

  

####   2. CSS基本语法

```css
选择器{
	属性1：属性值1；
	属性2：属性值2；
}
```



####  3. CSS基本规则

* 一个选择器的声明中可以包含多条属性。

* 声明中的每一行属性后需以分号`;`结尾。

* 声明中的属性和值以“键值对”的形式出现。

####  

#### 4. CSS导入方式

* ##### 内联样式

  把CSS样式直接放在HTML标签中，在style属性中直接定义样式。

  ![image-20241116201807914](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116201807914.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <p style="color:red">这是一个应用内联样式的文本</p>
</body>
</html>
```

* ##### 内部样式表

  在`<head>`标签中添加`<style>`标签来定义样式

![image-20241116202335978](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116202335978.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        p {
            color:green;
        }
    </style>
</head>

<body>
    <p>这是一个应用内部样式表的文本</p>
</body>
</html>
```

* ##### 外部样式表

  先在外部的CSS文件中定义样式，然后在HTML文档的`<head>`标签里添加`<link>`标签，把CSS文件链接到这个HTML文档中

![image-20241116202510847](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116202510847.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <link rel="stylesheet" href="./style.css">
</head>

<body>
    <p>这是一个应用外部样式表的文本</p>
</body>
</html>
```

```css
p {
    color: blue;
}
```



**三种导入方式的优先级：内联样式>内部样式表>外部样式表**



####  4. CSS常用属性

* **font复合属性：**可以通过这一个属性设置多个样式。

  ![image-20241116210035071](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116210035071.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1 style="font-weight: bolder;font-size: 50px;color: aqua">这是一个font复合属性示例</h1>
</body>
</html>
```

* **line-height属性：**调节行间距。

  ![image-20241116210246846](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116210246846.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <p style="line-height: 40px;">这是一个line-height属性示例这是一个line-height属性示例这是一个line-height属性示例这是一个line-height属性示例这是一个line-height属性示例这是一个line-height属性示例这是一个line-height属性示例这是一个line-height属性示例</p>
</body>
</html>

```

* **width属性：**调节宽度。

* **height属性：**调节高度。

![image-20241116210713449](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116210713449.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .block{
            background-color: aqua;
            width: 200px;
            height: 100px;
        }
    </style>
</head>
<body>
    <div class="block">这是一个块级元素,宽200px,高100px</div>
</body>
</html>
```



### 二.选择器



CSS选择器允许我们对选择的一个或一组元素定义样式，**常见选择器**如下：

* **元素选择器**（根据标签选择）

* **类选择器**（根据类名选择）

* **ID选择器**（根据ID选择）

* **通用选择器**（选择所有元素）

* **子元素选择器**（选择直接嵌套在父元素内部的子元素）

* **后代选择器**（选择嵌套在父元素内的所有后代元素，无论这些元素是否直接嵌套）

  （子元素一定是后代元素，后代元素不一定是子代元素，

  所以子元素选择器可以控制子元素，但不能控制除了子元素之外的后代元素，

  后代选择器可以控制子元素和后代元素，但当子代选择器已经控制子元素时，后代选择器控制不了子元素）

* **兄弟选择器**

  * **相邻兄弟选择器**（选择在同一级别下的紧跟在选中元素之后的第一个兄弟元素）
  * **通用兄弟选择器**（选择在同一级别下的所有位于选中元素之后的兄弟元素）

* **伪类选择器**（选择HTML文档中元素的特定状态或位置）

![image-20241116231927821](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116231927821.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CSS选择器</title>
    <style>
        /* 元素选择器 */
        h2{
            color: aqua;
        }
        /* 类选择器 */
        .highlight{
            background-color: yellow;
        }
        /* ID选择器 */
        #header{
            font-size: 34px;
            font-weight: bolder;
        }
        /* 通用选择器 */
        *{
            font-family: KaiTi;
        }
        /* 子元素选择器 */
        .father > .son{
            color: coral;
        }
        /* 后代选择器 */
        .father p{
            color: blueviolet;
        }
        /* 相邻兄弟选择器 */
        h3 + p{
            background-color: red;
        }
        /* 伪类选择器 *//*也可以选中第一个子元素 ：first-child，也可选中第n个子元素 ：nth-child,也可以选中链接的状态 ：active*/
        #element:hover{
            background-color: aquamarine;
        }
        /* 伪元素选择器 
          ::before(在选中元素之前插入虚拟的内容)
          ::after(在选中元素之后插入虚拟的内容)
         用于创建一个虚拟元素并样式化它们 */
    </style>
</head>
<body>
    <h1>不同类型的CSS选择器</h1>

    <h2>这是一个元素选择器示例</h2>

    <h3 class="highlight">这是一个类选择器示例</h3>

    <h4 id="header">这是一个ID选择器示例</h4>

    <div class="father">
        <p class="son">这是一个子元素选择器示例</p>
        <div>
            <p class="grandson">这是一个后代选择器示例</p>
        </div>
    </div>

    <p>这是相邻兄弟选择器前面的p标签</p>
    <h3>这是一个相邻兄弟选择器示例</h3>
    <p>这是紧跟在相邻兄弟选择器后面的p标签</p>
    <p>这是没有紧跟在相邻兄弟选择器后面的p标签</p>

    <h3 id="element">这是一个伪类选择器示例</h3>
</body>
</html>
```



### 三.盒子模型



**盒子模型：文档中的每个元素都可以被看成是一个矩形的盒子，这个盒子包含了内容（content）、内边距（padding）、文本边框（border）和 外边距（margin）**。

![e065c41ed61e2b0e5f680ed866e46dbc_720](C:\Users\zq113\Documents\Tencent Files\1134547237\nt_qq\nt_data\Pic\2024-11\Thumb\e065c41ed61e2b0e5f680ed866e46dbc_720.png)

|       属性名       |                             说明                             |
| :----------------: | :----------------------------------------------------------: |
|  内容（content）   |            盒子包含的实际内容，比如：文本、图片等            |
| 内边距（padding）  | 围绕在内容的内部，是内容与边框之间的空隙间隔，可以使用`padding`属性来设置，为复合属性 |
| 文本边框（border） | 围绕在内边距的外部，是盒子的边界，可以使用border属性来设置，为复合属性 |
|  外边距（margin）  | 围绕在边框的外部，是盒子与其他元素之间的空间，可以使用margin属性来设置，为复合属性 |



![image-20241116233210506](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116233210506.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>盒子模型</title>
    <style>
        .demo{
            background-color: aqua;
            display: inline-block;
            border: 5px solid yellowgreen;
            padding: 50px;
            margin: 40px;
        }
        .border-demo1{
            background-color: crimson;
            width: 300px;
            height: 200px;
            border-style: solid;
            border-width: 10px;
            border-color: darkorchid;
        }
        .border-demo2{
            background-color: crimson;
            width: 300px;
            height: 200px;
            border-left:10px solid yellow;
        }
    </style>
</head>
<body>
    <div class="demo">盒子模型练习</div>
    <br>
    <div class="border-demo1">这是一个边框示例</div>
    <br>
    <div class="border-demo2">这是一个边框示例</div>
</body>
</html>
```



### 四.CSS常见样式



#### 1. 颜色

* **CSS中表示颜色的常见方法有两种**

​    **（1）直接使用颜色名**，比如red。

​    **（2）用RGB或Hex自定义颜色。**

（RGB是红绿蓝三原色，每个原色的取值范围在0~255之间，越接近0则越接近黑色，越接近255则越接近该原      色，而Hex则是对取值范围从0~255的16进制转换，对应取值范围为00到ff）



* **颜色可以赋值给任何需要以颜色为值的CSS语句**，比如，

   文本颜色color；

   背景颜色background color；

   边框颜色border color等。

![image-20241116235615992](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241116235615992.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .rgb {
        /* rgb(red,green,blue)*/
        color:rgb(235,70,224);
        }

        .hex {
        /* #rrggbb */
        color:#eb46e0;
        }

        .background-color {
        background-color:rgb(235,70,224);
        }

        .border-color {
        border: 3px solid rgb(235,70,224);
        }
    </style>
</head>
<body>
    <!-- 颜色，这里介绍RGB和Hex两种 -->
    <div class="rgb">用RGB表示颜色</div>
    <div class="hex">用Hex表示颜色</div>

    <!-- 颜色可以赋给任何需要以颜色为值的语句 -->
    <div class="background-color">背景颜色</div>
    <div class="border-color">边框颜色</div>
</body>
</html>

```

#### 2. 透明度

rgba属性可以调节当前文本或背景颜色的透明度，取值范围0~1，

opacity属性调整整个元素及所有子元素的透明度，取值范围也是0~1。

![image-20241117000151632](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241117000151632.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
    .background-color {
    background-color:rgb(235,70,224);
}

    .text-opacity {
    color: rgba(235,70,224,0.5);
    }

    .opacity {
    opacity:0.5;
    }
    </style>
</head>
<body>
    <!-- 透明度 -->
    <div class="background-color text-opacity">颜色透明度</div>
    <div class="background-color opacity">Opacity调节整体透明度</div>
</body>
</html>

```

#### 3. 文本样式

![image-20241117001148535](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241117001148535.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
    @font-face{
    font-family:"ZiYan";
    src: url("./ZiYan.ttf")
    }

    .text-1 {
    font-family: Ziyan,'sans-serif';
    }

    .text-2 {
    font-size: 24px;
    line-height: 32px;
    }

    .text-3 {
    font-weight: 200;
    font-style: italic;
    text-decoration: underline;
    }

    </style>
</head>
<body>
        <!-- 文本 -->
        <p class="text-1">文本样式1</p>
        <p class="text-2">文本样式2</p>
        <p class="text-3">文本样式3</p>
</body>
</html>

```



### 五.CSS常见布局



#### **布局的本质：摆盒子**



* #### 普通流布局（元素按文档默认规则从上到下、从左到右排列）

  ![image-20241117084637490](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241117084637490.png)

  ```html
  <!DOCTYPE html>
  <html lang="en">
  <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>Document</title>
  </head>
  <body>
      <div>块级元素1</div>
      <div>块级元素2</div>
      <span>行内元素1</span><span>行内元素2</span>
  </body>
  </html>
  ```

* #### 浮动布局（元素通过 `float` 属性脱离文档流，实现左右浮动）

  * ##### 浮动的三大特性

    * **脱标：**脱离标准流（当元素设置了浮动之后，元素就会脱离标准流的控制移动到指定的位置，浮动的盒子就不再保留原来所占的位置，就像漂浮在空中一样，脱离了原来的地面）。
    * **一行显示：**顶部对齐（如果多个盒子同时设置了浮动，那么他们就会按照属性值一行内显示，并且顶端对齐排列）。
    * **具备行内块元素特性**（不管原先是什么模式的元素，添加了浮动之后，都会具有行内块元素相似的特性）。

  * ##### 浮动和行内块的区别：

    浮动的元素是相互贴靠在一起的，不会有缝隙，如果父级宽度装不下这些浮动的盒子，多出去的盒子就会另起一行，而我们如果使用行内块元素的话，它彼此之间是由缝隙的，不如浮动这么方便。

  * **浮动的典型应用：**让多个块级元素在同一行内排列显示。

  * **浮动注意：**浮动是相对于父元素浮动，只会在父元素的内部移动。

  * **父元素的坍塌：**看图二，父元素边框包裹着所有元素，如果把父元素的高度去掉，盒子就出现了高度的坍塌，边框就没有包裹这两个浮动盒子，这两个子元素就浮动了，而此时如果我们想在父元素的后面再写标签的话，它就会出现两个浮动盒子中间，而不是两个浮动盒子的下面。

  * **浮动的清除：**

    常见的清除浮动方式有两种：

    * 在浮动元素的父元素中添加overflow属性，指定值为hidden（图三）；
    * 伪元素清除法：在浮动标签的父元素添加一个伪元素（图四）。

    

  ![image-20241117084814582](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241117084814582.png)

  ```html
  <!DOCTYPE html>
  <html lang="en">
  <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>浮动</title>
      <style>
          .father{
              background-color: aqua;
              height: 150px;
              border: 3px solid brown;
          }
          .left-son{
              width: 100px;
              height: 100px;
              background-color: pink;
              float: left;
          }
          .right-son{
              width: 100px;
              height: 100px;
              background-color: yellow;
              float: right;
          }
          .father1{
              background-color: aqua;
              /* height: 150px; */
              border: 3px solid brown;
          }
          .left-son1{
              width: 100px;
              height: 100px;
              background-color: pink;
              float: left;
          }
          .right-son1{
              width: 100px;
              height: 100px;
              background-color: yellow;
              float: right;
          }
          .father2{
              background-color: aqua;
              /* height: 150px; */
              border: 3px solid brown;
              overflow: hidden;
          }
          .left-son2{
              width: 100px;
              height: 100px;
              background-color: pink;
              float: left;
          }
          .right-son2{
              width: 100px;
              height: 100px;
              background-color: yellow;
              float: right;
          }.father3{
              background-color: aqua;
              /* height: 150px; */
              border: 3px solid brown;
          }
          .father3::after{
              content: "";
              display: table;
              clear: both;
          }
          .left-son3{
              width: 100px;
              height: 100px;
              background-color: pink;
              float: left;
          }
          .right-son3{
              width: 100px;
              height: 100px;
              background-color: yellow;
              float: right;
          }
      </style>
  </head>
  <body>
      <div class="father">
          <div class="left-son">左浮动</div>
          <div class="right-son">右浮动</div>
      </div>
      <br>
      <div class="father1">
          <div class="left-son1">左浮动</div>
          <div class="right-son1">右浮动</div>
      </div>
      <p>这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本</p>
      <br>
      <div class="father2">
          <div class="left-son2">左浮动</div>
          <div class="right-son2">右浮动</div>
      </div>
      <p>这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本</p>
      <br>
      <div class="father3">
          <div class="left-son3">左浮动</div>
          <div class="right-son3">右浮动</div>
      </div>
      <p>这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本这是一段文本</p>
  </body>
  </html>
  ```

* #### 定位布局（让元素按照特定模式定位在页面中的某个位置）（详见定位）

* #### 弹性布局（一维布局，灵活实现水平或垂直居中对齐，适合导航栏等场景）

  ![image-20241117090347644](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241117090347644.png)

  ```html
  <!DOCTYPE html>
  <html lang="en">
  <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>Document</title>
      <style>
      .flex {
      display: flex;
      /*flex-direction: column;*/
  
      height: 120px;
      background-color:violet; 
      justify-content: center;
      align-items:center;
      }
  
      .item1 {
      flex: 1;
      display: flex;
      justify-content: center;
      align-items: center; 
      }
  
      .item2 {
      flex: 1;
      display: flex;
      justify-content: center;
      align-items: center; 
      }
  
      .item3 {
      flex: 1;
      display: flex;
      justify-content: center;
      align-items: center; 
      }
      </style>
  </head>
  <body>
      <div class="flex">
          <div class="item1">项目</div>
          <div class="item2">项目</div>
          <div class="item3">项目</div>
      </div>
  </body>
  </html>
  ```

  

* #### 网格布局（二维布局，可同时精确控制行和列，适合复杂页面设计）

![image-20241117090548126](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241117090548126.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
    .grid {
    display: grid;
    /*grid-template-rows: 100px 200px 300px;
    grid-template-columns: 1fr 2fr;*/
    grid-template-columns: repeat(4,1fr);
    grid-auto-rows: 100px;
    grid-auto-columns: 1fr;

    grid-gap: 12px;
}
    </style>
</head>
<body>
    <div class="grid">
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
        <div>网格项目</div>
    </div>
</body>
</html>
```



### 六.定位



* #### 简单介绍：

  定位是使用 CSS的`position` 属性精确控制元素位置，让元素按照特定模式定位到页面中的某个位置的方法。

* #### 常见定位：

  * **相对定位：**相对于元素在文档流中的正常位置进行定位。
  * **绝对定位：**相对于其最近的已定位的父级元素进行定位，不占据文档流的位置。
  * **固定定位：**相对于浏览器窗口进行定位。不占据文档流固定在屏幕上的位置，不随滚动而移动。

* #### 通俗理解：

  * **相对定位：**我没搬家，但可以稍微移动位置。
  * **绝对定位：**我搬家了，不再占原来的位置。
  * **固定定位：**我贴在屏幕上，不跟页面一起滚动。

* #### 定位布局与浮动布局优缺点

  浮动布局灵活性高，但不易控制；

  定位布局控制精准，但不够灵活。

![image-20241117094502455](C:\Users\zq113\AppData\Roaming\Typora\typora-user-images\image-20241117094502455.png)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>定位</title>
    <style>
        .box1{
            height: 350px;
            background-color: aqua;
        }
        .box-normal{
            width: 100px;
            height: 100px;
            background-color: purple;
        }
        .box-relative{
            width: 100px;
            height: 100px;
            background-color: pink;
            position: relative;
            left: 120px;
            right: ;
            top: 40px;
            bottom: ;
        }
        .box2{
            height: 350px;
            background-color: yellow;
            margin-bottom: 300px;
        }
        .box-absolute{
            width: 100px;
            height: 100px;
            background-color: yellowgreen;
            position: absolute;
            left: 120px;
        }
        .box-fixed{
            width: 100px;
            height: 100px;
            background-color: salmon;
            position: fixed;
            right: 0;
            top: 300px;
        }
    </style>
</head>
<body>
    <h1>相对定位</h1>
    <div class="box1">
        <div class="box-normal"></div>
        <div class="box-relative"></div>
        <div class="box-normal"></div>
    </div>
    <h1>绝对定位</h1>
    <div class="box2">
        <div class="box-normal"></div>
        <div class="box-absolute"></div>
        <div class="box-normal"></div>
    </div>
    <h1>固定定位</h1>
    <div class="box-fixed"></div>
</body>
</html>
```



## 参考资料

1. [Html和css入门（更新中）](https://zhuanlan.zhihu.com/p/268900362)

2. [史上最全的HTML、CSS知识点总结，浅显易懂。适合入门新手](https://zhuanlan.zhihu.com/p/62967499)

3. [前端三剑客 HTML、CSS、JavaScript 入门到上手](https://zhuanlan.zhihu.com/p/526785618)

4. [CSS入门最全笔记](https://zhuanlan.zhihu.com/p/429033172)

5. [十七分钟CSS快速入门 | 无废话且清晰流畅 | WEB前端必备程序语言~](https://www.bilibili.com/video/BV1Ci4y1W7H7?vd_source=c91cd0d85dd28e19f062edef99115834)

6. [3小时前端入门教程（HTML+CSS+JS）](https://www.bilibili.com/video/BV1BT4y1W7Aw?vd_source=c91cd0d85dd28e19f062edef99115834)
