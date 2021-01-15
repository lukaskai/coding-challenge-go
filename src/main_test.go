package main

import (
	"reflect"
	"testing"
)

func TestPostingTransactions(t *testing.T) {

	t.Run("test posting a single transaction to an empty batch", func(t *testing.T) {
		testBatch := TransactionBatch{1, 0, false}
		expected := TransactionBatch{1, 72, false}

		processTransaction(72, &testBatch, nil)

		if !reflect.DeepEqual(testBatch, expected) {
			t.Errorf("got %v want %v", testBatch, expected)
		}
	})

	t.Run("test posting a transaction that makes batch go over 100 in amount", func(t *testing.T) {
		var ledger []TransactionBatch
		testBatch := TransactionBatch{1, 99, false}
		expectedNewBatch := TransactionBatch{2, 0, false}
		expectedLedgerElement := TransactionBatch{1, 102, true}

		processTransaction(3, &testBatch, &ledger)

		if len(ledger) != 1 {
			t.Errorf("expected batch history to have length 1. got %v ", len(ledger))
		}

		if !reflect.DeepEqual(ledger[0], expectedLedgerElement) {
			t.Errorf("expected batch history element to be %v, but got %v", expectedLedgerElement, ledger[0])
		}

		if !reflect.DeepEqual(testBatch, expectedNewBatch) {
			t.Errorf("expected current batch to be %v, but got %v", expectedNewBatch, testBatch)
		}
	})
}
