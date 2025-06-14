package main

import "fmt"

func main() {
	// number
	var age bool = true
	fmt.Println("FPP", age)

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

}
