package main

import (
	"encoding/json"
	"io/ioutil"
)

type Salary struct {
	Basic  int
	Medical int
	HRA     int
	TA      int
	DA      int
}

type Employee struct {
	Name    string
	Age     int
	City    string
	Salary  []Salary
}
func main(){
	data := Employee{
	Name:    "John",
	Age:     30,
	City:    "New York",
	Salary: []Salary{
		{
			Basic:  10000,
			Medical: 5000,
			HRA:     2000,
			TA:      1000,
			DA:      5000,
		},
		{
			Basic:  15000,
			Medical: 6000,
			HRA:     2500,
			TA:      1500,
			DA:      6000,
		},
	},
}

 file,_ :=json.MarshalIndent(data,""," ")
_ = ioutil.WriteFile("test.json", file, 0644)
}
