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