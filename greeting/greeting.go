package greeting

import "fmt"

type inc func(digit int) int

type stringy func() string

type Printer func(string)

type Salutation struct {
	Name     string
	Greeting string
}

func Greet(salutation []Salutation, do Printer, isFormal bool) {
	for _, s := range salutation {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := "Hello "; isFormal {
			do(prefix + s.Name + ", " + message)
		} else {
			do(alternate)
		}
	}
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	// message = greeting + " " + name + ", welcome to go programming "
	message = "welcome to go programming "
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
