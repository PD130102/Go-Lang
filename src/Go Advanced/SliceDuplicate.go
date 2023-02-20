package main

import ("fmt")

func main(){
	 slice := []int{1,2,3,4,5,6,7,6,5,4,4,2,4,5,7,8,7,9,10,2,3,8,9,10}
	 fmt.Println(slice)
	 fmt.Println("After removing duplicate:(DuplicateRemoval)")
	 fmt.Println(removeDuplicates(slice))
	 fmt.Println("After removing duplicate:(Unique)")
	 fmt.Println(unique(slice))
}

func unique(slice []int) []int{
	var uniqueSlice []int
		var count int
		for i := 0; i < len(slice); i++ {
			for j := i + 1; j < len(slice); j++ {
				if slice[i] == slice[j] {
					count++
				}
			}
			if count == 0 {
				uniqueSlice = append(uniqueSlice, slice[i])
			}
			count = 0
		}
	return uniqueSlice
}


func removeDuplicates(slice []int) []int {
	m := make(map[int]bool)
	var result []int
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}