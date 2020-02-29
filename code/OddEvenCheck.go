package main
import "fmt"

func main(){
	var num1 int
	fmt.Print("Enter your number to check even or odd?   :")
	fmt.Scan(&num1)
	if num1%2==0 {
		fmt.Print(num1, " is Even number")
	}else {

		fmt.Print(num1,"is odd number")
	}

}
