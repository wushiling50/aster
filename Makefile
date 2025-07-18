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

DOCKER := contribution developer id_generator relation repo analysis api_processor fetcher api

.PHONY: env-up
env-up:
	sudo rm -f ./docker/data/mysql/mysql.sock
	@ docker compose -f ./docker/docker-compose.yml up -d

.PHONY: env-down
env-down:
	@ cd ./docker && docker compose down

.PHONY: aster-build-all
aster-build-all:
	for svc in $(DOCKER); do \
		docker build -f $(DOCKER_PATH)/Dockerfile.$$svc -t aster/$$svc:$(IMAGE_TAG) . ;\
	done

.PHONY: $(addprefix aster-build-,$(DOCKER))
$(addprefix aster-build-,$(DOCKER)): aster-build-%:
	docker build -f $(DOCKER_PATH)/Dockerfile.$* -t aster/$*:$(IMAGE_TAG) .

.PHONY: aster-run-all
aster-run-all: 
	for svc in $(DOCKER); do \
		if echo 'api' | grep -wq $$svc ; then \
			docker run -di -p 20001:20001 --name aster-$$svc --network aster aster/$$svc:$(IMAGE_TAG) ;\
		else \
			docker run -di --name aster-$$svc --network aster aster/$$svc:$(IMAGE_TAG) ;\
		fi \
	done

.PHONY: $(addprefix aster-run-,$(DOCKER))
$(addprefix aster-run-,$(DOCKER)): aster-run-%:
	if echo 'api' | grep -wq $* ; then \
			docker run -di -p 20001:20001 --name aster-$* --network aster aster/$*:$(IMAGE_TAG) ;\
	else \
			docker run -di --name aster-$* --network aster aster/$*:$(IMAGE_TAG) ;\
	fi

.PHONY: aster-remove-all
aster-remove-all:
	for svc in $(DOCKER); do \
		docker stop aster-$$svc ;\
		docker rm aster-$$svc ;\
	done

.PHONY: $(addprefix aster-remove-,$(DOCKER))
$(addprefix aster-remove-,$(DOCKER)): aster-remove-%:
	docker stop aster-$* ;\
	docker rm aster-$*

# -----------------------
.PHONY: api-go
api-go:
	goctl api format --dir ${IDL_PATH}
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
