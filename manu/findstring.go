package main

import (
	"fmt"
)

func Find(name string, names ...string) {
	fmt.Printf("type of names is %T\n", names)
	found := false
	for i, v := range names {
		if v == name {
			fmt.Println(name, "found at index", i, "in", names)
			found = true
		}
	}
	if !found {
		fmt.Println(name, "not found in ", names)
	}
	fmt.Printf("\n")
}
func letsFind() {
	Find("john", "jacob", "will", "john", "smith")
	Find("bheem", "ram", "shyam", "bheem")
	Find("manu", "ajit", "pratim", "manu", "rohini")
	Find("jack", "jill")
}
