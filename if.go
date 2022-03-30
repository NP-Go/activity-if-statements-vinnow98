package main

import "fmt"

func main() {
	number1 := 17
	number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string
	if number1 < number2 {
		fmt.Printf("%v is bigger than %v", number2, number1)
		difference := number2 - number1
		fmt.Printf("\nNumber2 is bigger by number1 by %v", difference)
	} else {
		fmt.Printf(resultMessage)
	}
}
