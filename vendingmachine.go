package main

// VendingMachine ...
type VendingMachine struct {
	Display
	CashInput
}

// NewVM ...
func NewVM() *VendingMachine {
	return &VendingMachine{}
}
