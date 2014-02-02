package main

import (
	"./greeting"
	"./sql"
	"fmt"
)

func main() {

	greeting.PrintMe("Hello Marcio from Main function!")
	g := greeting.GetIncbyFunction
	h := g(4)
	i := g(6)
	fmt.Printf("The value for h = %d\n", h(5))
	fmt.Printf("The value for i = %d\n", i(1))
	fmt.Println("---------------------------------------")
	//---------------------------------------

	salutations := greeting.Salutations{
		{"Marcio", "Yo"},
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
		{"John", "Too late"},
	}
	//_, _ = createMessage(greet.name, greet.greeting)
	//var prn = Print("Hi, ")
	prn := greeting.CreatePrintFunction("|:=)")
	salutations.Greet(prn, true)
	//Greet(greet, prn)
	//fmt.Println(createMessage(greet.name,greet.greeting)
	db.DBPing()
	// db.DBDelete("Panda")
	// db.DBInsert("Bella")
	db.DBUpdate("Bella", "Panda")
	// db.DBExec()
}
