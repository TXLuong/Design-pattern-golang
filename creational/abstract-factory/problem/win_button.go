package main

import "fmt"

type WinButton struct{}

func (w WinButton) display() {
	fmt.Printf("Display WinButton \n")
}