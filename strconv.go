package main

import (
	"fmt"
	"strconv"
)

// Konversi tipe data string ke lainnya

func main()  {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	}else{
		fmt.Println("Error", err.Error())
	}
}
