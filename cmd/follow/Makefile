include ../common.mk

SERVICE_NAME = follow

.PHONY: server
server:
	sh build.sh
	cd output && sh bootstrap.sh

.PHONY: gen
gen:
	kitex \
	-gen-path ../../kitex_gen \
	-service "$(SERVICE_NAME)" \
	-module "$(MODULE)" \
	-type thrift \
	../../idl/$(SERVICE_NAME).thrift
	go mod tidy
