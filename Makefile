.PHONY: build install gen_proto default help 
.ONESHELL:
.SILENT: all

CURRDIRBASE := $(shell basename "${PWD}")
CURRDIRFULL := $(CURDIR)
NAME := $(shell uname -s)

define clear
	@clear > $(shell tty)
endef

define exit
	$(error 1)
endef

define log_print
	@printf "\n --- $(1) --- \n"
endef



# go_install:
# 	$(call clear)
# ifeq ($(NAME), Darwin) 
# 	$(call log_print,"Installs golang via brew")
# ifeq ($(shell ${USER}), root)
# 	$(call log_print,"Switch user to non root")
# 	$(call exit)
# endif
# 	@brew install golang > /dev/null
# endif
	
# ifeq ($(NAME), Linux) 
# 	$(call log_print, Installs golang via apt)
# 	@sudo apt install golang
# endif
default: help

help:
	@echo "ALL AVAILABLE INSTRUCTIONS!"
	@echo "\n"
	@echo "dev: builds, installs and generates proto"
	@echo "build: builds project"
	@echo "install: installs built project"
	@echo "gen_proto: generates proto file"
	@echo "\n"

dev: build install gen_proto

build:
	$(call clear)

	@$(call log_print, Assets transfer)
	@mkdir -p /usr/local/share/games/HideSeek/assets
	@mkdir -p /usr/local/share/games/HideSeek/log
	@mkdir -p /usr/local/share/games/HideSeek/db
	@mkdir -p /usr/local/share/games/HideSeek/pprof

	$(call clear)

	$(call log_print, Builds project)
	@go build

install: 
	$(call clear)
	$(call log_print, Installs project)
	@go install

gen_proto:
	$(call clear)
	@protoc -I internal/core/networking/api --go_out=. api.proto
	@protoc -I internal/core/networking/api --go-grpc_out=. api.proto
