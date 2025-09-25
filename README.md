# go_blog

#### 介绍
个人博客项目

#### 软件架构
软件架构说明:
整体架构：
```
├── go_blog
    ├── server (后端)
    └── web    (前端)
```
后端架构
```
├── server
    ├── api               (api层)
    ├── assets            (静态资源包)
    ├── config            (配置包)
    ├── core              (核心文件)
    ├── flag              (flag命令)
    ├── global            (全局对象)
    ├── initialize        (初始化)
    ├── log               (日志文件)
    ├── middleware        (中间件层)
    ├── model             (模型层)
    │   ├── appTypes      (自定义类型)
    │   ├── database      (mysql结构体)
    │   ├── elasticsearch (es结构体)
    │   ├── other         (其他结构体)
    │   ├── request       (入参结构体)
    │   └── response      (出参结构体)
    ├── router            (路由层)
    ├── service           (service层)
    ├── task              (定时任务包)
    ├── uploads           (文件上传目录)
    └── utils             (工具包)
        ├── hotSearch    (热搜接口封装)
        └── upload        (oss接口封装)
```

#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

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
