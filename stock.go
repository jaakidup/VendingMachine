package main

import "fmt"

// Stock ...
type Stock struct {
	items StockItems
}

// Add ...
func (s *Stock) Add(item Item) {
	if s.items == nil {
		s.items = make(StockItems)
	}
	s.items[item.Code] = item
}

// DumpList ...
func (s *Stock) DumpList() {
	fmt.Println(s.items)
}

// StockItems ...
type StockItems map[string]Item

// Item ...
type Item struct {
	Code  string
	Name  string
	Price int64
	Type  string
	Level int
}
