FROM golang AS dev

WORKDIR /app
COPY . .

# 仅安装调试工具
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# 调试专用命令
CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient", "./src"]