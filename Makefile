.PHONY: build

prebuild:
	cd channel-to-slack/ && GOOS=linux GOARCH=amd64 go build -o channel-to-slack .
build:
	sam build
