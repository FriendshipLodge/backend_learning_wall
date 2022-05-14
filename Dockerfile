FROM golang:latest

MAINTAINER gaoyang gaoyangbenyang@outlook.com

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
	
# 移动到工作目录：/home/ubuntu/fl_backend_learning_wall 这个目录 是你项目代码 放在linux上  
# 这是我的代码跟目录 
# 你们得修改成自己的
WORKDIR /home/ubuntu/fl_backend_learning_wall

# 将代码复制到容器中
COPY . .

# 声明服务端口
EXPOSE 8081

# 启动容器时运行的命令
CMD ["go run main.go"]