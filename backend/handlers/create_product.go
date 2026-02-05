package handlers

import (
	"ecommerceProject47/database"
	"ecommerceProject47/utility"

	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "plz give me valid JSON", 400)
		return
	}

	fmt.Println("Received product:", newProduct)

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	utility.SendData(w, newProduct, 201)

}
