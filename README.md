# net


## Simple TCP client and TCP server


### Installation

> `git clone git@github.com:BakitD/net.git`

> `cd net`

> `make`

Run server (don't change port)

> `./srv`

Files that client could download is located in /files/output
directory.
Downloaded files are stored in /files/input.

To download file use client with next comand

> `./cli FILENAME`

FILENAME - is the name of file.



#### TODO

 - logging (change print_error, handle_error, print_message)
 - add multicasting for servers
 - add directory where files will be located for search
 - add application layer protocol
 - change schema of server and client interaction
 - format code