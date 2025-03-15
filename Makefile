BIN_NAME=imgconv
BUILD_DIR=build

default: help

.PHONY: help
help:
	@echo "可用命令："
	@echo "  make help    显示此帮助信息"
	@echo "  make run     本地运行，以便开发"
	@echo "  make build   执行构建"
	@echo "  make install 执行本地安装命令,以便开发和自测"
	# 在这里添加其他命令的说明

.PHONY: install
install:
#	go install github.com/taadis/imgconv/cmd/imgconv
# 使用相对路径从本地安装,而不是从远端安装
	go install ./cmd/imgconv
