package main

// Stock ...
type Stock struct {
	items   StockItems
	soldOut StockItems
}

// Add ...
func (s *Stock) Add(item Item) {
	if s.items == nil {
		s.items = make(StockItems)
	}

	_, found := s.items[item.Code]
	if found {
		restock := s.items[item.Code]
		restock.Level += item.Level
		s.items[item.Code] = restock
	} else {
		s.items[item.Code] = item
	}

	_, exists := s.soldOut[item.Code]
	if exists {
		delete(s.soldOut, item.Code) // deleting regardless of whether it was there
	}
}

// RemoveItem from Stocklist
func (s *Stock) RemoveItem(item Item) {

	temp := s.items[item.Code]
	temp.Level--
	s.items[item.Code] = temp

	if s.items[item.Code].Level == 0 {
		s.MoveToSoldOut(item)
	}
}

// MoveToSoldOut ...
// Move item over to soldout list
func (s *Stock) MoveToSoldOut(item Item) {
	if len(s.soldOut) == 0 {
		s.soldOut = make(StockItems)
	}
	s.soldOut[item.Code] = item
	delete(s.items, item.Code)
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
