package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type Customer struct{
	Id string 			`json:"id"`
	FirstName string`json:"first_name"`
	LastName string	`json:"last_name"`
	Cash int		`json:"cash"`
	Basket Basket	`json:"basket"`
}


 type Basket struct{
	Id string			  `json:"id"`
	Products []Product`json:"products"`
	Total int		  `json:"total"`
}


type Product struct {
	Id string		`json:"id"`
	Category string	`json:"category"`
	Name string		`json:"name"`
	Price int		`json:"price"`
	Quantity int	`json:"quantity"`
}
var(
	Customers = []Customer{}
)

func OpenReadFile() {
	filejson, err := os.OpenFile("first_exam.json", os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error while opening json file!", err)
		return
	}

	err = json.NewDecoder(filejson).Decode(&Customers)
	if err != nil{
		fmt.Println("Error while decoding Customers!", err)
		return
	}
	
}
