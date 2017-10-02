package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aguimaraes/boringcoin/pkg/transaction"
)

var transactions []transaction.Transaction

func AddTransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}

	t := transaction.Transaction{}
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		panic(err)
	}
	defer r.Body.Close()
	transactions = append(transactions, t)
	println("New transaction:")
	fmt.Printf("From: %s\nTo: %s\nWhen: %s\nAmount: %d\n\n", t.From, t.To, t.When, t.Amount)
}
