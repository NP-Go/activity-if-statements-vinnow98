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
	fmt.Println("")
	activity2()
}
func activity2() {
	number1 := 17
	resultMessage := "No outcome."
	if number1%2 == 0 {
		fmt.Print("This is an even number")
	} else if number1%2 == 1 {
		fmt.Println("This is an odd number")
	} else {
		fmt.Println(resultMessage)

	}
	if number1 > 9 {
		fmt.Println("This number has more than 1 digit")
	} else if number1 < -9 {
		fmt.Println("This number has more than 1 digit")
	} else {
		fmt.Println("This number has only 1 digit")
	}
}

func activity3() {

}
