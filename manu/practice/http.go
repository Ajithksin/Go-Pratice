// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/Home", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Welcome to Golang")

// 		http.HandleFunc("/Gadgets", func(w http.ResponseWriter, r *http.Request) {
// 			fmt.Fprintf(w, "Mobiles, cameras, laptops")
// 		})
// 		http.ListenAndServe(":5050", nil)
// 	})
// }

package main

import (
	"fmt"
	"net/http"
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

	http.ListenAndServe(":5050", nil)
}
