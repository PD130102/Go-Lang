package main

import (
	"fmt"
)

func main(){

	var n int 
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	fmt.Println("The number is",n)
    fmt.Println("The Sum of the Digits of the Number is",SumofDigits(n))
}
func SumofDigits(n int) int{
	var sum int
	for n>0{
		sum+=n%10
		n=n/10
	}
	return sum
}
