package main

// Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode
// kita pada saat aplikasi sedang berjalan

import (
	"fmt"
	"reflect"
)

// Kode Program StructTag

type Sample struct {
	Name string `required:"true" max:"10"`
}

func main() {
	sample := Sample{"Rizal"}
	sampleType := reflect.TypeOf(sample)
	samplefield := sampleType.Field(0)
	required := samplefield.Tag.Get("required")

	fmt.Println(samplefield)
	fmt.Println(required)
}
