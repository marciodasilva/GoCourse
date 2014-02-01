package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

func DBPing() {
	err = db.Ping()
	panicIf(err)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func DBUpdate() {
	newValue := "Panda"
	oldValue := "Bella"
	stmDel, err := db.Prepare("UPDATE task SET label=?  WHERE label=?")
	panicIf(err)
	_, err = stmDel.Exec(newValue, oldValue)
	panicIf(err)
}

func DBDelete() {
	value := "Bella"
	stmIns, err := db.Prepare("DELETE FROM task WHERE label =? ")
	panicIf(err)

	_, err = stmIns.Exec(value)
	panicIf(err)
}

func DBInsert() {
	value := "Bella"
	stmIns, err := db.Prepare("INSERT INTO task set label =? ")
	panicIf(err)

	_, err = stmIns.Exec(value)
	panicIf(err)

	defer stmIns.Close()
}

func DBExec() {
	// Open database connection
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
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
