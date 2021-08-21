# Summer2021-No.104 开发openEuler社区原生的运维管理平台

#### 介绍

https://gitee.com/openeuler-competition/summer-2021/issues/I3OCZ3

#### 软件架构

前端采用原生jQuery，后端采用Golang Gin框架。

#### 安装教程

1.  git clone git@gitee.com:openeuler-competition/summer2021-104.git
2.  cd summer2021-104/wbe
3.  go run main.go
4.  cd summer2021-104/wbe
5.  cp config/wfe.conf /etc/nginx/conf.d/
6.  systemctl start nginx

#### 使用说明

1.  浏览器打开 http://127.0.0.1
2.  输入系统管理员账户信息或运维工程师账户信息登陆，账号信息如下
    a) 系统管理员账号：admin 密码：200705
    b) 运维工程师账号：ops 密码：200705
3.  点击左侧导航栏使用系统

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)

#### 效率工具

- [x] [YAML to Go](https://yaml2go.prasadg.dev/)

- [x] [JSON to Go](https://mholt.github.io/json-to-go/)

- [x] [BootStrap](https://www.runoob.com/try/bootstrap/layoutit/)

- [x] [mysql --login-path](https://www.jianshu.com/p/feb178a677a2)
