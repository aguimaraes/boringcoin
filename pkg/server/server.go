package server

import (
	"encoding/json"
	"fmt"
	"github.com/aguimaraes/boringcoin/pkg/transaction"
	"net/http"
)

var transactions []transaction.Transaction

func AddTransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	t := transaction.Transaction{}
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	transactions = append(transactions, t)
	println("New transaction:")
	fmt.Printf("From: %s\nTo: %s\nWhen: %s\nAmount: %d\n\n", t.From, t.To, t.When, t.Amount)
}
