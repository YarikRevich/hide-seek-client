.PHONY: go_install build install generate
.ONESHELL:
.SILENT: all

CURRDIRBASE := $(shell basename "${PWD}")
CURRDIRFULL := $(CURDIR)
NAME := $(shell uname -s)

define clear
	clear > $(shell tty)
endef

define exit
	$(error 1)
endef

define log_print
	@printf "\n --- $(1) --- \n"
endef

go_install:
	clear > $(shell tty)
ifeq ($(NAME), Darwin) 
	$(call log_print,"Installs golang via brew")
ifeq ($(shell ${USER}), root)
	$(call log_print,"Switch user to non root")
	$(call exit)
endif
	brew install golang > /dev/null
endif
	
ifeq ($(NAME), Linux) 
	$(call log_print,"Installs golang via apt")
	@sudo apt install golang
endif

build:
	$(call clear)
	$(call log_print,"Builds project")
	@go build -o HideSeek
install: 
	$(call clear)
	$(call log_print,"Installs project")
	@go install 
	@go generate ./...

generate:
	@$(shell $(CURDIR)/scripts/transfer_assets.sh)
	$(call log_print,"Assets transfered")
