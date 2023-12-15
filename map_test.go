package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Mac Pro","price":"2000"}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonByte, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "POO1",
		"name":  "Aplle Mac Pro",
		"price": 2000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
