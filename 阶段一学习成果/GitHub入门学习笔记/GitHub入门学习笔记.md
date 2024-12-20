### GitHub入门学习笔记大纲

#### 一、Git 和 GitHub介绍

##### 1. Git

##### 2. GitHub

#### 二、GitHub 的特殊词汇涵义

##### 1. Repository（仓库）

##### 2. Commit（提交）

##### 3. Branch（分支）

##### 4. Merge（合并）

##### 5. Pull Request（合并请求）

##### 6. Fork（复制仓库）

##### 7. Issues（问题追踪）

##### 8. Star（收藏）

##### 9. GitHub Actions（自动化操作）

##### 10. README 文件

#### 三．GitHub 的基础界面和功能介绍

##### 1. 首页侧边导航栏及功能介绍

##### 2. 首页个人菜单导航栏及功能介绍

##### 3. 仓库页面顶部导航栏及功能介绍 

---



### 一、Git 和 GitHub介绍



#### 1.Git：

**1）概念：**Git 是一个分布式版本控制软件，用于管理代码和文件的变更。

**2）特点：分布式。**Git允许每个开发者都拥有项目的完整副本(包括所有文件和历史版本)，使他们能够在不同地点独立地对项目进行修改和管理。而且，这些独立的修改可以在需要时进行合并和同步，确保项目的协同一致性。



#### 2.GitHub：

**1）概念：**GitHub 是一个基于 Git 的在线平台，用于文件托管和项目协作。

**2）特点：**GitHub允许开发者将项目上传以便进行团队协作和资源分享。它提供了一个共享、管理和协作代码的环境，既支持开源项目的公开分享，同时也支持私有项目的单独创建。



---



### 二、GitHub 的特殊词汇涵义



#### **1.Repository（仓库）：仓库是存放项目代码和文件的地方。**

1）可以在本地创建仓库，或者把本地仓库推送到 GitHub，以便进行云端存储和协作。

2）公有仓库允许任何人查看，私有仓库则仅允许授权用户访问。



#### **2.Commit（提交）：提交是对文件的更改记录。**

1）每次提交都会保存一次文件变更，记录代码修改的内容、时间和提交人等信息。

2）每次提交都有一个唯一的 ID，便于查找和回溯到特定历史版本。

 

#### **3.Branch（分支）：分支是 Github中的独立开发线。**

分支允许开发者在主分支之外并行开发不同功能，自由修改代码，最终可以将完成的功能合并回主分支。



#### **4.Merge（合并）：合并是将不同分支的更改共同整合到一个分支中。**

1）开发者在分支中完成新功能的开发后，可以通过合并将新功能整合到主分支中。

2）多个开发者在各自分支做出的更改后，也可以通过合并将更改整合到主分支中。

 

#### **5.Pull Request（合并请求）：合并请求是将更改合并到主分支的请求。**

当开发者在分支上完成更改后，可以向仓库的原始项目提交合并请求，以便将更改合并到主分支。



#### 6.Fork（复制仓库）：fork是生成其他用户项目副本的操作。

 Fork 功能允许用户将其他用户创建的项目复制到自己仓库里，并在其中自由修改，而不会影响其他用户的原始仓库。

 

#### 7.Issues（问题追踪）： Issue 是记录和追踪项目任务和问题的工具。

开发者可以在 Issues 中查看、创建和管理项目中的bug 、功能请求和任务进度，并和其他协作者讨论。

 

#### 8.Star（收藏）:用户可以点击star为自己喜欢的项目点赞和收藏。

Star 数量通常反映了项目的受欢迎程度。



####  9.GitHub Actions（自动化操作）：GitHub Actions 是 GitHub 的自动化工具 。

 GitHub Actions 允许用户设置一系列自动化操作（比如运行测试和代码部署），当代码被提交或合并时，系统会自动执行这些操作。

 

#### 10.README 文件：  README 是仓库中的介绍文档。

README文件通常包含项目的简介、安装和使用说明等内容，帮助其他用户快速了解项目的基本信息。

 

---



### 三．GitHub 的基础界面和功能介绍



#### 1.首页侧边导航栏及功能介绍

**1）Home：返回用户主页面。**用户可以在这里查看动态、通知和推荐项目。

**2）Issues：问题跟踪。**用户可以在这里查看、跟踪项目中的任务、功能请求和 bug。

**3）Pull requests：合并请求。**用户可以在这里查看和提交将更改合并到主分支的请求。

**4）Projects：项目管理。**用户可以在这里查看、管理和规划项目。

**5）Discussions：讨论区。**用户可以在这里交流和讨论问题、想法。

**6）Codespaces：代码空间。**用户可以在这里实现在浏览器中编辑和运行代码的功能。

**7）Explore：探索页面。**用户可以在这里浏览 GitHub 上的热门项目、开源代码库和开发者。

**8）Marketplace：应用商店。**用户可以在这里使用第三方集成工具来扩展 GitHub 的功能。

**9）Repositories：仓库。**用户可以在这里访问自己创建或参与的项目仓库。

 

#### 2.首页个人菜单导航栏及功能介绍

**1)Set status：设置状态。**用户可以在这里设置对外展示的状态（比如“在忙”“在线”）。

**2)Your profile：个人资料。**用户可以在这里查看设置个人相关信息（比如“用户名”“简介”）。

**3)Your repositories：个人仓库。**用户可以在这里访问自己创建和参与的所有仓库列表。

**4)Your Copilot：个人Copilot 。**用户可以在这里使用AI 编程助手GitHub Copilot（需付费）。

**5)Your projects：个人项目。**用户可以在这里查看、管理和规划项目。

**6)Your stars：个人收藏。**用户可以在这里查看自己收藏的项目和仓库列表。

**7)Your gists：个人代码片段。**用户可以在这里查看、管理和分享自己创建的代码片段。

**8)Your organizations：个人加入的组织。**用户可以在这里查看并进入自己参与的组织。

**9)Your enterprises：个人加入的企业。**用户可以在这里查看企业账户及相关资源。

**10)Your sponsors：个人赞助情况。**用户可以在这里查看自己的赞助对象和赞助商。

**11)Try Enterprise：试用企业版。**用户可以在这里体验为企业提供的高级服务。

**12)Feature preview：功能预览。**用户可以在这里体验 GitHub 的新功能。

**13)Settings：设置。**用户可以在这里管理个人账户信息、安全和通知等。

**14)GitHub Docs：GitHub文档。**用户可以在这里查阅帮助和使用指南。

**15)GitHub Support：GitHub支持。**用户可在这里寻求技术帮助和支持。

**16)GitHub Community：GitHub 社区。**用户可以在这里和其他用户交流。

**17)Sign out：登出。**用户可以在这里退出当前 GitHub账户。 

 

#### 3.仓库页面顶部导航栏及功能介绍 

**1)Code：代码文件。**用户可以在这里可以查看文件和提交代码。

**2)Issues：问题跟踪。**用户可以在这里查看、跟踪项目中的任务、功能请求和 bug。

**3)Pull requests：合并请求。**用户可以在这里查看和提交将更改合并到主分支的请求。

**4)Actions：自动化操作。**用户可以在这里查看和设置一系列自动化操作。

**5)Projects：项目管理。**用户可以在这里查看、管理和规划项目。

**6)Security：安全报告。**用户可以在这里查看仓库中的漏洞和潜在风险以及进行安全设置。

**7)Insights：数据分析。**用户可以在这里查看仓库的信息统计和使用情况。

**8)Settings：仓库设置。**用户可以在这里此设置仓库的基本信息。

 

 
