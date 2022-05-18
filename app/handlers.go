package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World !")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	Customers := []Customer{
		{Name: "Mohan", City: "Karimnagar", Zipcode: "505001"},
		{Name: "Shravya", City: "Dharmaram", Zipcode: "505470"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(Customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customers)

	}

}

