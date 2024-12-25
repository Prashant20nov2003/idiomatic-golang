package main

import (
	"fmt"

	"github.com/Prashant20nov2003/ch16/sample_code/unexported_field_access/one_package"
	"github.com/Prashant20nov2003/ch16/sample_code/unexported_field_access/other_package"
)

func main() {
	huf := one_package.HasUnexportedField{
		A: 10,
		C: "hello",
	}
	fmt.Println(huf)
	other_package.SetBUnsafe(&huf)
	fmt.Println(huf)
}
