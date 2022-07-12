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
		fmt.Fprintf(w, "Golang history\n")
	})

	http.HandleFunc("/introduction", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get set GO!!!!")
	})

	http.HandleFunc("/time&date", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Today is:", time.Now())
	})
	http.HandleFunc("/developerinfo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Developer Name: Manu Vijaykumar Hali\n Email: manu.hali3d@gmail.com\n")
	})

	http.ListenAndServe(":5050", nil)
}

func logger(f func(http.ResponseWriter, *http.Request)) {

}
