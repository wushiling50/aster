DIR = $(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
API_PATH = $(DIR)/api
GEN_PATH = $(DIR)/gen
RPC_PATH = $(DIR)/rpc
GO_MODULE := github.com/wushiling50/aster

# 服务名
SERVICES := analysis contribution developer id_generator relation repo
NO_DB_SERVICES := id_generator

.PHONY: init
init:
	sh ./script/init.sh
	
.PHONY: env-up
env-up:
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

$(SERVICES): gen-base

.PHONY: $(SERVICES)
$(SERVICES):
	goctl rpc protoc idl/$@.proto \
		--proto_path=. \
		--go_out=. \
		--go_opt=module=$(GO_MODULE) \
		--go-grpc_out=. \
		--go-grpc_opt=module=$(GO_MODULE) \
		--zrpc_out=$(RPC_PATH)/$@; 
	
	if echo '$(NO_DB_SERVICES)' | grep -wq '$@'; then \
		echo "Skipping database model generation for $@"; \
	else \
		echo "Generating database model for $@"; \
		goctl model mysql ddl \
			--dir ./pkg/model/$@ \
			--cache true  \
			--src ./config/sql/$@.sql; \
	fi

.PHONY: gen-base
gen-base:
	protoc idl/base.proto \
		--proto_path=. \
		--go_out=. \
		--go_opt=module=$(GO_MODULE) \
		--go-grpc_out=. \
		--go-grpc_opt=module=$(GO_MODULE)

# 格式化代码，我们使用 gofumpt，是 fmt 的严格超集
.PHONY: fmt
fmt:
	gofumpt -l -w .
