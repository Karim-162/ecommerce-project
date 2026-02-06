package handlers

import (
	"ecommerceProject47/database"
	"ecommerceProject47/utility"
	"net/http"
	"strconv"
)

func Getproductbyid(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "please give a valid ID", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			utility.SendData(w, product, 200)
			return
		}
	}
	utility.SendData(w, "data pai nai", 404)
}
