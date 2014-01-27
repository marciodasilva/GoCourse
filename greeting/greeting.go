package greeting

import "fmt"

type inc func(digit int) int

type stringy func() string

type Printer func(string)

type Salutation struct {
	Name     string
	Greeting string
}

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	//fmt.Println(message + ", " + alternate)
	if prefix := "Mr."; isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = "Hi, " + name
	alternate = greeting
	return
}
func Print(s string) {
	fmt.Print(s)
}
func PrintMessage(s string) {
	fmt.Println("Hello PrintMessage, %s\n", s)
}

func Printline(s string) {
	fmt.Println(s)
}

func PrintMe(s string) {
	fmt.Print(s)
}

// Uses the closure process
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func GetIncbyFunction(n int) inc {
	return func(value int) int {
		return value + n
	}
}
