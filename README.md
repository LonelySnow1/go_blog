# go_blog

#### 介绍
个人博客项目

#### 软件架构
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

### 启动容器

```bash
docker run -itd --name mysql -p 3306:3306 -e  MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=blog_db -d mysql

docker run --name es -p 127.0.0.1:9200:9200 -e "discovery.type=single-node" -e "xpack.security.http.ssl.enabled=false" -e "xpack.license.self_generated.type=trial" -e "xpack.security.enabled=false" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -d elasticsearch:8.17.0

docker run --name redis -p 6379:6379 -d redis
```

### 启动服务

```bash
# 进入 server 文件夹，修改配置文件 config.yaml
go mod tidy
# 初始化 mysql
go run main.go -sql
# 初始化 es
go run main.go -es
# 创建管理员
go run main.go -admin
# 运行后端
go run main.go

# 进入 web 文件夹
npm install
# 运行前端
npm run dev
```

## 项目部署

### 部署环境

CentOS（linux）

### 环境准备

```bash
# 安装 docker
yum install -y docker-ce
systemctl start docker
systemctl enable docker

# 安装 supervisor
yum install -y supervisor
systemctl start supervisord
systemctl enable supervisord

# 安装 nginx
yum install -y nginx
systemctl start nginx
systemctl enable nginx
```

### 准备工作

编译后端，得到 main 文件

```bash
# windows环境下，打开项目所在目录，进入 server 文件夹，打开 cmd （不是 powershell）
set GOOS=linux
set GOARCH=amd64
go mod tidy
go build main.go
```

编译前端，得到 dist 文件夹

```bash
# windows环境下，打开项目所在目录，进入 web 文件夹，打开 cmd
npm install
# 请将 index.html 中 http://127.0.0.1:8080/api/website/logo 替换为您的域名 https://www.your_domain/api/website/logo
npm run build
```

### 服务端目录

将文件按照下述目录上传

```bash
# /opt/go_blog
├── go_blog
    ├── server
    │   ├── data
    │   │   ├── es
        │   └── mysql
    │   ├── main
    │   └── config.yaml
    └── web
        └── dist
```

### 容器配置

### nginx配置

```

创建 /etc/nginx/conf.d/nginx.conf

将 your_domain 替换为你的域名，请自行获取 ssl 证书，上传证书文件至 /etc/nginx/cert/

```nginx
server {
    listen 80;
    server_name your_domain www.your_domain;
    return 301 https://www.your_domain$request_uri;
}

server { 
    listen 443 ssl; 
    server_name your_domain;  # 仅匹配非 www 的域名
    ssl_certificate /etc/nginx/cert/your_domain.crt; # 证书公钥
    ssl_certificate_key /etc/nginx/cert/your_domain.key; # 证书私钥
    return 301 https://www.your_domain$request_uri;  # 强制跳转到 www.your_domain
}

server {
    gzip on;
    gzip_vary on;
    gzip_disable "MSIE [1-6]\.";
    gzip_static on;
    gzip_min_length 256;
    gzip_buffers 32 8k;
    gzip_http_version 1.1;
    gzip_comp_level 5;
    gzip_proxied any;
    gzip_types text/plain text/css text/xml application/javascript application/x-javascript application/xml application/xml+rss application/emacscript application/json image/svg+xml;

    listen 443 ssl;
    server_name www.your_domain; # 多个域名⽤空格分开 
    ssl_certificate /etc/nginx/cert/your_domain.crt; # 证书公钥
    ssl_certificate_key /etc/nginx/cert/your_domain.key; # 证书私钥
    ssl_session_timeout 5m; 
    ssl_session_cache shared:MozSSL:10m;  # 设置会话缓存以提⾼性能 
    ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384;  # 配置加密算法 
    ssl_protocols TLSv1.2 TLSv1.3;  # 配置加密协议 
    ssl_prefer_server_ciphers on;

    add_header Strict-Transport-Security "max-age=63072000; includeSubDomains; preload" always; #可选配置，开启HSTS 
    add_header X-Frame-Options DENY; # 可选配置，防⽌点击劫持 
    add_header X-Content-Type-Options nosniff; # 可选配置，防⽌MIME类型嗅探 
    add_header X-XSS-Protection "1; mode=block"; # 可选配置，防⽌XSS攻击

    location / {
        try_files $uri $uri/ /index.html;
        root   /opt/go_blog/web/dist;
        index  index.html index.htm;
    }

    location /api/ {
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header REMOTE-HOST $remote_addr;
        proxy_pass http://127.0.0.1:8080/api/;
    }

    location /image {
        alias /opt/go_blog/web/dist/image;
    }

    location /emoji {
        alias /opt/go_blog/web/dist/emoji;
    }

    location /uploads/ {
        alias /opt/go_blog/server/uploads/;
    }
}
```

重新启动

```bash
systemctl restart nginx
```

### supervisor配置

给 main 执行权限，并初始化项目

```bash
# 进入 /opt/go_blog/server
chmod +x ./main

./main -sql
./main -es
./main -admin
```

创建 /etc/supervisord.d/go_blog.ini

```ini
[program: go_blog]
command=/opt/go_blog/server/main
directory=/opt/go_blog/server/
autorestart=true ; 程序意外退出是否自动重启
autostart=true ; 是否自动启动
user=root ; 进程执行的用户身份
stopsignal=INT
startsecs=1 ; 自动重启间隔
stopasgroup=true ;默认为false,进程被杀死时，是否向这个进程组发送stop信号，包括子进程
killasgroup=true ;默认为false，向进程组发送kill信号，包括子进程
```

重新启动

```bash
systemctl restart supervisord
```
