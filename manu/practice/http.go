package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to golang server!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Golang history")
	})

	http.HandleFunc("/introduction", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get set GO!!!!")
	})

	http.HandleFunc("/time&date", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Today is:", time.Now())
	})
	http.ListenAndServe(":5050", nil)
}


func logger(f func(http.ResponseWriter, *http.Request)){

}