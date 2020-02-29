package main

import "fmt"

func main(){
	var num1, num2 int
	fmt.Println("Enter First Number")
	fmt.Scan(&num1)

	fmt.Println("Enter Second Number")
	fmt.Scan(&num2)
	if num1>num2{
		fmt.Print(num1, "  is greater than ", num2)
	}else{
		fmt.Print(num2, "  is greater than ", num1)
	}

}
