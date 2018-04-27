GOCMD=go
GOBUILD=$(GOCMD) build
MKDIR=mkdir -p

SRC_ROOT=$(shell pwd)
SERVER_SRC=src/server/main/main.go
CLIENT_SRC=src/client/main/main.go

SERVER_NAME=srv
CLIENT_NAME=cli

build:
	$(MKDIR) $(SRC_ROOT)/files/input
	$(MKDIR) $(SRC_ROOT)/files/output
	$(GOBUILD) -o $(SERVER_NAME) ${GOPATH}/$(SERVER_SRC)
	$(GOBUILD) -o $(CLIENT_NAME) ${GOPATH}/$(CLIENT_SRC)

clean:
	@[ -f ./$(SERVER_NAME) ] && rm ./$(SERVER_NAME) || true
	@[ -f ./$(CLIENT_NAME) ] && rm ./$(CLIENT_NAME) || true

run:
	$(MKDIR) $(SRC_ROOT)/files/input
	$(MKDIR) $(SRC_ROOT)/files/output
	$(GOBUILD) -o $(SERVER_NAME) $(SERVER_SRC)
	$(GOBUILD) -o $(CLIENT_NAME) $(CLIENT_SRC)
	./srv ${ARGS}