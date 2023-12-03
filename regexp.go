package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`z([a-z])l`)

	fmt.Println(regex.MatchString("zal"))
	fmt.Println(regex.MatchString("zel"))
	fmt.Println(regex.MatchString("zAl"))

	fmt.Println(regex.FindAllString("zal zil zul zAl ", 10))
}
