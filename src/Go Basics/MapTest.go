package main
 
import (
	"fmt"	
)

func main() {
	studentMap := make(map[string]string)
	var n int
    fmt.Println("Enter the Number of Students")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var name string
		var age string
		fmt.Println("Enter the Name of Student")
		fmt.Scanln(&name)
		fmt.Println("Enter the Age of Student")
		fmt.Scanln(&age)
		studentMap[name] = age
	}

	for key, value := range studentMap {
		fmt.Println("Name:", key, "Age:", value)
	}
}