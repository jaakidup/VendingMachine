package main

import (
	"fmt"
	"testing"
)

func TestCashInput(t *testing.T) {

	ci := &CashInput{}

	ci.SetNeededAmount(1000)

	if ci.totalNeeded != 1000 {
		t.Error("Should be needing 1000")
	}

	ci.Insert(500)
	if ci.Check() {
		t.Error("Shouldn't pass amount check")
	}
	if ci.inserted != 500 {
		t.Error("Should have had 500 already inserted")
	}

	ci.Insert(500)
	if !ci.Check() {
		t.Error("Should have passed amount check")
	}

}
func TestCashInputChange(t *testing.T) {

	ci := &CashInput{}

	ci.SetNeededAmount(1000)

	ci.Insert(500)
	ci.Insert(200)
	ci.Insert(200)
	ci.Insert(200)
	fmt.Println("Accepted: ", ci.accepted)
	fmt.Println("Change ", ci.change)

	if !ci.accepted {
		t.Error("Shouldn't have been accepted")
	}

	if ci.change != 100 {
		t.Error("Should have been 100 change")
	}

	if ci.GiveChange() != 100 {
		t.Error("I need 100 change")
	}
}
