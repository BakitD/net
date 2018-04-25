package utils

import (
	"fmt"
	"os"
)

func Handle_error(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, ERROR_MESSAGE, e.Error())
		os.Exit(ERROR_CODE_AT_EXIT)
	}
}

func Print_error(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, ERROR_MESSAGE, e.Error())
	}
}

func Print_message(message string) {
	fmt.Println(message)
}
