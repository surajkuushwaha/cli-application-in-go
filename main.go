package main

import (
	"fmt"
	"slices"
)

func main() {
	// number
	// var age bool = true
	// fmt.Println("FPP", age)

	// Array
	// names := [5]string{"suraj", "kushwaha", "vscode", "go", "terminal"}
	// names[1] = "kkk"
	// fmt.Println("This is ", names[1])

	// Slice : dynamic array
	// names := []string{"suraj", "kushwaha", "vscode", "go", "terminal"}
	// names = append(names, "last value")
	// fmt.Println("This is ", names[len(names)-1])

	// Map key value pair : similar to js object
	// map[<key type>]<value type> : same type of restriction for the type of key like javascript

	// ages := map[string]int{
	// 	"suraj": 25,
	// 	"ekta":  24}
	// ages["jack"] = 45

	// fmt.Println("Suraj's age is ", ages["suraj"])
	// fmt.Println("ekta's age is ", ages["ekta"])
	// fmt.Println("jack's age is ", ages["jack"])

	// type Person struct {
	// 	Age  int
	// 	Name string
	// }

	// suraj := Person{
	// 	Age:  25,
	// 	Name: "suraj",
	// }

	// fmt.Println(suraj.Age)

	// ifConditions()
	loops()

}

func ifConditions() {
	powerLever := 500

	isOver9000 := powerLever >= 9000
	if isOver9000 {
		fmt.Println("It's over 9000")
	} else {
		fmt.Println("It is not over 9000")
	}
}

func panicChecks() {
	fruits := []string{"banana", "apple", "grapes", "pineapple"}

	fruit := fruits[0]

	switch fruit {
	case "banana":
		fmt.Println("It's banana")
	default:
		fmt.Println("not a banana")
	}

}

func loops() {
	fruits := []string{"banana", "apple", "grapes", "pineapple"}
	x := 10
	for x > 0 {
		fmt.Println("Hello", x)
		x--
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Println("It's ", fruits[i])
	}

	for i := range 10 {
		fmt.Println("Hello from range loop: ", i)

	}

	fruits2 := []string{"banana", "apple", "grapes", "pineapple"}

	for _, value := range fruits2 {
		fmt.Println("Fruit | ", value)
	}

	brands := []string{"apple", "sumsung", "oneplus", "mi", "vivo"}

	brandCategory := map[string]string{
		"apple":   "premium",
		"sumsung": "mix",
		"oneplus": "budget",
		"mi":      "low budget",
		"vivo":    "low budget",
	}

	brands = append(brands, "nothing")


	// iterate over array of stings
	for _, value := range brands {
		fmt.Println(value, "->", brandCategory[value])
	}

	fmt.Println("_________________________ Map loop ________________________")
	// iterate over map
	for key, value := range brandCategory {
		fmt.Println(key, "->", value)
	}

}
