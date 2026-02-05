package handlers

import (
	"ecommerceProject47/database"
	"ecommerceProject47/utility"
	"net/http"
)

func Getproducts(w http.ResponseWriter, r *http.Request) {

	utility.SendData(w, database.ProductList, 200)
}
