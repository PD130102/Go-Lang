package main

import (
  "fmt"
  "math/rand"
)
/*
    rand.Int() which calls the function to return a random integer.
    rand.Intn() which calls the function to return a random element from 0 up to the specified number provided.

*/
func main() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%d) %d\n", i, rand.Intn(25))
  }
  fmt.Println("**********************************")
  for i:=0;i<10;i++{
	fmt.Printf("%d) %d\n",i,rand.Int())
  }
}
