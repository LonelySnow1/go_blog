# go_blog

#### Introduction
Personal Blog Project

#### Software Architecture
Overall Architecture:
```
├── go_blog
    ├── server (backend)
    └── web    (frontend)
```
Backend Architecture
```
├── server
    ├── api               (api layer)
    ├── assets            (static resource package)
    ├── config            (configuration package)
    ├── core              (core files)
    ├── flag              (flag command)
    ├── global            (global objects)
    ├── initialize        (initialization)
    ├── log               (log files)
    ├── middleware        (middleware layer)
    ├── model             (model layer)
    │   ├── appTypes      (custom types)
    │   ├── database      (mysql structure)
    │   ├── elasticsearch (es structure)
    │   ├── other         (other structures)
    │   ├── request       (request parameter structure)
    │   └── response      (response parameter structure)
    ├── router            (router layer)
    ├── service           (service layer)
    ├── task              (scheduled task package)
    ├── uploads           (file upload directory)
    └── utils             (tool package)
        ├── hotSearch    (hot search interface encapsulation)
        └── upload        (oss interface encapsulation)
```

### Start Containers

```bash
docker run -itd --name mysql -p 3306:3306 -e  MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=blog_db -d mysql

docker run --name es -p 127.0.0.1:9200:9200 -e "discovery.type=single-node" -e "xpack.security.http.ssl.enabled=false" -e "xpack.license.self_generated.type=trial" -e "xpack.security.enabled=false" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -d elasticsearch:8.17.0

docker run --name redis -p 6379:6379 -d redis
```

### Start Services

```bash
# Enter the server folder and modify the configuration file config.yaml
go mod tidy
# Initialize mysql
go run main.go -sql
# Initialize es
go run main.go -es
# Create administrator
go run main.go -admin
# Run backend
go run main.go

# Enter the web folder
npm install
# Run frontend
npm run dev
```

## Project Deployment

### Deployment Environment

CentOS（linux）

### Environment Preparation

```bash
# Install docker
yum install -y docker-ce
systemctl start docker
systemctl enable docker

# Install supervisor
yum install -y supervisor
systemctl start supervisord
systemctl enable supervisord

# Install nginx
yum install -y nginx
systemctl start nginx
systemctl enable nginx
```

### Preparation

Compile the backend to get the main file

```bash
# Under Windows environment, open the directory where the project is located, enter the server folder, and open cmd (not powershell)
set GOOS=linux
set GOARCH=amd64
go mod tidy
go build main.go
```

Compile the frontend to get the dist folder

```bash
# Under Windows environment, open the directory where the project is located, enter the web folder, and open cmd
npm install
# Please replace http://127.0.0.1:8080/api/website/logo in index.html with your domain name https://www.your_domain/api/website/logo
npm run build
```

### Server-side Directory

Upload files according to the following directory structure

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

### Container Configuration

### Nginx Configuration

Create /etc/nginx/conf.d/nginx.conf

Replace your_domain with your domain name. Please obtain the SSL certificate by yourself and upload the certificate files to /etc/nginx/cert/

```nginx
server {
    listen 80;
    server_name your_domain www.your_domain;
    return 301 https://www.your_domain$request_uri;
}

server { 
    listen 443 ssl; 
    server_name your_domain;  # Only match non-www domain name
    ssl_certificate /etc/nginx/cert/your_domain.crt; # Certificate public key
    ssl_certificate_key /etc/nginx/cert/your_domain.key; # Certificate private key
    return 301 https://www.your_domain$request_uri;  # Force redirect to www.your_domain
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
    server_name www.your_domain; # Separate multiple domain names with spaces 
    ssl_certificate /etc/nginx/cert/your_domain.crt; # Certificate public key
    ssl_certificate_key /etc/nginx/cert/your_domain.key; # Certificate private key
    ssl_session_timeout 5m; 
    ssl_session_cache shared:MozSSL:10m;  # Set session cache to improve performance 
    ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384;  # Configure encryption algorithms 
    ssl_protocols TLSv1.2 TLSv1.3;  # Configure encryption protocols 
    ssl_prefer_server_ciphers on;

    add_header Strict-Transport-Security "max-age=63072000; includeSubDomains; preload" always; # Optional configuration, enable HSTS 
    add_header X-Frame-Options DENY; # Optional configuration, prevent clickjacking 
    add_header X-Content-Type-Options nosniff; # Optional configuration, prevent MIME type sniffing 
    add_header X-XSS-Protection "1; mode=block"; # Optional configuration, prevent XSS attacks

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
Restart

```bash
systemctl restart nginx
```

### Supervisor Configuration
Grant execution permission to main and initialize the project

```bash
# Enter /opt/go_blog/server
chmod +x ./main

./main -sql
./main -es
./main -admin
```
Create /etc/supervisord.d/go_blog.ini

```ini
[program: go_blog]
command=/opt/go_blog/server/main
directory=/opt/go_blog/server/
autorestart=true ; Whether to automatically restart if the program exits unexpectedly
autostart=true ; Whether to start automatically
user=root ; User identity for process execution
stopsignal=INT
startsecs=1 ; Automatic restart interval
stopasgroup=true ; Default is false. When the process is killed, whether to send a stop signal to this process group, including child processes
killasgroup=true ; Default is false. Send a kill signal to the process group, including child processes
```

Restart

```bash
systemctl restart supervisord
```
