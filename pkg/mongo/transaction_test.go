package mongo_test

import (
	"context"
	"testing"
)



func Test_Commit(t *testing.T) {

	tx, err := instance.NewTx(context.Background())
	if err != nil {
		t.Errorf("NewTx() failed: %v", err)
		return
	}

	_, err = queries.FetchAllStockBatch(tx, stockName)
	if err != nil {
		t.Errorf("FetchAllStockBatch() failed: %v", err)
		return
	}

	if err := tx.Commit(); err != nil {
		t.Errorf("Commit() failed: %v", err)
		return
	}

}



func Test_Rollback(t *testing.T) {

	tx, err := instance.NewTx(context.Background())
	if err != nil {
		t.Errorf("NewTx() failed: %v", err)
		return
	}

	batch, err := queries.FetchAllStockBatch(tx, stockName)
	if err != nil {
		t.Errorf("FetchAllStockBatch() failed: %v", err)
		return
	}
	count := len(batch)

	if err := queries.InsertStockBatch(tx, stockName, stockBatch); err != nil {
		t.Errorf("InsertStockBatch() failed: %v", err)
		return
	}

	if err := tx.Rollback(); err != nil {
		t.Errorf("Rollback() failed: %v", err)
		return
	}

	batch, err = queries.FetchAllStockBatch(tx, stockName)
	if err != nil {
		t.Errorf("FetchAllStockBatch() failed: %v", err)
		return
	}
	updatedCount := len(batch)

	if updatedCount != count {
		t.Errorf("count = %d, updatedCount = %d, Rollback() does not works", count, updatedCount)
		return
	}
}



func Test_CommitAfterRollback(t *testing.T) {

	tx, err := instance.NewTx(context.Background())
	if err != nil {
		t.Errorf("NewTx() failed: %v", err)
		return
	}

	if err := queries.InsertStockBatch(tx, stockName, stockBatch); err != nil {
		t.Errorf("FetchAllStockBatch() failed: %v", err)
		return
	}

	if err := tx.Rollback(); err != nil {
		t.Errorf("Rollback() failed: %v", err)
		return
	}

	if err := tx.Commit(); err == nil {
		t.Errorf("Commit() = nil, expected = error")
		return
	}

}