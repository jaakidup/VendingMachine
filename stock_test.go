package main

import "testing"

func TestStockAdd(t *testing.T) {

	stock := &Stock{}

	item := Item{
		Code:  "A1",
		Name:  "Kale",
		Price: 420,
		Type:  "Veg",
		Level: 10,
	}
	stock.Add(item)

	items := stock.items[item.Code]
	if items.Name != "Kale" {
		t.Error("Should be Kale")
	}
	if items.Level != 10 {
		t.Error("Should be 10")
	}
	if items.Price != 420 {
		t.Error("Should be 420")
	}
	if items.Type != "Veg" {
		t.Error("Should be Veg")
	}
}
