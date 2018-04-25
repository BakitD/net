GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

SRC_ROOT=$(shell pwd)
SERVER_SRC=$(SRC_ROOT)/src/server/*.go
CLIENT_SRC=$(SRC_ROOT)/src/client/*.go

SERVER_NAME=sstart
CLIENT_NAME=cstart


build:
	$(GOBUILD) -o $(SERVER_NAME) $(SERVER_SRC)
	$(GOBUILD) -o $(CLIENT_NAME) $(CLIENT_SRC)
