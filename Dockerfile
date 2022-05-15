FROM golang:latest

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
#GOOS指的是目标操作系统，支持以下操作系统
    GOOS=linux \
#GOARCH指的是目标处理器的架构，支持一下处理器架构
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
	
# 移动到工作目录：/home/www/goWebBlog 这个目录 是你项目代码 放在linux上  
# 这是我的代码跟目录 
# 你们得修改成自己的
WORKDIR /home/ubuntu/LearningWall

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件  可执行文件名为 LearningWall
RUN go build main.go

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /home/www/goWebBlog 目录复制到这里
RUN cp /home/ubuntu/LearningWall/LearningWall .

# 声明服务端口
EXPOSE 8081

# 启动容器时运行的命令
CMD ["/dist/main"]