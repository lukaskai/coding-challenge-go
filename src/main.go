package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// TransactionBatch is the definition of a batch that would be sent to a bank account
type TransactionBatch struct {
	BatchNumber  int
	Amount       int
	IsDispatched bool
}

// AmountPayload is struct representing the body of a POST
type AmountPayload struct {
	Amount int
}

// Globals
var (
	batchHistory []TransactionBatch
	currentBatch TransactionBatch
	mu           = &sync.Mutex{}
)

func main() {

	//initialize first batch
	currentBatch = TransactionBatch{1, 0, false}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getBatchHistory(w)
	case http.MethodPost:
		var payload AmountPayload
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&payload)

		processTransaction(payload.Amount, &currentBatch, &batchHistory)
	default:
		http.Error(w, "Invalid request method.", 405)
	}
}

/*
*	getBatchHistory returns the information regarding the current un-dispatched batch,
*	and the history of dispatched batches in a human readable way
*
*	Parameters:
*	w, http.ResponseWriter: Output writer
 */
func getBatchHistory(w http.ResponseWriter) {

	fmt.Fprintf(w, "Current batch:\n")
	fmt.Fprintf(w, "\tBatch number: %d\n", currentBatch.BatchNumber)
	fmt.Fprintf(w, "\tAmount: %d\n", currentBatch.Amount)
	fmt.Fprintf(w, "\tDispatched: %t\n", currentBatch.IsDispatched)
	fmt.Fprintf(w, "-----------------------------------------\n\n")
	fmt.Fprintf(w, "Batch history:\n")

	if len(batchHistory) > 0 {
		for _, v := range batchHistory {
			fmt.Fprintf(w, "\tBatch number: %d\n", v.BatchNumber)
			fmt.Fprintf(w, "\tAmount: %d\n", v.Amount)
			fmt.Fprintf(w, "\tDispatched: %t\n", v.IsDispatched)
			fmt.Fprintf(w, "-----------------------------------------\n")
		}
	} else {
		fmt.Fprintf(w, "\n##### Batch history is empty #####")
	}
}

/*
*	processTransaction takes in a transaction amount, and adds it to the current unposted batch.
*	If the added amount to the current unposted batch goes over the threshold of 100,
*	it will be dispatched, and recorded into the batch history
*
*	Parameters:
*	amount, int: The amount of the transaction being posted
* 	batch, *TransactionBatch: The current unposted batch
*	ledger, *[]TransactionBatch: The collection of the dispatched batches
 */
func processTransaction(amount int, batch *TransactionBatch, ledger *[]TransactionBatch) {

	mu.Lock()

	// log.Printf("Adding %d to current amount of %d", amount, currentBatch.Amount)

	batch.Amount += amount

	if batch.Amount > 100 {

		// log.Printf("-------------Batch amount is over 100. Amount is: %d-------------", batch.Amount)

		batch.IsDispatched = true
		*ledger = append(*ledger, *batch)

		batch.BatchNumber = batch.BatchNumber + 1
		batch.Amount = 0
		batch.IsDispatched = false
	}

	mu.Unlock()
}
