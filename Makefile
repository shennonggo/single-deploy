# 基础变量定义
BINARY_NAME=deployer
MAIN_FILE=cmd/deployer/main.go

# 版本信息
VERSION=$(shell git describe --tags --always || echo "unknown")
BUILD_TIME=$(shell date "+%F %T")
GIT_COMMIT=$(shell git rev-parse --short HEAD || echo "unknown")
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD || echo "unknown")

# Go 构建标记
GO=go
GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

# 构建标记
LDFLAGS=-ldflags "\
	-w -s \
	-X 'main.Version=${VERSION}' \
	-X 'main.BuildTime=${BUILD_TIME}' \
	-X 'main.GitCommit=${GIT_COMMIT}' \
	-X 'main.GitBranch=${GIT_BRANCH}' \
"

# 默认目标
.PHONY: all
all: clean build

# 清理构建文件
.PHONY: clean
clean:
	@echo "清理构建文件..."
	@go clean

# 构建应用
.PHONY: build
build:
	@echo "开始构建..."
	@echo "GOOS: ${GOOS}"
	@echo "GOARCH: ${GOARCH}"
	@echo "Version: ${VERSION}"
	@echo "BuildTime: ${BUILD_TIME}"
	@echo "GitCommit: ${GIT_COMMIT}"
	@echo "GitBranch: ${GIT_BRANCH}"
	@GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} \
		${GO} build ${LDFLAGS} \
		-o ./${BINARY_NAME} ${MAIN_FILE}
	@echo "构建完成: ./${BINARY_NAME}"