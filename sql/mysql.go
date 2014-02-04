package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DBPing() {
	var db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	err = db.Ping()
	panicIf(err)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func DBUpdate(oldValue string, newValue string) (res sql.Result) {
	var db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	// var tx, txerr = db.Begin()
	stmDel, err := db.Prepare("UPDATE task SET label=?  WHERE label=?")
	panicIf(err)
	res, err = stmDel.Exec(newValue, oldValue)
	var count, errR = res.RowsAffected()
	if errR != nil {
		fmt.Println("res.RowAffected() returned error: %s", errR.Error())
	}
	if count != 0 {
		fmt.Println("Expected 0 affected row, go %d", count)
	}
	// defer tx.Rollback()
	return res
}

func DBDelete(delLabel string) (res sql.Result) {
	if len(delLabel) == 0 {
		err := errors.New("Misssing parameter ...!")
		panicIf(err)
	}
	// Open database connection
	var db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	stmIns, err := db.Prepare("DELETE FROM task WHERE label =?")
	panicIf(err)

	res, err = stmIns.Exec(delLabel)
	panicIf(err)
	return res
}

func DBInsert(insValue string) (res sql.Result) {
	value := insValue
	// Open database connection
	var db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	stmIns, err := db.Prepare("INSERT INTO task set label =? ")
	panicIf(err)

	res, err = stmIns.Exec(value)
	panicIf(err)

	defer stmIns.Close()
	return res
}

func DBExec() {
	// Open database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}
func DBSimpleExec() {
	var (
		id    int
		label string
	)
	// Open database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total # of rows is %s\n", rows)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &label)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("The id is %d and name is %s\n", id, label)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
