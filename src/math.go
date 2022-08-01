package main

func main() {
	// call the Add function
	result := Add(1, 2)
	// print the result
	println(result)
}

// the Add Function takes two integers and returns their sum
func Add(number1 int, number2 int) int {
	return number1 + number2
}
