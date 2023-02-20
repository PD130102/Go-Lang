package main

import "fmt"

func main(){
  
	var n int
	fmt.Println("Enter the number of elements in the array")
	fmt.Scanln(&n)
	fmt.Println("Enter the elements of the array")
	var slice = make([]int,n)

	for i:=0;i<n;i++{
		fmt.Scan(&slice[i])
	}
	fmt.Println("The array is",slice)
	fmt.Println("The Max of the array is",Max(slice))
}
func Max(slice []int) int{
	var max int
	for i:=0;i<len(slice);i++{
		if slice[i]>max{
			max=slice[i]
		}
	}
	return max
}