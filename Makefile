.PHONY: all exit
.ONESHELL:
.SILENT: all exit

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

all:
	clear > $(shell tty)
ifeq ($(NAME), Darwin) 
# Creates the environment for the game
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

ifeq ($(NAME), Windows)
	$(call log_print,"Your OS is Windows, sorry dude ㋡ hahahhah")
	$(call exit)
endif
	$(call clear)
	$(call log_print,"Creates the environment for the game")

	mkdir -p $$GOPATH/src
	mv $(CURRDIRFULL) $$GOPATH/src

	$(call clear)
	$(call log_print,"Installs all the important packages for project")

	$(call clear)
	$(call log_print,"Build project")
	@go build main.go