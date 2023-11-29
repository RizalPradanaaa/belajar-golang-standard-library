package main

import (
	"errors"
	"fmt"
)
var (
	validationError = errors.New("Validation Error")
	notFoundError = errors.New("Not Found Error")
)

func getId(id string) error {
	if id == "" {
		return validationError
	}

	if id != "eko" {
		return notFoundError
	}

	return nil
}
func main() {
	err := getId("eko")
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("Validation Error")
		}else if errors.Is(err, notFoundError) {
			fmt.Println("Not Found Error")
		}else{
			fmt.Print("Unkown Error")
		}
	}


}
