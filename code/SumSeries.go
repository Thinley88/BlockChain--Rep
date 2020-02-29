package main

import "fmt"

func main(){
	var  i, sum int
	sum = 0
	for i=1; i<=10; {
		sum= sum + i/(2*i)
		fmt.Print("\nThe sum of number in 10 series is = ", sum)
		i++;
	}
}
