# 定义变量
BINARY_NAME=main
DOCKER_IMAGE_NAME=gin-online-chat

# 默认目标
.PHONY: help build run clean tidy deps vet test

help: ## 显示帮助信息
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## 构建项目
	go build -o ${BINARY_NAME} .

build-opt: ## 使用优化参数构建项目
	go build -ldflags="-w -s" -o ${BINARY_NAME} .

run: ## 运行项目
	go run main.go

clean: ## 清理构建产物
	rm -f ${BINARY_NAME}

tidy: ## 整理 Go 模块依赖
	go mod tidy

deps: ## 显示项目依赖
	go list -m all

deps-graph: ## 生成依赖关系图
	go mod graph | dot -Tpng -o deps.png

vet: ## 检查代码问题
	go vet ./...

test: ## 运行测试
	go test ./...

docker-build: ## 构建 Docker 镜像
	docker build -t ${DOCKER_IMAGE_NAME} .

docker-run: ## 运行 Docker 容器
	docker run -p 8080:8080 ${DOCKER_IMAGE_NAME}