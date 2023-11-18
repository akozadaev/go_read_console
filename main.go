package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Чтение параметров
	// go run main.go 22 3 "Строка"
	arguments := os.Args
	first, _ := strconv.ParseFloat(arguments[1], 64)
	second, _ := strconv.ParseFloat(arguments[2], 64)
	asString := arguments[3]

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(asString)

	// Чтение из консоли, вывод в консоль
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)

	fmt.Println(name, age)

	// Чтение из консоли, вывод в консоль c потока os.Stdin:
	var name1 string
	var age1 int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Println(name1, age1)
}
