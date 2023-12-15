package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Ardi",
		MiddleName: "",
		LastName:   "Hidayat",
		Hobbies:    []string{"Gaming, Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Ardi","MiddleName":"","LastName":"Hidayat","Age":0,"Married":false,"Hobbies":["Gaming, Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Ardi",
		Addresses: []Address{
			{
				Street:     "Jl. Raya",
				Country:    "IND",
				PostalCode: "13900",
			},
			{
				Street:     "Jl. Gang K",
				Country:    "IND",
				PostalCode: "18900",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Ardi","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl. Raya","Country":"IND","PostalCode":"13900"},{"Street":"Jl. Gang K","Country":"IND","PostalCode":"18900"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl. Raya","Country":"IND","PostalCode":"13900"},{"Street":"Jl. Gang K","Country":"IND","PostalCode":"18900"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)

}
