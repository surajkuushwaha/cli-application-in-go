package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	foo, _ := os.ReadFile("./words.txt")

	fmt.Println("Word Count:", CountWords(foo))
}

// Task: Implement the count words function
func CountWords(data []byte) int {
	const spaceChar = ' '
	wordCount := 0
	for _, value := range data {
		if value == spaceChar {
			wordCount++
		}

	}
	wordCount++
	return wordCount
}

func bac() {
	// number
	// var age bool = true
	// fmt.Println("FPP", age)

	// Array
	// names := [5]string{"suraj", "kushwaha", "vscode", "go", "terminal"}
	// names[1] = "kkk"
	// fmt.Println("This is ", names[1])

	// Slice : dynamic array
	// names := []string{"apple", "orange", "pineapple", "papaya", "guava"}
	// names = append(names, "banana")
	// fmt.Println("Last Fruit is:", names[len(names)-1])

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
	// loops()
	// mapExample()
	// practicing()
	// mapExample()
	//
	// Create your instances here
	// printMultiples(5, 15)
	// printStrongest(map[string]int{"Goku": 9001, "Vegeta": 8500, "Gohan": 8000})
	fusedPower := fusionBoost(5000, 7000)
	fmt.Println("Fused power level:", fusedPower)
}

// Complete the function signature
func fusionBoost(powerLevel1 int, powerLevel2 int) int {
	return (powerLevel1 + powerLevel2) * 10
}

func printStrongest(powerLevels map[string]int) {
	// Your code here
	var highestCharacterName string
	var highestCharacterPower int
	for name, power := range powerLevels {
		if highestCharacterPower < power {
			highestCharacterPower = power
			highestCharacterName = name
		}
	}
	fmt.Println(highestCharacterName, "is the strongest with ", highestCharacterPower)
}

func sumPowerLevels(levels []int) {
	// Your code here
	var sum int = 0
	for _, value := range levels {
		sum += value
	}
}

func printMultiples(n uint, max uint) {
	// Your code here
	for i := n; i <= max; i += n {
		fmt.Println(i)
	}
}

func countDown(n int) {
	// Your code here
	for n > 0 {
		fmt.Println(n)
		n--
	}
}

func countUp(num int) {
	// Add your code here
	var counter int = 0
	for {
		fmt.Println(counter)
		counter++
		if counter == num {
			break
		}

	}
}

func practicing() {
	var nums [5]int
	nums[0] = 1
	fmt.Println(nums)
}

func mapExample() {
	foo1 := make(map[string]int)
	foo1["Hello"] = 10
	fmt.Print(foo1)
	return

	type NestedData map[string]map[string]map[string]string
	// 1. Data Initialization:
	// In JS, this is just a nested object literal.
	// In Go, we must explicitly declare the full nested type: map[string]map[string]map[string]string
	ages := NestedData{
		"suraj": {
			"foo": {
				"foo": "foo",
			},
		},
	}

	// 2. The "Container":
	// We declare 'foo' with the exact same structure as 'ages'.
	// If the structure didn't match, Unmarshal would fail or skip fields.
	var foo NestedData

	// 3. Marshaling:
	// Transform the Go map into a JSON byte slice ([]byte).
	// JS Equivalent: const jsonData = JSON.stringify(ages);
	jsonData, err := json.Marshal(ages)

	// In Go, we check errors immediately (idiomatic "guard clause" pattern)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// 4. Unmarshaling:
	// We pass 'jsonData' (the source) and '&foo' (the destination address).
	// '&' is essential because Unmarshal needs to modify the 'foo' variable directly.
	// JS Equivalent: foo = JSON.parse(jsonData);
	err = json.Unmarshal(jsonData, &foo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 5. Output:
	// 'jsonData' is []byte, so we must cast it to 'string' to see the JSON text.
	fmt.Println("JsonData: ", string(jsonData))

	// 'foo' is now populated with the data from the JSON
	fmt.Println("foo: ", foo)
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
