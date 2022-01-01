package main

import (
	"fmt"
)

// Functions
// A function is a block of code that performs a specific task
// A function takes an input, performs some calculations on the input, and generates an output.

// Syntax =>
/*
			func functionname(parametername type) returntype {
	 			//function body
			}

	OR

			func functionname(){

			}

*/

// 1st Function Way
func calculateBill(price int, no int) int { // also can be  => (print,no int) if same Type
	totalPrice := price * no
	return totalPrice
}

// 2nd Function Way
func rectProp(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter // remember (float64,float64)  2 times so return is 2 times
}

// 3rd Function Way
func maxNum(num1, num2 int) (result int) {
	if num1 > num2 {
		return num1 // if result not declared above then it will be -> return result
	} else {
		return num2
	}

}

// 4th Function
func swap(a, b string) (string, string) {
	return b, a
}

// 5th sums Function - Variadic Function
func sums(numm ...int) int {
	res := 0
	for _, n := range numm {
		res += n
	}
	return res
}

func main() {

	fmt.Println("2nd Part of Basics Golang Started - Only about Functions")

	// calculateBill Function =>
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is => ", totalPrice)

	// rectProp Function
	area, perimeter := rectProp(14.5, 8.4)
	fmt.Printf("Area : %f\nPerimeter: %f\n", area, perimeter)

	// maxNum Function
	result := maxNum(92, 53)
	fmt.Printf("Maximum Number :=> %d", result)

	// swap Function
	x, y := swap("Aakash", "Choudhary")
	fmt.Println("\nSwap Function Applies => ", x, y)

	// sum Function

	// 5th Function => Anonymous
	adds := func(a, b, c int) int {
		return a + b + c
	}(4, 3, 9)

	fmt.Println(adds)

	// 6th variadic Function =>
	// A variadic function can accept variable number of parameters
	s1 := sums(1, 2, 3)
	s2 := sums(3, 1, 4, 5, 1)
	s3 := sums(4, 1, 5, 77, 3, 6, 2, 1, 5, 6)
	fmt.Println(s1, s2, s3)
	// See we have many parameters

}
