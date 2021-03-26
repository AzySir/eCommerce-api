// Followed: https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://www.sohamkamani.com/golang/parsing-json/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"ecommerce-backend/product"
	"reflect" //only needed for debugging purposes
)

const (
    serverPort = "10000"
)

type Product struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Brand              string   `json:"brand"`
	Colour             string   `json:"colour"`
	Price              string   `json:"price"`
	Image              string   `json:"image"`
	Sizes              []string `json:"sizes"`
	Special            bool     `json:"special"`
	SpecialPrice       int      `json:"special_price"`
	SpecialDescription string   `json:"special_description"`
	Category           string   `json:"category"`
	Length             string   `json:"length"`
	Width              string   `json:"width"`
}

type Error struct {
	RequestCode int `json:"request_code"`
	ErrorCode int `json:"error_code"`
	ErrorName string `json:"name"`
	ErrorDescription string `json:"description"`	
}

func myProduct(w http.ResponseWriter, r *http.Request) {
	p := product.GetProduct(12312312)
	json.NewEncoder(w).Encode(p)
	log.Println(reflect.TypeOf(p))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/products", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/products", getProducts)
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/backendproduct", myProduct)
	r.Use(mux.CORSMethodMiddleware(r))	
	log.Fatal(http.ListenAndServe(":" + serverPort, r))

}

func getProduct(w http.ResponseWriter, r *http.Request) {
	var found bool = false;
	var Products []Product;
	vars := mux.Vars(r)
	fmt.Printf("Searching Id is " + vars["id"]);	
	
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-Type", "application/json")

	Products = getAllProducts()

	for i := 0; i < len(Products); i++{
		fmt.Println(Products[i].Id)
		if string(Products[i].Id) == vars["id"] {
			var p []Product
			p = append(p, Products[i])
			// json.NewEncoder(w).Encode(Products[i])
			json.NewEncoder(w).Encode(p)
			found = true
			break
		}
	}
	
	if !found {
		// fmt.Printf("Failed not found! 404 Error!")
		json.NewEncoder(w).Encode(Error{
			RequestCode: 200,
			ErrorCode: 404,
			ErrorName: "Product Not Found!",
			ErrorDescription: "Product Code " + vars["id"] + "has not been found or does not exist- please try again",
		})
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	Products := product.GetAllProducts()
	json.NewEncoder(w).Encode(Products)
}

























// func getAllProducts() []Product {
	
// 	var Products []Product

// 	// //Append Product 1
// 	// Products = append(Products,
// 	// 	Product{
// 	// 		Id:     "111",
// 	// 		Name:   "Shoes",
// 	// 		Colour: "red",
// 	// 		Price:  "100",
// 	// 		Image:  "https://images.asos-media.com/products/adidas-originals-ozweego-trainers-in-grey-multi/22322025-4?$XXL$&wid=513&fit=constrain",
// 	// 		Sizes: []string{
// 	// 			"s",
// 	// 			"m",
// 	// 			"l",
// 	// 		},
// 	// 		Special:            false,
// 	// 		SpecialPrice:       0,
// 	// 		SpecialDescription: "None",
// 	// 		Category:           "Tops, Jumpers, Warm, Winter",
// 	// 	})

// 	// 	//Append Product 2
// 	// Products = append(Products, 
// 	// 	Product{
// 	// 		Id:     "23232",
// 	// 		Name:   "TShirt",
// 	// 		Brand:  "Adidas",
// 	// 		Colour: "blue",
// 	// 		Price:  "50",
// 	// 		Image:  "https://images.asos-media.com/products/under-armour-sportstyle-logo-t-shirt-in-burgundy/21992029-1-burgundy?$XXL$&wid=513&fit=constrain",
// 	// 		Sizes: []string{
// 	// 			"S",
// 	// 			"M",
// 	// 			"L",
// 	// 		},
// 	// 		Special:            false,
// 	// 		SpecialPrice:       0,
// 	// 		SpecialDescription: "None",
// 	// 		Category:           "Tops, Jumpers, Warm, Winter",
// 	// 		Length:             "10cm",
// 	// 		Width:              "20cm",
// 	// 	})

// 		// fmt.Printf(reflect.TypeOf(Products))
// 		// println("----")
// 		// fmt.Printf(reflect.ValueOf(Products).Kind())
// 		// println("----")
// 		// fmt.Printf("%T\n", Products[0])
// 		// fmt.Println(Products[0])

// 		return Products
// }

