CWD=$(shell pwd)
DIR_CONF=${CWD}/config

# dockerfile构建本地镜像
USER_VAR=$(shell echo $$USER)
.PHONY: mp
mp:
	docker build -t ${USER_VAR}/prometheus:v1 .

# 直接用官方镜像启动prometheus
.PHONY: prometheus
prometheus:
	docker run --name my-prometheus \
        -p 9090:9090 \
        -v ${DIR_CONF}/prometheus.yml:/etc/prometheus/prometheus.yml \
        prom/prometheus