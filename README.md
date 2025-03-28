# go-web
## 准备工作

app相关

log相关

mysql相关

redis相关

### 读取配置文件 viper

1、传入配置文件路径 flag.StringVar flag.parese

2、初始化配置文件信息

这里需要定义各类的结构体

`// 获取配置文件路径viper.SetConfigFile`

`// 读取配置文件err = viper.ReadInConfig()`

`// 读取配置文件,反序列化err = viper.Unmarshal(Conf)`

`// 监视viper.WatchConfig()viper.OnConfigChange(func(e fsnotify.Event)}`

### 初始化log、mysql、redis、雪花算法、gin

**log zap**

1、编码器 
`zapcore.EncoderConfig`
`zapcore.NewJSONEncoder`

2、日志输出
`zapcore.WriteSyncer`
`&lumberjack.Logger`
`zapcore.AddSync`

3、日志等级
`new(zapcore.Level)`

4、核心
`zapcore.NewCore(encoder, writerSyncer, l)`

`zap.New(core, zap.AddCaller())`

`// 全局替换zap.ReplaceGlobals(lg)`

**mysql sqlx**

1、配置信息

2、连接数据库  `sqlx.Connect("mysql", dsn)`

3、配置数据库

`db.SetMaxOpenConns(cfg.MaxOpen)`

`db.SetMaxIdleConns(cfg.MaxIdle)`

4、关闭数据库

**redis go-redis**

1、配置信息

2、连接redis  `redis.NewClient`

3、配置redis  `rdb.Ping(context.Background()).Result()`

4、关闭redis

**雪花算法 snowflake**

1、时间偏移

2、设置开始时间

3、机器码

`sf.Epoch = st.UnixNano() / 1000000`

`node, err = sf.NewNode(machineId)`

**gin框架**

1、gin.New()

2、中间件，在logger中定义

router.Use()

3、加载html文件 `router.LoadHTMLFiles`

4、加载静态文件 router.Static

5、设置组 api/v1 router.Goup

## 启动服务 http.Server

监听 srv.ListenAndServer 这里需要一个goroutin

优雅关机 

信号管道`make(chan os.Signal, 1)`

捕捉信号 

`signal.Notify(quit, syscall.*SIGINT*, syscall.*SIGTERM*)`

获取管道信号

设置上下文
`context.WithTimeout(context.Background(), 5*time.*Second*)`

关闭`srv.Shutdown(ctx)`

## 用户功能

### 用户注册 router.Post

参数信息

1、实例登录用户对象

2、绑定参数信息 **c.ShouldBindJson()**

3、验证参数信息 validator.ValidatorErrors

注册业务处理

1、查询数据库是否存在用户 **db.Get  select count**

2、生成用户id  **雪花算法 node.Generate().Int64()**

3、赋值到用户结构体

4、加密用户密码 md5 

`h := md5.New()`

`h.Write([]byte(*secret*))`

`hex.EncodeToString(h.Sum([]byte(str)))`

5、添加到数据库 **db.Exec insert into**

返回响应

### 用户登录

获取参数信息

1、实例化登录用户对象参数

2、绑定参数信息

3、验证参数信息

登录业务处理

1、实例化用户

2、数据查询

根据用户名查询用户密码

判断密码是否与传入密码相同（这里传入的密码也需要用md5加密后判断）

3、用户正确→生成token**(JWT)**

声明：(这里定义一个jwt的自定义的结构体)
`&CustomClaims{`uid,uname,`jwt.RegisteredClaims`{签发时间、签发人}}

创建前面对象`jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)`

生成token`token.SignedString(CustomSecret)`

返回响应

### 获取用户信息

根据用户id获取用户名称 **db.Get**

根据用户名称判断是否存在 **db.Get count**

## 身份验证：JWT中间件

1、从请求头部获取Authorization

2、判断→no→返回需登录

3、按空格划分拿到token，并验证

4、解析token

实例一个自定义声明JWT结构体

解析

`jwt.ParseWithClaims(tokenString, cm, func(token *jwt.Token) (interface{}, error) {    return CustomSecret, nil})`

5、拿到用户id

6、判断是否有用户

从redis中找到是否这个用户是否登录

并拿到其token

判断是否和传入的token有区别，如果有，就判断是多个用户登录

7、设置上下文，将用户id传入上下文

## 社区功能

### 创建社区

直接通过mysql创建

### 获取社区信息

获取全部社区信息 

