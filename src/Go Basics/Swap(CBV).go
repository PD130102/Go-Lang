package main

import "fmt"

func main(){
 var a,b int
 fmt.Println("Enter the two numbers")
 fmt.Scanln(&a,&b)
 fmt.Println("The number a before swapping is",a)
 fmt.Println("The number b before swapping is",b)

 Swap(a,b)
 
 fmt.Println("The number a after swapping is (in the main)",a)
 fmt.Println("The number b after swapping is (in the main)",b)
}
func Swap(a,b int) {
	a,b=b,a
	fmt.Println("The number a after swapping is (in the function)",a)
	fmt.Println("The number b after swapping is (in the function)",b)
}