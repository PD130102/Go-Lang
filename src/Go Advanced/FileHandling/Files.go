package main

import (
	"os"
	"log"
)

func main(){
	ELife,err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Log file created")
	log.Println(ELife)
	defer ELife.Close()
}