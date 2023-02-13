package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello")
	// takes imput from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Pizza rating : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your Pizza rating is : ", input)
	one, _ := strconv.Atoi(input)
	fmt.Printf("Type of input is : %T \n", one)
}
