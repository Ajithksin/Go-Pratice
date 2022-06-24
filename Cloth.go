//Rohini 
//package main
//import "encoding/json"
//myJsonString := '{"name": "Json"}'
//myStoreValue is address of variable we want to store our data 
//err := json.Unmarshal(Data, &country)
//unmarshal is provided by Go (Json standard lib)

package cloths 

import ( 

    "fmt" 

    "encoding/json" 

) 
type Cloths interface { 

    Cloths() string 

}

//declaring a struct 

type Brands struct { 

    //defining struct variables 

     Name    string 
     Size    string 
     Country string 

} 

//main function 

func cloths() { 

    //defining a struct instance 

    var brands []Brands 

    //JSON array to be decoded 
    //to an array in golang 

    Data := []bytes(' 

    [ 

        {"Name": "Allen Solly",      "Size": "Large",     "Country": "India"}, 
        {"Name": "Arrow",            "Size": "Small",     "Country": "America"}, 
        {"Name": "Emporio Armani",  "Size": "Large",     "Country": "Italy"}, 
        {"Name": "Fila ",            "Size": "Mediun",    "Country": "South Korea"}, 
        {"Name": "Jack&Jones",       "Size": "Small",     "Country": "Europe"}, 

    ]') 

    //decoding JSON array to brand array 
    err := json.Unmarshal(Data, &brand) 

    if err != nil { 

        //if error is not nil 

        //print error 

        fmt.Println(err) 

    } 

    //printing decoded array 
    //values one by one 
    for i := range brand { 

        fmt.Println(brands[i].Name + " - " + brandsbrands[i].Size + 

            " - " + brands[i].Countries) 

    } 

} 
