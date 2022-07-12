package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type customer_details struct {
	Name       string `json:"fullname"`
	Acc_number string `json:"Acc_number"`
	City       string `json:"city"`
}

func accounts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "accounts list here\n")
}

func getCustomerList(w http.ResponseWriter, r *http.Request) {
	customer := []customer_details{
		{Name: "Manu", Acc_number: "321456", City: "Dharwad"},
		{Name: "mike", Acc_number: "456123", City: "hubli"},
	}

	//fmt.Fprint(w, json.NewEncoder(w).Encode(customer))

	//marshaling to xml
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		//marshaling to json
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
