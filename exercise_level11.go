package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type customer struct {
	name	string
	cart	[]string
}

type sqrtError struct {
	lat		string
	long	string
	err1	error
	err2	error
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
		1. Create a struct “customErr” which implements the built-in error interface. 
		2. Create a func “foo” that has a value of type error as a parameter. 
		3. Create a value of type “customErr” and pass it into “foo”.
	*/
	customErr := customer {
		name: "James Bond",
		cart: []string{"Walther PPK", "silencer", "50 rounds 380 ACP"},
	}

	foo(customErr)

	/*
		Exercise 4
		1. Use the sqrt.Error struct as a value of type error. 
		2. Use these numbers for your lat and longa:
			- lat: "24.2192N"
			- long: "79.3314W"
	*/
	_, err = sqrt(-1.0)
	if err != nil {
		log.Println(err)
	}
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Error: %v", err)
	}
	return bs, nil
}

func (c customer) Error () string {
	return fmt.Sprintf("Out of stock on: %v", c.cart[1])
}

func foo(e error) {
	fmt.Println("foo ran")
	
	// By the way if we wanted to print out a value of customer from e we would first have to assert 
	// that the passed value is of type customer because here it is being passed as an error.
	fmt.Println("This will run", e.(customer).name)

	// If we tried to run it by accessing the value of name directly from e it wouldn't run:
 	// fmt.Println("This won't run", e.name)
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v %v", se.lat, se.long, se.err1, se.err2)
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		// Here we are passing two different type of error to demonstrat that with Errorf we can also calpture the value passed.
		se := sqrtError {
			lat: "24.2192N",
			long: "79.3314W",
			err1: errors.New("I'm sorry Dave, I'm afraid I can't do that."),
			err2: fmt.Errorf("The value passed was: %v", n),
		}
		return math.NaN(), se
	}
	return math.Sqrt(n), nil
}