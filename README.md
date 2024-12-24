# deployer

# 使用
只需要修改 `configs/deploy-config.json` 文件，然后执行 `./deployer` 即可。

# 安装依赖
go mod tidy

# 构建
go build -o deployer cmd/deployer/main.go

# 运行
./deployer