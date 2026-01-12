package main

type WinFactory struct {}

func (w WinFactory) createButton() Button {
	return WinButton{}
}

func (w WinFactory) createCheckbox() Checkbox {
	return WinCheckbox{}
}