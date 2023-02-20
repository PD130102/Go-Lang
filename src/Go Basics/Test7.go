package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	var s string
	fmt.Println("Enter the String")
	scanner.Scan()
	s = scanner.Text()
	fmt.Println("The String is:", s)
}