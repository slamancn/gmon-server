# 使用 Node.js 14 官方镜像作为基础镜像
FROM node:14

# 设置工作目录为 /app
WORKDIR /app

# 复制 package.json 和 package-lock.json 到工作目录
COPY package*.json ./

# 安装应用程序的依赖
RUN npm install

# 复制当前目录的所有文件到工作目录
COPY . .

# 暴露容器内部的 3000 端口
EXPOSE 3000

# 定义容器启动时运行的命令
CMD ["node", "server.js"]
