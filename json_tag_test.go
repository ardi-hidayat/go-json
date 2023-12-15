package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Mac Pro",
		ImageURL: "macbook.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))

}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"Id":"P001","name":"Apple Mac Pro","image_url":"macbook.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
