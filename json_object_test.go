package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Ardi",
		MiddleName: "",
		LastName:   "Hidayat",
		Age:        27,
		Married:    false,
		Hobbies:    []string{"Travelling"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
