GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
MKDIR=mkdir -p
EXPORT=export

SRC_ROOT=$(shell pwd)
SERVER_SRC=$(SRC_ROOT)/src/server/*.go
CLIENT_SRC=$(SRC_ROOT)/src/client/*.go


SERVER_NAME=srv
CLIENT_NAME=cli


build:
	$(MKDIR) $(SRC_ROOT)/files/input
	$(MKDIR) $(SRC_ROOT)/files/output
	$(GOBUILD) -o $(SERVER_NAME) $(SERVER_SRC)
	$(GOBUILD) -o $(CLIENT_NAME) $(CLIENT_SRC)
