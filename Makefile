DIR = $(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
OUTPUT_PATH = $(DIR)/output

.PHONY: env-up
env-up:
	sh init.sh
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down