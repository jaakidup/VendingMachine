package main

func main() {
	vm := NewVM("My Shop System")
	vm.ShowMessage("Hello")

	// vm.dispenser.AddItem(Item{"Coke", 1000})
	// vm.dispenser.AddItem(Item{"Pepsi", 900})

	// vm.ListItems(vm.dispenser.GetStock())
}