[**db.Select](http://db.Select) 从mysql中的社区中获取全部列表，这里可以分页和偏移量也可以不用**

根据id获取社区信息

[**db.Select](http://db.Select)  从mysql中的社区中根据社区id获取**

## 帖子功能

### 发布帖子

获取参数信息

1、实例化个帖子对象

2、绑定参数

3、验证参数信息

获取用户id

通过中间件传入的上下文id拿到

赋值到实例化对象

帖子发布业务逻辑处理

用雪花算法生成一个帖子id赋值到id中

将数据传到数据库 **insert into db.Exec**

将帖子id和社区传入到redis中

3个键值对  pipeline 管道

帖子id和时间 ZAdd

帖子id和分数 ZAdd

梯子id和社区 SAdd

返回响应

### 获取帖子信息

获取帖子信息

1、获取参数信息 分页和偏移量

2、到数据库中获取帖子信息

到数据库中查找帖子信息 **selecr from db.Select**

根据帖子信息中返回的用户id获取作者姓名

根据帖子信息中返回的社区id获取社区的信息

3、合并为帖子详情，赋值到结构体

4、返回响应

根据id获取帖子信息

1、获取参数信息 id

2、到数据库中查找id到帖子信息

根据帖子信息中返回的用户id获取作者姓名

根据帖子信息中返回的社区id获取社区的信息

3、合并为帖子详情，赋值到结构体

4、返回响应

根据指定顺序获取帖子信息

1、获取参数信息

2、实例参数对象，分页，偏移量和指定顺序（时间和点赞数）

3、绑定参数，并验证

4、业务逻辑处理一、没给定社区id

1、从redis中获取排序安指定顺序排序的帖子id

判断要查询到key即上面的redis中一个键值对

按照分页偏移量获取**rdb.ZRevRange（从大到小获取)**

2、根据帖子id到数据获取详细信息

二、给定社区id

1、根据社区id在redis中获取

社区的key

指定顺序的key

**取交集rdb.ZInterStore**

```go
pipeline.ZInterStore(context.Background(), key, &redis.ZStore{
			Aggregate: "MAX",
			Keys:      []string{orderKey, ckey},
		})
		pipeline.Expire(context.Background(), key, time.Hour) // 设置超时时间
```

2、按照分页偏移量获取**rdb.ZRevRange（从大到小获取)**

3、根据帖子id到数据获取详细信息

5、返回响应

## 投票功能

### 投票

1、实例投票对象（梯子id，赞成、反对）

2、绑定参数信息

3、获取当前用户id

4、处理业务逻辑

更新redis中的投票信息

查看是否超过投票时间（一周）

当前用户的投票状态和之前的投票状态

更新帖子的票数

**pipeline ZIncrBy**

更新用户的投票信息

如果取消 就移除 **ZRem**

否则，添加 **ZAdd 用户id 投票状态**

5、返回响应

## 限流和性能测试

漏桶算法

```go

// 漏桶算法
func Ratelimit1(rate int) func(c *gin.Context) {
	rl := ratelimt1.New(rate) //生成一个限流器
	return func(c *gin.Context) {
		// 去水滴
		if rl.Take().Sub(time.Now()) > 0 {
			c.String(http.StatusOK, "reta limit ...")
			c.Abort()
			return
		}
		c.Next()
	}
}
```

令牌算法

```go
// 令牌算法
func Ratelimit2(rate time.Duration, capital int64) func(c *gin.Context) {
	// NewBucketWithQuantum 创建指定填充速率、容量大小和每次填充的令牌数的令牌桶
	// NewBucketWithRate 创建填充速度为指定速率和容量大小的令牌桶
	rl := ratelimt2.NewBucket(rate, capital) // 填充速率和容量
	return func(c *gin.Context) {
		// Take 可以赊账
		if rl.TakeAvailable(1) == 1 { // 有就能拿
			c.Next()
			return
		}
		c.String(http.StatusOK, "reta limit ...")
		c.Abort()
	}
}
```

**采集性能数据：**pprof 

```go
pprof.Register(router)
```

压测wrk

```go
go-wrk -n 50000 http://127.0.0.1:8080/book/list
go-torch -u http://127.0.0.1:8080 -t 30\\性能采集
```

## 部署

### 编译

makefile

```go
.PHONY: all build run gotool clean help

BINARY="bluebell"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ]; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
```

```go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/bluebell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/bluebell

```

### docker部署

dockerfile

```go
FROM golang:alpine AS builder

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录
WORKDIR /build

# 将代码复制到容器中
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执问价app
RUN go build -o bluebell_app .

###################
# 接下来创建一个小镜像
###################
FROM debian:buster

COPY ./wait-for.sh /
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf

# 从builder镜像中把/bulid/app 拷贝到当前目录
COPY --from=builder /build/bluebell_app /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 声明服务端口
EXPOSE 8084

# 启动容器时的运行命令
#ENTRYPOINT ["/bluebell_app", "conf/config.yaml"]

docker build . -t goweb_app
docker run -p 8888:8888 goweb_app
// 关联其他容器
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v /Users/q1mi/docker/mysql:/var/lib/mysql -d mysql:8.0.19
```

dock-compose

