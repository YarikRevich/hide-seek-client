.PHONY: help test gen_proto build install
.ONESHELL:
.SILENT: all

help:
	@echo "These are all available commands:\n\n--build: builds all the components('hide-seek-server', 'hide-seek-client', 'hide-seek-services')\n--install: installs all the components('hide-seek-server', 'hide-seek-client', 'hide-seek-services')\n";

test:
	@bazel test //...

gen_proto:
	@bazel build //api:server_external_proto_go
	@bazel build //api:services_external_proto_go
	
	@cp bazel-bin/api/server_external_proto_go_/internal/core/networking/api/server_external/server_external.pb.go internal/core/networking/api/server_external
	@cp bazel-bin/api/services_external_proto_go_/internal/core/networking/api/services_external/services_external.pb.go internal/core/networking/api/services_external

build: gen_proto
	@bazel build //cmd:hide-seek-client

build:
	@mkdir -p /usr/local/share/games/HideSeek/assets
	@mkdir -p /usr/local/share/games/HideSeek/log
	@mkdir -p /usr/local/share/games/HideSeek/db
	@mkdir -p /usr/local/share/games/HideSeek/pprof

	@go build -o HideSeek cmd/main.go

install: 
	@go install
