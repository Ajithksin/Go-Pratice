package main

import (
	"log"
)

type Database interface{
	GetUser() string
	GetAllUsers() []string
}

type Defaultdatabase struct {}

func (db Defaultdatabase) GetUser() string{
	return "Rohini"
}

type FamousDatabase struct {}

func (db FamousDatabase) GetUser() string{
	return "Ajit"
}

func (db FamousDatabase) GetAllUsers() []string {
	return [] string {}
}

type Greeter interface {
	Greet (username string)
}

type NiceGreeter struct {}

func (ng NiceGreeter) Greet(username string){
	log.Printf(format:"Hello %s!! nice to meet you", username)
}

type Program struct{
	Db Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main(){
	db := FamousDatabase
	Greeter := NiceGreeter{}
	p := Program
	Db : db,
	Greeter : greeter,
	p.Execute()
}
