package main

import "testing"

func TestVendingMachine(t *testing.T) {

	vm := NewVM("Test VendingMachine")
	vm.ShowMessage("Hello")

	vm.SetNeededAmount(1000)
	vm.ShowMessage("Please insert 1000")
	if vm.totalNeeded != 1000 {
		t.Error("Need 1000 more")
	}

}
