package main

import (
	"errors"
	"fmt"
)

func createFactory(os string) (error, GUIFactory) {
	if os == "windows" {
		return nil, WinFactory{}
	} else if os == "mac" {
		return nil, MacFactory{}
	}
	fmt.Printf("ERR os %s not found \n", os)
	return  errors.New("OS NOT FOUND"), nil
}

func CreateUI(os string) {
	err, factory := createFactory(os)
	if err != nil {
		fmt.Printf("ERROR when showing UI \n")
		return
	}
	// Logic of this function doesn't depend on concrete UI Compoments
	factory.createButton().display()
	factory.createCheckbox().display()
}
func main() {
	CreateUI("windows")
	CreateUI("mac")
	CreateUI("Ubuntu")
}