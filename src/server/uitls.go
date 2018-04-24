package main


import (
	"os"
	"fmt"
)


func handle_error(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, ERROR_MESSAGE, e.Error())
		os.Exit(ERROR_CODE_AT_EXIT)
	}
}


func print_error(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, ERROR_MESSAGE, e.Error())
	}
}


func print_message(message string) {
	fmt.Println(message)
}
