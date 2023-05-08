package main

import "fmt"

func main(){

	var a = []int{1,2,3}
	fmt.Println(a)
    var numbers = make([]int,3,5)
    printSlice(numbers)
}
func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}