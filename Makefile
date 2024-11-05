DIR = $(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
OUTPUT_PATH = $(DIR)/output
API_PATH = $(CMD)/api

.PHONY: env-up
env-up:
	sh init.sh
	@ docker compose -f ./docker/docker-compose.yml up -d

.PHONY: env-down
env-down:
	@ cd ./docker && docker compose down

# TODO: finish build 
.PHONY: aster
aster:
	mkdir -p output
	@echo "TODO"
	@echo "aster running"

# run ci
.PHONY: lint
lint:
	golangci-lint run --config=./.golangci.yml