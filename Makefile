.PHONY: build

build:
	sam build --use-container --container-env-var-file env.json
