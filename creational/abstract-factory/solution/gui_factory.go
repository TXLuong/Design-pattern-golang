package main

type GUIFactory interface {
	createButton() Button
	createCheckbox() Checkbox
}