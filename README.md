# net


## Simple TCP client and TCP server

Simple TCP server and client for file transfer in local network


### Installation

> `git clone git@github.com:BakitD/net.git`

> `cd net`

> `export GOPATH=($pwd)`

> `make`

To remove all object files

> `make clean`

### Usage

Run server (-port is optional, default value is 9999)

> `./srv -port=PORT`


Directory /files/output includes files that client can
download. They are stored at servers side.
After download all files will be stored in /files/input directory.

To get file use client with next comand

> `./cli SERVER_ADDRESS FILES`

SERVER_ADDRESS - server IPv4 address with port in next format ip:port (example 172.10.20.30:9999)
FILES - file names divided with whitespaces
