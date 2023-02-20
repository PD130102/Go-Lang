package main

import (
	"encoding/csv"
	"log"
	"os"
)


func main(){
  data := [][]string{
	{"Name", "Age", "City"},
	{"John", "30", "New York"},
	{"Jenny", "20", "New York"},
	{"Sam", "35", "London"},
	{"Phil", "25", "Paris"},
}
 csvfile,err := os.Create("sample.csv")
 if err != nil {
	 log.Fatalf("Error creating csv file: %s", err)
 }

 defer csvfile.Close()

 writer := csv.NewWriter(csvfile)
 for _, value := range data {
	 _ = writer.Write(value)
 }
 
writer.Flush()
csvfile.Close()
}