package transaction

import "time"

type Transaction struct {
	From   string    `json:"from"`
	To     string    `json:"to"`
	When   time.Time `json:"when"`
	Amount uint64    `json:"amount"`
}

func New(f string, t string, w time.Time, a uint64) Transaction {
	return Transaction{f, t, w, a}
}
