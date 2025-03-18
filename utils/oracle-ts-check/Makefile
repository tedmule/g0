VERSION ?= $(shell date +%Y%m%d-%H%M)
APP_NAME ?= oracle-ts-check
IMAGE_NAME ?= $(APP_NAME)
IMAGE_TAG ?= $(VERSION)

# 编译目标的操作系统和架构
GOOS ?= linux
GOARCH ?= amd64

# 声明伪目标
.PHONY: all build docker clean

# 默认目标：编译并打包镜像
all: build docker

# 编译 Go 二进制
build:
	@echo "Building $(APP_NAME) for $(GOOS)/$(GOARCH)..."
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP_NAME) 
	@echo "Binary built: $(APP_NAME)"

# 打包 Docker 镜像
docker: build
	@echo "Building Docker image $(IMAGE_NAME):$(IMAGE_TAG)..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	@echo "Docker image built: $(IMAGE_NAME):$(IMAGE_TAG)"

# 查看版本
version:
	@echo "Version: $(VERSION)"