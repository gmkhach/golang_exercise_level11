package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First		string
	Last 		string
	Sayings		[]string
}

func main() {
    /*
		Exercise 1
		Instead of using the blank identifier, make sure the code is checking and handling the error in the following block of code:
			p1 := person{
				First:   "James",
				Last:    "Bond",
				Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
			}

			bs, _ := json.Marshal(p1)
			fmt.Println(string(bs))
    */
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	fmt.Println(string(bs))
	
	
	/*
		Exercise 2

    */
	
	
	
	/*
		Exercise 3

    */
	
	
	
	/*
		Exercise 4

    */
	
	
	
	/*
		Exercise 5

    */
}