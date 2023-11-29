package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("      Hello    ", " "))	// Memotong cutset di awal dan akhir string
	fmt.Println(strings.ToLower("hello world"))
	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.Split("hello world", " "))
	fmt.Println(strings.Contains("ha ha ha ha", "ha"))
	fmt.Println(strings.ReplaceAll("ha ha ha ha", "ha", "hi"))
}
