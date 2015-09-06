package main

import "fmt"
import "encoding/json"

type Product struct {
	// fields to retrieve
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	var pro Product
	s := `{"name":"Galaxy Nexus", "price":3460.00, "ignored":"field"}`
	err := json.Unmarshal([]byte(s), &pro)

	if err == nil {
		fmt.Print("Product ", pro.Name, " has price ")
		fmt.Printf("$%+v\n", pro.Price)
		fmt.Println("Fields:", pro)
	} else {
		fmt.Println(err)
	}
}
