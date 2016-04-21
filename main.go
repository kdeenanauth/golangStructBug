package main

type LocalInterface interface {
	SomeMethod(string) error
	SomeOtherMethod(string) error
}

type LocalStruct struct {
	LocalInterface
	myOwnField string
}

func main() {
	var localInterface LocalInterface = &LocalStruct{myOwnField:"test"}
	
	localInterface.SomeMethod("calling some method")
}