```go
FROM golang:alpine AS builder

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录
WORKDIR /build

# 将代码复制到容器中
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执问价app
RUN go build -o bluebell_app .

###################
# 接下来创建一个小镜像
###################
FROM debian:buster

COPY ./wait-for.sh /
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf

# 从builder镜像中把/bulid/app 拷贝到当前目录
COPY --from=builder /build/bluebell_app /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 声明服务端口
EXPOSE 8084

# 启动容器时的运行命令
#ENTRYPOINT ["/bluebell_app", "conf/config.yaml"]

docker-compose up
```

**nohup：**用于在系统后台**不挂断**地运行命令，不挂断指的是退出执行命令的终端也不会影响程序的运行。

**supervisor：一个普通的命令行进程变为后台守护进程，并监控该进程的运行状态，当该进程异常退出时能将其自动重启。**

```go
sudo yum install epel-release
// 安装
sudo yum install supervisor
// 修改配置 /etc/supervisord.d/
[include]
files = /etc/supervisord.d/*.conf
// 启动
sudo supervisord -c /etc/supervisord.conf
// 配置文件
[program:bluebell]  ;程序名称
user=root  ;执行程序的用户
command=/data/app/bluebell/bin/bluebell /data/app/bluebell/conf/config.yaml  ;执行的命令
directory=/data/app/bluebell/ ;命令执行的目录
stopsignal=TERM  ;重启时发送的信号
autostart=true  
autorestart=true  ;是否自动重启
stdout_logfile=/var/log/bluebell-stdout.log  ;标准输出日志位置
stderr_logfile=/var/log/bluebell-stderr.log  ;标准错误日志位置
// 更新配置
sudo supervisorctl update # 更新配置文件并重启相关的程序
sudo supervisorctl status bluebell

supervisorctl status       # 查看所有任务状态
supervisorctl shutdown     # 关闭所有任务
supervisorctl start 程序名  # 启动任务
supervisorctl stop 程序名   # 关闭任务
supervisorctl reload       # 重启supervisor

```

### **nginx**

```go
sudo yum install epel-release
//安装
sudo yum install nginx
//开机启动
sudo systemctl enable nginx
//启动
sudo systemctl start nginx
sudo systemctl status nginx

nginx -s stop    # 停止 Nginx 服务
nginx -s reload  # 重新加载配置文件
nginx -s quit    # 平滑停止 Nginx 服务
nginx -t         # 测试配置文件是否正确

worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       80;
        server_name  localhost;

        access_log   /var/log/bluebell-access.log;
        error_log    /var/log/bluebell-error.log;

        location / {
            proxy_pass                 http://127.0.0.1:8084;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}

```

负载均衡

```go
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    upstream backend {
      server 127.0.0.1:8084;
      # 这里需要填真实可用的地址，默认轮询
      #server backend1.example.com;
      #server backend2.example.com;
    }

    server {
        listen       80;
        server_name  localhost;

        access_log   /var/log/bluebell-access.log;
        error_log    /var/log/bluebell-error.log;

        location / {
            proxy_pass                 http://backend/;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}

```

## 接口文档 swag

```go
package main

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url 

// @host 这里写接口服务的host
// @BasePath 这里写base path
func main() {
	r := gin.New()

	// liwenzhou.com ...

	r.Run()
}

// contoller
// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]

go get -u github.com/swaggo/swag/cmd/swag
swag init

// import
	_ "bluebell/docs"  // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

// 注册路由
r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

http://localhost:8080/swagger/index.html

```

### air 热监督，启动后程序一直运行，有修改保存后自动编译运行

```go
# [Air](https://github.com/cosmtrek/air) TOML 格式的配置文件

# 工作目录
# 使用 . 或绝对路径，请注意 `tmp_dir` 目录必须在 `root` 目录下
root = "."
tmp_dir = "tmp"

[build]
# 只需要写你平常编译使用的shell命令。你也可以使用 `make`
# Windows平台示例: cmd = "go build -o tmp\main.exe ."
cmd = "go build -o ./tmp/main ."
# 由`cmd`命令得到的二进制文件名
# Windows平台示例：bin = "tmp\main.exe"
bin = "tmp/main"
# 自定义执行程序的命令，可以添加额外的编译标识例如添加 GIN_MODE=release
# Windows平台示例：full_bin = "tmp\main.exe"
full_bin = "./tmp/main ./conf/config.yaml"
# 监听以下文件扩展名的文件.
include_ext = ["go", "tpl", "tmpl", "html"]
# 忽略这些文件扩展名或目录
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 监听以下指定目录的文件
include_dir = []
# 排除以下文件
exclude_file = []
# 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间
delay = 1000 # ms
# 发生构建错误时，停止运行旧的二进制文件。
stop_on_error = true
# air的日志文件名，该日志文件放置在你的`tmp_dir`中
log = "air_errors.log"

[log]
# 显示日志时间
time = true

[color]
# 自定义每个部分显示的颜色。如果找不到颜色，使用原始的应用程序日志。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# 退出时删除tmp目录
clean_on_exit = true
```

### 测试单元
