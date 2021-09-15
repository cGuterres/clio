package main

import "os"

var (
	global  = 1010
	global2 = 2020
)

const (
	number1 = 3030
	number2 = 123
)

func main() {
	// variaveis
	var message = "Hello, world!"

	if number1 > 10 {
		println("Big")
	} else {
		println("Small")
	}

	switch number1 {
	case 100:
		println("Ok")
	case 200:
		println("Bla")
	default:
		println("Default")
	}

	println(message, number1, number2, global, global2)

	c := 10
	for c > 0 {
		println("Started")
		c--
	}

	args := os.Args

	for i := 0; i < len(args); i++ {
		println(args[i])
	}

	// or
	for i, arg := range args {
		println(i, " => ", arg)
	}

	// array
	arrays := []string{"Teste", "Teste2"}

	for i := 0; i < len(arrays); i++ {
		println(arrays[i])
	}
}
