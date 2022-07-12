package main

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/account", accounts)
	http.HandleFunc("/CustomerList", getCustomerList)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
