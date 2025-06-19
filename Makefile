DIR = $(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
API_PATH = $(DIR)/api
OUTPUT_PATH = $(DIR)/output

.PHONY: env-up
env-up:
	sh init.sh
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down

.PHONY: api-format
api-format:
	goctl api format --dir ${IDL_PATH}

.PHONY: api-go
api-go:
	goctl api go --dir=${API_PATH} --api ${IDL_PATH}/api.api