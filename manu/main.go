package main

import (
	"fmt"
)

type car interface {
	tata()
	mahindra()
}

type spec struct {
	company, drive, fuel, model string
	price                       int 
}

func (s *spec) tata() spec {
	fmt.Println("TATA NEXA Specifications")
	return spec{}
}

func (s *spec) mahindra() spec {
	fmt.Println("MAHINDRA XUV Specifications")
	return spec{}
}
func CarType() spec {
	car1 := spec{
		company: "TATA",
		model:   "NEXA",
		drive:   "Four-wheel Drive",
		fuel:    "EV",
		price:   750000,
	}
	fmt.Printf("%v", car1.tata())

	car2 := spec{
		company: "MAHINDRA",
		model:   "XUV",
		drive:   "Four-wheel Drive",
		fuel:    "EV",  
		price:   1500000,
	}
	fmt.Printf("%v", car2.mahindra())

return spec{}
}
