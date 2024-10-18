# 使用 Ubuntu 作为基础镜像
FROM ubuntu:20.04

# 避免交互式提示
ENV DEBIAN_FRONTEND=noninteractive

# 设置 Go 版本
ENV GO_VERSION=1.21.3

# 设置工作目录
WORKDIR /app

# 更新包列表并安装必要的工具
RUN apt-get update && apt-get install -y \
    wget \
    git \
    curl \
    tmpl \
    python3 \
    python3-pip \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

# 安装 Go
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz

# 设置 Go 环境变量
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# 安装 Python 依赖
RUN pip3 install ujson

# 安装 Go 依赖
RUN GO111MODULE=off go get -u \
    github.com/gogo/protobuf/protoc-gen-gofast \
    github.com/gogo/protobuf/proto \
    github.com/gogo/protobuf/jsonpb \
    github.com/gogo/protobuf/protoc-gen-gogo \
    github.com/gogo/protobuf/gogoproto \
    github.com/golang/protobuf/protoc-gen-go

# 设置工作目录
WORKDIR /app

# 如果有源代码需要复制到镜像中，可以使用 COPY 命令
# COPY . .

# 如果需要运行特定命令，可以使用 CMD 或 ENTRYPOINT
# CMD ["your-command-here"]
