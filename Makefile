# Basic variable definitions
BINARY_NAME=single-deploy
MAIN_FILE=cmd/single-deploy/main.go
GO=go
CGO_ENABLED=0
CONFIG_DIR=configs
BUILD_DIR=build
BUILD_CONFIG_DIR=build/configs

# Version information
VERSION=$(shell git describe --tags --always || echo "unknown")
BUILD_TIME=$(shell date "+%F %T")
GIT_COMMIT=$(shell git rev-parse --short HEAD || echo "unknown")
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD || echo "unknown")

# Build flags
LDFLAGS=-ldflags "\
	-w -s \
	-X 'main.Version=${VERSION}' \
	-X 'main.BuildTime=${BUILD_TIME}' \
	-X 'main.GitCommit=${GIT_COMMIT}' \
	-X 'main.GitBranch=${GIT_BRANCH}' \
"

# Platform specific builds
.PHONY: build-all
build-all: build-linux build-windows build-darwin

.PHONY: build-linux
build-linux:
	@echo "Building for Linux..."
	@mkdir -p ${BUILD_CONFIG_DIR}
	@cp -r ${CONFIG_DIR}/* ${BUILD_CONFIG_DIR}/
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=${CGO_ENABLED} \
		${GO} build ${LDFLAGS} \
		-o ${BUILD_DIR}/${BINARY_NAME}-linux-amd64 ${MAIN_FILE}
	@echo "Linux build complete"

.PHONY: build-windows
build-windows:
	@echo "Building for Windows..."
	@mkdir -p ${BUILD_CONFIG_DIR}
	@cp -r ${CONFIG_DIR}/* ${BUILD_CONFIG_DIR}/
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=${CGO_ENABLED} \
		${GO} build ${LDFLAGS} \
		-o ${BUILD_DIR}/${BINARY_NAME}-windows-amd64.exe ${MAIN_FILE}
	@echo "Windows build complete"

.PHONY: build-darwin
build-darwin:
	@echo "Building for macOS..."
	@mkdir -p ${BUILD_CONFIG_DIR}
	@cp -r ${CONFIG_DIR}/* ${BUILD_CONFIG_DIR}/
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=${CGO_ENABLED} \
		${GO} build ${LDFLAGS} \
		-o ${BUILD_DIR}/${BINARY_NAME}-darwin-amd64 ${MAIN_FILE}
	@echo "macOS build complete"

# Update clean target
.PHONY: clean
clean:
	@echo "Cleaning build files..."
	@go clean
	@rm -rf ${BUILD_DIR}
	@mkdir -p ${BUILD_DIR}