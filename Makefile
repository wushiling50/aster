DIR = $(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
API_PATH = $(DIR)/api
GEN_PATH = $(DIR)/gen
RPC_PATH = $(DIR)/rpc
DOCKER_PATH = $(DIR)/docker
GO_MODULE := github.com/wushiling50/aster
IMAGE_TAG := 0.1

# 服务名
SERVICES := contribution developer id_generator relation repo analysis
NO_DB_SERVICES := id_generator

DOCKER_BUILD := contribution developer id_generator relation repo analysis api_processor fetcher api


.PHONY: env-up
env-up:
	sudo rm -f ./docker/data/mysql/mysql.sock
	@ docker compose -f ./docker/docker-compose.yml up -d

.PHONY: env-down
env-down:
	@ cd ./docker && docker compose down

.PHONY: api-go
api-go:
	goctl api format --dir ${IDL_PATH}
	goctl api go --dir=${API_PATH} --api ${IDL_PATH}/api.api

# TODO: add dockerfile build & run command
.PHONY: rpc-run
rpc-run: 
	echo "TODO"

.PHONY: api-run
api-run:
	go run ./rpc/developer/developer.go
	go run ./rpc/contribution/contribution.go
	go run ./rpc/relation/relation.go
	go run ./rpc/repo/repo.go
	
	go run ./rpc/analysis/analysis.go
	
	go run ./rpc/id_generator/idgenerator.go

	go run ./api/applet.go

.PHONY: aster-build-all
aster-build-all:
	for svc in $(DOCKER_BUILD); do \
		docker build -f $(DOCKER_PATH)/Dockerfile.$$svc -t aster/$$svc:$(IMAGE_TAG) . ;\
	done

.PHONY: $(addprefix aster-build-,$(DOCKER_BUILD))
$(addprefix aster-build-,$(DOCKER_BUILD)): aster-build-%:
	docker build -f $(DOCKER_PATH)/Dockerfile.$* -t aster/$*:$(IMAGE_TAG) .

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
