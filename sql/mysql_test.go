package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestRun(t *testing.T) {
	DBPing()
}

func TestUpdate(t *testing.T) {
	var res = DBUpdate("Bella", "Panda")
	var count, err = res.RowsAffected()
	fmt.Println("Rows count for UDPATE is = %d", count)
	if err != nil {
		t.Errorf("res.RowAffected() returned error: %s", err.Error())
	}
	if count != 1 {
		t.Errorf("Expected 1 affected row, got %d", count)
	}
}

func TestDelete(t *testing.T) {
	var res = DBDelete("Panda")
	var count, err = res.RowsAffected()
	fmt.Println("Rows count for DELETE is = %d", count)
	if err != nil {
		t.Errorf("res.RowAffected() returned error: %s", err.Error())
	}
	if count != 1 {
		t.Errorf("Expected 1 affected row, got %d", count)
	}
}

func TestInsert(t *testing.T) {
	var res = DBInsert("Bella")
	var count, err = res.RowsAffected()
	fmt.Println("Rows count for INSERT is = %d", count)
	if err != nil {
		t.Errorf("res.RowAffected() returned error: %s", err.Error())
	}
	if count != 1 {
		t.Errorf("Expected 1 affected row, got %d", count)
	}
}

func TestExec(t *testing.T) {
	// db.DBExec()
}
