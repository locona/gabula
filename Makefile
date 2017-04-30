.DEFAULT_GOAL := all

# APP
all:
	@go run src/main.go

install:
	@sh script/installer
