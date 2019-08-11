Package
---
- go get
   - `github.com/astaxie/beego`
   - `github.com/jmoiron/sqlx`
   - `github.com/go-sql-driver/mysql`

当前项目目录
---
~~~
- src/
    - commons
        - EgoResult.go
        - DBUtils.go
    - user/
        - User.go
- static/
    - js/
    - css/
    - image/
- view/
    - login.html
main.go

~~~

**测试**
* 运行 `go run main.go`
* 访问 `http://localhost:8088`


目录结构 - 在 /mock_back 文件夹内
---

~~~
- conf/ 配置
   - app.conf
- controller/ 控制器
   - shop.go 
- maintain/
   - main.go 
- router/
   - router.go 路由
~~~


软件项目管理
---

**1. 瀑布式开发**
* 需求分析
* 设计↑↓
* 编码↑↓
* 测试↑↓
* 运行↑↓

> 优点
* 容易理解，易于管理
* 强调需求调查、产品测试
* 强调开发的阶段性早期测试，即编码前的工作
* 项目延期风险大

> 缺点
* 客户必须完整描述需求 - 可能客户不知道自己需要什么
* 阶段工作难以逆转
* 项目结束前都不能完整演示系统

**2. 敏捷性开发**
* 发布←→测试←→计划←→开发
* 需求确认
* 功能拆分
* 迭代开发
* 重构
* 完整发布

> 优点
* 增量开发，进度可控
* 持续测试和集成
* 快速试错，拥抱变化

> 缺点
* 项目成员稳定性要求高
* 文档管理弱