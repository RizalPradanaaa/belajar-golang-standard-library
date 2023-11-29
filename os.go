package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	for _, value := range args {
		fmt.Println(value)
	}

	host, err := os.Hostname()
	if err == nil {
		fmt.Println(host)
	}else{
		fmt.Println("Error", err.Error())
	}
}
