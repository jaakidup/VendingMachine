package main

// VendingMachine ...
type VendingMachine struct {
	Name string
	Display
	CashInput
}

// NewVM ...
func NewVM(name string) *VendingMachine {
	return &VendingMachine{
		Name: name,
	}
}
