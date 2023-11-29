package main

import (
	"flag"
	"fmt"
)

func main() {

	host := flag.String("hostname", "localhost", "Put you hostname")
	username := flag.String("username", "root", "Put you username")
	password := flag.String("password", "root", "Put you Hostname")

	fmt.Println(*host, *username, *password)
}
