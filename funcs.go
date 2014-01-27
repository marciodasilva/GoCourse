package main

import (
	"./greeting"
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

	var greet = greeting.Salutation{"Marcio", "Yo"}
	//_, _ = createMessage(greet.name, greet.greeting)
	//var prn = Print("Hi, ")
	prn := greeting.CreatePrintFunction("|:=)")
	greeting.Greet(greet, prn, false)
	//Greet(greet, prn)
	//fmt.Println(createMessage(greet.name,greet.greeting)
}