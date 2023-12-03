package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Rizal", "Nawang", "Pradana"}
	values := []int{10, 20, 30}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "rizal"))
	fmt.Println(slices.Index(names, "Pradana"))
}
