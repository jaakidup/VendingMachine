package main

import "fmt"

// Display ...
type Display struct{}

// ShowMessage ...
func (d *Display) ShowMessage(message ...string) {
	fmt.Println("VM: ", message)
}
