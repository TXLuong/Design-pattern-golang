package main

import "fmt"

func CreateUI(os string) {
	var button Button
	var checkbox CheckBox
	// This function depends on concrete UI components
	if os == "windows" {
		button = WinButton{}
		checkbox = MacCheckbox{} // This is possible error when developers can create wrong family of UI components
	} else if os == "mac" {
		button = MacButton{}
		checkbox = MacCheckbox{}
	} else {
		fmt.Printf("ERR not found OS %s ", os)
		return
	}
	button.display()
	checkbox.display()
}
func main() {
	CreateUI("windows")
	CreateUI("mac")
	CreateUI("Ubuntu")
}