package main

type MacFactory struct {}

func (m MacFactory) createButton() Button{
	return MacButton{}
}

func (m MacFactory) createCheckbox() Checkbox {
	return MacCheckbox{}
}