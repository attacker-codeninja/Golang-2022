// Every Go program starts with a package declaration which provides a way for use to reuse code

package main

// import allows us to use code from other packages

import (
	"fmt" // The format package provides formatting for input and output
)

// Comments  Types ->
// 1.  Single Comment -> //
// 2. Multi-line Comment -> /**/

func main() { // Functions start with func and surround the code with { }
	// main is the function that is executed when you execute your program

	fmt.Println("Hello World") // Println is a function in the fmt package that outputs a string, which
	// is surrounded by double quotes and a newline to the screen

	// Want HELP about specific function or keyword ? =>
	// Try -> go doc <keyword> => go doc fmt => By this we can get decription of particular function or keyword

	// How to execute GO Programs ? => go run program-name.go
	// How to Build GO Programs ?   => go build program-name.go => It will make binary file of program name => so run this binary file as => ./program-name.go

	// VARIABLES =>

	// Variable => statically typed => means their type can't change
	// Variable => Variable names must start with a letter and may contain letters, numbers or the _

	// Assign variables =>

	// Syntax => var <variable-name> <type-name> = "Strings"
	var a string = "I am String1" // 1st Type by using "="
	fmt.Println(a)

	// Syntax => <variable-name> := "Strings"
	b := "I am String2" // 2nd Type by using shot-hand ":="
	fmt.Println(b)
	// By omitting <type-name> => GO will understand it by checking value => i.e  no need to define the data type
	// i.e value is surrounded in ""  => its mean STRING

	// INT or FLOAT =>

	// 1. INT => positive or negative number without decimals
	// Versions
	// uint8 : unsigned  8-bit integers (0 to 255)
	// uint16 : unsigned 16-bit integers (0 to 65535)
	// uint32 : unsigned 32-bit integers (0 to 4294967295)
	// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
	// int8 : signed  8-bit integers (-128 to 127)
	// int16 : signed 16-bit integers (-32768 to 32767)
	// int32 : signed 32-bit integers (-2147483648 to 2147483647)
	// int64 : signed 64-bit integers (-9223372036854775808 to
	// 9223372036854775807)

	// Example =>
	var age int = 30
	fmt.Println("My age is => ", age) // See, We can use both String & Variable in Print function

	// 2. FLOAT => a number with decimals
	// Versions : float32, float64

	// EXAMPLE =>
	var percentage float32 = 70.89
	fmt.Println("My Percentage is => ", percentage)

	// ARITHMETIC

	// We can perform arithmetic in Println (Note that floats aren't accurate)

	num1 := 93.87
	num2 := 39.32
	fmt.Println(num1 - num2)

	// Arithmetic Operators : +, -, *, /, %
	fmt.Println("6 + 4 =", 6+4)
	fmt.Println("6 - 4 =", 6-4)
	fmt.Println("6 * 4 =", 6*4)
	fmt.Println("6 / 4 =", 6/4)
	fmt.Println("6 % 4 =", 6%4)

	// CONSTANT

	// A constant is a variable with a value that can't be changed
	const pi float64 = 3.14159265359
	// Though "pi" we used in small letters but better if we use const variable as CAPITAL
	// Specially First Letter

	// We can declare multiple variables like this
	const (
		varA = 2
		varB = 3
	)

	// CONSTANT - IOTA

	// The iota keyword represents successive integer constants 0, 1, 2,…
	// It resets to 0 whenever the word const appears in the source code,
	// and increments after each const specification.

	const (
		one = iota
		two
		three
		// one -> 0, two -> 1, three 2

		/*

			Extra notes on iota =>

			We can either use iota as
				name1 = iota
				name2 = iota
				name3 = iota
			OR
				name1 = iota
				name2
				name3

			Also => If we want to start from 1 instead of 0 =>
				name 1 = iota + 1 => now it will become 1 and so on
			Also => We can Skip value by using _

			Example =>
			C1 = iota + 1
			_
			C3
			C4

			above will print => 1,3,4
		*/
	)
	fmt.Println("one will be => ", one)
	fmt.Println("two will be => ", two)
	fmt.Println("three will be => ", three)

	// STRINGS

	// Strings are a series of characters between " or '

	var myName string = "Aakash"
	fmt.Println("My name is : ", myName)

	// Get Lenght of String -> len()
	fmt.Println("Length of string : ", len(myName)) // Aakash => 6

	// Variable and String can be combine using + also
	robo := "CyperZ"
	misson := "Mars Mission"
	fmt.Println(robo + "\n is on " + misson)

	// \n => for newline

	// Printf => used for format printing
	// %f => floats
	// %v => value
	// %T => Type [string,int,float etc]
	// %t => Booleans [true,false]
	// %d => Integers
	// %b => Binary
	// %c prints the character associated with a keycode
	// %x prints in hexcode
	// %e prints in scientific notation

	// Let's see Examples =>

	//const Pi = 3.14159265359
	//var boole = true
	//integer := 72

	// fmt.Printf("Float Value %f =>", Pi)
	// fmt.Printf("\nFloat Value with decimal precision %.3f => ", Pi)
	// fmt.Printf("\nData Type => %T", Pi)
	// fmt.Printf("\nValue  => %v", Pi)
	// fmt.Printf("\nPrint Boolean => %t", boole)
	// fmt.Printf("\nInteger => %d", integer)
	// fmt.Printf("\nBinary => %b", 10)
	// fmt.Printf("\nHexadecimal => %x", 17)
	// fmt.Printf("\nCharacter => %c", 44)
	// fmt.Printf("\nScientific Notation => %e", Pi)

	// Logical Operators

	fmt.Println("true && false =", true && false) // && => and
	fmt.Println("true || false =", true || false) // || => or
	fmt.Println("!true =", !true)                 // ! => not

	// FOR Loop => 2 Ways =>

	// One Way =>

	// i := 1

	// for i <= 10 { // 1 to 10
	// 	println(i)
	// 	i++
	// }

	// // 2nd Way
	// for j := 15; j <= 20; j++ {
	// 	println(j)
	// }

	// Relational Operators include ==, !=, <, >, <=, and >=

	// CONDITIONS => if, if else, else if else, Switch

	// IF  ELSE Condition

	// myAge := 17

	// if myAge >= 18 {
	// 	println("You can vote")
	// } else {
	// 	println("You can't vote")
	// }

	// IF ELSE IF ELSE Condition

	// fruit := "Grapes"

	// if fruit == "Mango" {
	// 	println(fruit + " is Superb")
	// } else if fruit == "Grapes" {
	// 	println(fruit + " => Easily can be eaten")
	// } else {
	// 	println("Awesome")
	// }

	// Switch statements are used when you have limited options

	// enter := "Yes"

	// switch enter {
	// case "Yes":
	// 	fmt.Println("You entered : Yes")
	// case "No":
	// 	fmt.Println("You entered : No")
	// default:
	// 	fmt.Println("Please enter Yes or No ")
	// }

	// ARRAYS & SLICES

	// 1. ARRAYS

	// An Array holds a fixed number of values of the same type

	// Syntax Defining 1 =>
	// Using var
	// var array_name = [length]datatype{values} // here length is defined
	//			or
	// var array_name = [...]datatype{values} // here length is inferred
	// Using :=
	// 	array_name := [length]datatype{values} // here length is defined
	// 			or
	//  array_name := [...]datatype{values} // here length is inferred

	// Syntax Defining 2 =>
	// var array_name = [length]datatype
	// array_name[0] = value
	// array_name[1] = value2

	// Like above we can define Arrays

	// Examples =>
	var favNum = [5]float32{23.4, 12.4, 43.0, 53.4, 24.1}
	fmt.Println(favNum)
	favNum2 := [5]int{21, 42, 43, 56, 77}
	fmt.Println(favNum2)
	var favName3 [3]string
	favName3[0] = "Aakash"
	favName3[1] = "Vikas"
	favName3[2] = "Maerifat"
	fmt.Println(favName3[2]) // Accessing particular Array Index

	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // here we didn't define size

	for n, m := range nums {
		fmt.Println(n, m)
	}

	// Iterate Array
	for k, value := range favName3 {
		fmt.Println(k, value)
	}

	// SLICES
	// Slices are like arrays but you leave out the size
	numSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	// Remember from above =>
	/*
		A = 0
		B = 2
		C = 2 and so on
	*/

	// We can create a slice by defining the first index value to take through the last
	numSlice2 := numSlice[3:5] // D,E
	fmt.Println("numSlice2[0] => ", numSlice2[0])
	fmt.Println("numSlice2[1] => ", numSlice2[1])
	fmt.Println("numSlice[5:] => ", numSlice[5:])
	fmt.Println("numSlice[:5] => ", numSlice[:5])
	fmt.Println("numSlice[:]  => ", numSlice[:])        // All
	fmt.Println("Length of numSlice =>", len(numSlice)) // 9

	// Create slice using make()
	numSlice3 := make([]string, 3, 5) // Empty Slice
	// Note => 3 => length & 5 = capacity [max capacity]
	// Now use copy function
	copy(numSlice3, numSlice) // copy numSlice into numSlice3
	fmt.Println(numSlice3[0])

	// Append values to the end of a slice

	abc := []int{1, 2}
	abc = append(abc, 0, -1)
	fmt.Println(abc)

	numSlice = append(numSlice, "hi", "bye")
	fmt.Println(numSlice)
	fmt.Println(numSlice[9])

	numSlice3 = append(numSlice3, "jake", "cake")
	fmt.Println(numSlice3)
	fmt.Println(numSlice3[4])

	// MAPS
	// A Map is a collection of key value pairs

	/* Syntax1 => varName := make(map[keyType] valueType)

	varName[0] = "Yes1"
	varName[1] = "Yes2"
	*/

	// Syntax 2 => varName := make(map[keyType]valueType{Key:Value,Key2:Value2})

	// Initialize an empty map
	// var varName = map[string]int{}

	// Adding items (key-value pairs) to a map => varName[key] = value
	// Examples =>

	details := make(map[string]int)
	details["Aakash"] = 30
	details["Vikas"] = 23
	details["Maerifat"] = 27
	fmt.Println("Age of Aakash => ", details["Aakash"])
	fmt.Println("Number of elements using len => ", len(details)) //3
	// Add new Element
	details["Kashish"] = 31
	fmt.Println("Length Increased now => ", len(details)) //4

	// Delete element =>
	delete(details, "Kashish")
	fmt.Println(details)

	// Checking key in Maps

	// Syntax => value, ok := map[key] => whether a particular key is present in a map
	// If ok is true, then the key is present and its value is present in the variable value,
	// else the key is absent.

	value, ok := details["Aakash"] // details is map & "Aakash" is Key

	if ok == true {
		fmt.Println(value, " is Present")
	} else {
		fmt.Println("Not Found")
	}

	/*

		If just want to check for the existence of a key without retrieving the value
		associated with that key, then we can use an _ (underscore)
		in place of the first value -

		_, ok := map[key]

	*/

	// Iterate Maps

	for key, value := range details {
		fmt.Printf("Key[%s] = Value[%d]\n", key, value)
	}

	// NOTE =>
	/*
		Maps are reference types.

		When we assign a map to a new variable, they both refer to the same underlying data
		structure.

		Therefore changes done by one variable will be visible to the other
	*/

	// Maps in Maps => Nested Maps =>
	// Syntax => var x = map[string]map[string]string{}

	// Example =>

	// 1st Way using make function
	xyz := make(map[string]map[string]string)

	xyz["fruits"] = make(map[string]string)
	xyz["colors"] = make(map[string]string)

	xyz["fruits"]["a"] = "Apple"
	xyz["fruits"]["g"] = "Grapes"

	xyz["colors"]["r"] = "Red"
	xyz["colors"]["g"] = "Green"

	fmt.Println(xyz["colors"])

	// Similary we can create nested maps without make Function
	// Let's see how

	superHeros := map[string]map[string]string{
		"Superman": map[string]string{
			"Species": "Alien",
			"Power":   "Super Lasik",
		},
		"Spiderman": map[string]string{
			"Species": "Human-Spider",
			"Power":   "Spider Abilities",
		},
	}

	// Now let's access it =>
	if key, value := superHeros["Superman"]; value {
		fmt.Println(key["Species"], key["Power"])
	}

}
