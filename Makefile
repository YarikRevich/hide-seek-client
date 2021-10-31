.PHONY: go_install build install_bin update_assets full_install
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

go_install:
	$(call clear)
ifeq ($(NAME), Darwin) 
	$(call log_print,"Installs golang via brew")
ifeq ($(shell ${USER}), root)
	$(call log_print,"Switch user to non root")
	$(call exit)
endif
	brew install golang > /dev/null
endif
	
ifeq ($(NAME), Linux) 
	$(call log_print, Installs golang via apt)
	@sudo apt install golang
endif

dev: build install_bin

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

install_bin: 
	$(call clear)
	$(call log_print, Installs project)
	@go install

update_assets:


full_install:
	$(MAKE) go_install
	$(MAKE) build
	$(MAKE) install_bin
	$(MAKE) update_assets	
