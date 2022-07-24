BINARY=bin/tls-server-rest
GOPACKAGE=cncc/tls-server-rest
BP_RUN_BASE=cncc/tls-server-rest
BP_RUN_GM=cncc/tls-server-rest:cnccgm
BP_RUN_SGM=cncc/tls-server-rest:scnccgm
UB_TIME=ubuntu:cncc

# 构建tls-server-rest
tls-server-rest:
	 go build -tags '!cnccgm' -o ${BINARY} -mod vendor
tls-server-rest-cncc:
	go build  -tags "cnccgm" -v -o ${BINARY} ${GOPACKAGE}


run:
	go run ${GOPACKAGE}
	
# 构建基础ubuntu镜像
ubuntu:
	docker build -t ${UB_TIME} -f ./script/docker/Dockerfile .
.PHONY: ubuntu

# 镜像中添加时区
docker-gm:
	docker build -t ${BP_RUN_GM} -f ./script/docker/Dockerfile.cncc .
.PHONY: docker-gm

docker-cnccgm:
	docker build -t ${BP_RUN_GM} -f ./script/docker/Dockerfile.cncco .
.PHONY: docker-cnccgm

# 软国密镜像
docker:
	docker  build --platform=linux/amd64 -t ${BP_RUN_SGM} -f ./script/docker/Dockerfile.scncc .
.PHONY: docker