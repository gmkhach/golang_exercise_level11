package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
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
		Create a custom error message using "fmt.Errorf" starting with the following code:

		main() {
			p1 := person{
				First:   "James",
				Last:    "Bond",
				Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
			}

			bs, err := toJSON(p1)

			fmt.Println(string(bs))
		}

		// toJSON needs to retunr an error also
		func toJSON(a interface{}) []byte {
			bs, _ := json.Marshal(a)
		}
	*/
	p1 = person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err = toJSON(p1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))

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

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Error: %v", err)
	}
	return bs, nil
}
