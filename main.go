package main

import (
	"fmt"
	"./packageA"
)

type SomeStruct struct {
	packageA.SomeInterface
	myOwnField string
}

func (b *SomeStruct) SomeMethod(a string) error {
	fmt.Println(a)
	return nil
}

func main() {
	a := SomeStruct{myOwnField:"test"}
	
	var b packageA.SomeInterface = &a
	b.SomeMethod("calling some method")
	b.SomeOtherMethod("test")
}