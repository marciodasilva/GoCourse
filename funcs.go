package main

import (
	"./greeting"
	"./sql"
	"./ws"
	"fmt"
	"github.com/codegangsta/martini"
)

func main() {

	fmt.Println("-----------Starting Martini-----------------")
	m := martini.Classic()
	// setting out the route
	m.Get("/", func() string { return ws.Json("i") })

	fmt.Println("-----------Starting Greeting logic-----------------")
	greeting.PrintMe("Hello Marcio from Main function!")
	g := greeting.GetIncbyFunction
	h := g(4)
	i := g(6)
	fmt.Printf("The value for h = %d\n", h(5))
	fmt.Printf("The value for i = %d\n", i(1))
	fmt.Println("---------------------------------------")

	salutations := greeting.Salutations{
		{"Marcio", "Yo"},
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
		{"John", "Too late"},
	}
	// _, _ = createMessage(greet.name, greet.greeting)
	//var prn = Print("Hi, ")
	prn := greeting.CreatePrintFunction("|:=)")
	salutations.Greet(prn, true)
	//Greet(greet, prn)
	//fmt.Println(createMessage(greet.name,greet.greeting)

	fmt.Println("-----------Starting DataBase transaction logic-----------------")
	db.DBPing()
	db.DBDelete("Panda")
	db.DBInsert("Bella")
	db.DBUpdate("Bella", "Panda")
	db.DBExec()
	db.DBSimpleExec()

	// Execute Json function to fetch the message.json file
	ws.Json("Hi")
	db.DBJson()
	// Run Martini Server
	m.Run()
}
