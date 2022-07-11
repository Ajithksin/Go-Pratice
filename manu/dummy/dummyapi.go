package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type customer_details struct {
	Name       string `json:"fullname`
	Acc_number string `json:"Acc_number"`
	City       string `json:"city"`
}

func main() {
	http.HandleFunc("/account", accounts)
	http.HandleFunc("/CustomerList", getCustomerList)

	http.ListenAndServe(":8080", nil)

	// http.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Account list")
	// })
}

func accounts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "accounts list here\n")
}

func getCustomerList(w http.ResponseWriter, r *http.Request) {
	customer := []customer_details{
		{Name: "Manu", Acc_number: 321456, City: "Dharwad"},
		{Name: "mike", Acc_number: 456123, City: "hubli"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
	//fmt.Fprint(w, json.NewEncoder(w).Encode(customer))
}
