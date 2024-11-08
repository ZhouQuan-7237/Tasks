# GitHub进阶学习笔记-周全



## 学习背景



## 学习目录

#### 一.Git的学习

#####    1. Git工作原理

#####    2. Git基本操作

#####    3. Git与GitHub的关系

#### 二.Github的学习

#####    1. Fork（复制仓库）

#####    2. Pull Request（合并请求）

## 学习内容



### 一.Git的学习



#### 1.Git工作原理     



**Git有四个区域：**

- **Workspace（工作区）：**平时编辑和存放项目代码的地方。
- **Index / Stage（暂存区）：**暂时存放已修改但尚未提交的文件的地方。
- **Repository（仓库区或本地仓库）：**在本地存放提交历史和版本数据的地方。
- **Remote（远程仓库）:**  在远程服务器上存放提交历史和版本数据的地方。 



**文件在这四个区域之间的转换关系如下图所示：**



![git架构](https://www.wdg.pub/assets/git-svn-5.jpg)

 **Git 的基本工作流程：工作区修改文件 -> 添加文件到暂存区 -> 提交文件到本地仓库 -> 推送文件到远程仓库 -> 从远程仓库获取更新**

* **工作区修改编辑文件:**   在工作区编辑和修改文件。

* **添加文件到暂存区域：**将修改的文件从工作区添加到暂存区。

* **提交文件到本地仓库**：将修改的文件从暂存区提交到本地仓库。

* **推送文件到远程仓库：**将提交的文件从本地仓库推送到远程仓库，使本地仓库的更新同步到远程仓库。

* **从远程仓库获取更新：**将文件从远程仓库推送到本地仓库，使远程仓库的更新同步到本地仓库。

  

**文件的四种状态**:

- **Untracked（未跟踪）：** 文件在文件夹中, 但并没有加入到Git库。
- **Unmodify（未修改）：**文件已经入库，但未修改。
- **Modified（已修改）：**文件仅仅是修改, 并没有进行其他的操作。
- **Staged（暂存）：**文件已经进入暂存区。



#### 2.Git基本操作

**（1）创建本地仓库：**在一个空文件夹中使用`Git Bash`输入`$ git init`指令即可初始化一个本地仓库，文件夹会多出一个`.git`文件。

**（2）克隆远程仓库：**在一个空文件夹中使用`Git Bash`输入`$ git clone [远程仓库链接]` ,比如：

```text
$ git clone https://github.com/ZhouQuan-7237/Tasks.git
```

**（3）检查仓库状态：**使用`Git Bash`输入`$ git status`指令即可。

**（4）添加文件到暂存区：**使用`Git Bash`输入`$ git add .`指令即可。

**（5）提交暂存区文件到本地仓库：**使用`Git Bash`输入`$ git commit -m "提交的消息"`指令即可。比如：

```text
$ git commit -m "new file hello.md"
```

**（6）推送本地仓库文件到远程仓库：**使用`Git Bash`依次输入类似以下的指令即可。

```
$ git remote add origin https://github.com/ZhouQuan-7237/Tasks.git
$ git branch -M main 
$ git push -u origin main
```

**（7）从远程仓库获取更新：**使用`Git Bash`输入`$ git pull origin main`指令即可。



#### 3.Git与GitHub的关系

**Git 是核心，GitHub 是基于 Git 的协作平台。**

**(1)Git是核心：**Git是一个分布式版本控制系统，负责管理代码的版本历史和分支操作等核心功能。

**(2)GitHub是基于Git的协作平台：**GitHub 是在 Git 的基础上构建的在线协作平台，拓展了Git的功能，使其更适合团队协作和开源项目管理。



### 二.GitHub的学习



#### 1.Fork（复制仓库）：Fork是生成其他用户项目副本的操作。

Fork 操作可以将其他用户的仓库复制到自己的 GitHub 帐户中，

而且自己可以在其中自由修改，不会影响其他用户的原始仓库。



#### 2.Pull Request（合并请求）：Pull Request是请求将分支上的更改合并到主分支的操作。

当自己把Fork的其他用户仓库中的文件修改后，希望更新到其他用户的原始仓库时，

可以新建一个Pull Resquest，向其他用户发起合并请求，以便将分支上的更改合并到主分支。



即：用户1Fork用户2的仓库并修改了其中的文件，想把修改更新到用户2的仓库中，

那么他可以点击Pull Resquest，向用户2发起更新请求，

用户2点击Merge Pull Request和Confirm merge后即可完成更新合并。



## 参考资料

1.[廖雪峰的Git教程](https://liaoxuefeng.com/books/git/introduction/index.html)

2.[知乎Git 使用教程](https://zhuanlan.zhihu.com/p/135183491?utm_psn=1837968519708413952)

3.[停止重构Git 详解](https://mp.weixin.qq.com/s/mz9m1otMcL7BLvWxL_m1fQ)

4.[狂神说Git教程](https://mp.weixin.qq.com/s/Bf7uVhGiu47uOELjmC5uXQ)

5.[狂神说Git视频教程](【【狂神说Java】Git最新教程通俗易懂】https://www.bilibili.com/video/BV1FE411P7B3?p=9&vd_source=c91cd0d85dd28e19f062edef99115834)

6.[黑马程序员Github使用教程](【Github使用教程-黑马程序员】https://www.bilibili.com/video/BV1st411r7Sj?p=5&vd_source=c91cd0d85dd28e19f062edef99115834)
