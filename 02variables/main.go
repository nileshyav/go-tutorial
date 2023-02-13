package main

import "fmt"

var CupKake string = "anything" // first Capital letter used for making it global

const ApiKey string = "sdfsdfdsfdsfsfsf"

func main() {
	// finding type of a value  :-- start
	var userName bool = true
	fmt.Printf("Variable type is : %T \n", userName)

	var intValues int = 1212121
	fmt.Printf("Variable type is : %T \n", intValues)

	var floatValue float64 = 12.0212345678987654321
	fmt.Println(floatValue)
	fmt.Printf("Variable type is : %T \n", floatValue)

	var stringVal string
	fmt.Println(stringVal)
	fmt.Printf("Variable type is : %T \n", stringVal)

	var emptyVal int
	fmt.Println("----------\n", emptyVal)
	fmt.Printf("Variable type is : %T \n", emptyVal)
	// finding type of a value  :-- end

	//  another way to declear a variable

	var courses = "Devops Pro"
	fmt.Println(courses)

	//  another way to declear a variable
	// := walrus allowed inside function but not in global
	anotherCourses := "KodeKloud courses"
	fmt.Println(anotherCourses)

	fmt.Println(CupKake)
	fmt.Println(ApiKey)
}
