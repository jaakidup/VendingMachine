package main

// CashInput ...
type CashInput struct {
	totalNeeded int64
	inserted    int64
	change      int64

	busy     bool
	accepted bool
}

// SetNeededAmount ...
func (ci *CashInput) SetNeededAmount(amount int64) {
	if !ci.busy {
		ci.accepted = false
		ci.busy = true
		ci.totalNeeded = amount
	}
}

// Insert ...
func (ci *CashInput) Insert(amount int64) {
	if amount != 0 {
		ci.inserted += amount
		if ci.inserted >= ci.totalNeeded {
			ci.accepted = true
			ci.change = ci.inserted - ci.totalNeeded
			ci.busy = false
		}
	}
}

// Check ...
func (ci *CashInput) Check() bool {
	if ci.inserted >= ci.totalNeeded {
		ci.accepted = true
		ci.change = ci.inserted - ci.totalNeeded
		ci.busy = false
	}
	return ci.accepted
}

// GiveChange ...
func (ci *CashInput) GiveChange() int64 {
	return ci.change
}
