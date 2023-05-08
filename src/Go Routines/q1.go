package main

import "fmt"

func main() {
	ch := make(chan int)
	go product(2, 3, ch)
	go product(4, 5, ch)

	result1 := <-ch
	result2 := <-ch

	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)

	fmt.Println("Product :", result1*result2)
}

func product(a int, b int, ch chan int) {
	product := a * b
	ch <- product
}