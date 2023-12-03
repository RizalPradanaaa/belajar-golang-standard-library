package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc:= time.Date(2009, time.March, 10, 1, 0,0,0 ,time.UTC)
	fmt.Println(utc.Local())


	// Duration
	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 10
	var duration3 time.Duration = duration1 - duration2

	fmt.Printf("duration1 : %d\n", duration3)
}
