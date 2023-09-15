# weblog

基于Go语言和beego框架 前端使用layui 布局 开发的个人博客系统

本项目灵感来源于：<https://github.com/Echosong/beego_blog>，笔者在此基础上进行了一些修改，将beego的版本升级到了v2版（原版本为v1版），增加了一些新功能，以及一些bug的修复。

# 使用说明

1.克隆到本地
```bash
git clone https://github.com/palp1tate/weblog.git
```
2.下载依赖

在项目根目录依次运行：
```bash
go mod init weblog
go mod tidy
```
3.修改app.conf
```go
# MYSQL
host = 127.0.0.1
port = 3306
username = root
password = pwd
database = weblog
```
4.数据库操作

首先手动创建数据库`weblog`，然后执行：
```bash
go build
./weblog
```
或者只执行`bee run`让项目跑起来。

项目跑起来后，它会自动建表，接下来运行`weblog.sql`给数据库插入数据。注意先跑项目再执行`sql`文件。

4.演示

<http://localhost:8080> (前台)

<http://localhost:8080/admin/login> (后台)
```
账号： Palp1tate  密码 :123456
```
前台部分页面：
![image](https://github.com/palp1tate/weblog/assets/120303802/d4955c9c-cacb-4816-92eb-e10eb2cd55b1)
![image](https://github.com/palp1tate/weblog/assets/120303802/76ecf537-33d4-4cab-b3e6-0f7837926132)
![image](https://github.com/palp1tate/weblog/assets/120303802/111af7f1-1bee-4715-a333-7deb71818bb0)
![image](https://github.com/palp1tate/weblog/assets/120303802/c9430372-40fe-4527-ab3f-adf4518e2311)
![image](https://github.com/palp1tate/weblog/assets/120303802/2431c882-e018-4a7a-bef2-a8a5bd39c1fd)
![image](https://github.com/palp1tate/weblog/assets/120303802/94f883a5-6466-4a2b-b190-77e7fbe27ded)
后台部分页面：
![image](https://github.com/palp1tate/weblog/assets/120303802/6253565c-2da5-4113-933d-296f303cbfd8)
![image](https://github.com/palp1tate/weblog/assets/120303802/85f9a7de-fe24-4d50-90dc-8754d9e2e751)
![image](https://github.com/palp1tate/weblog/assets/120303802/0034e25b-4e35-46cc-b1da-a5cb158e5c54)
![image](https://github.com/palp1tate/weblog/assets/120303802/e3aeb313-ad64-436f-9ea1-c9e2be5ce979)








