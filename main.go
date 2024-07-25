package main

import (
	"encoding/json"
	"fmt"
)
type User struct{
	Name string  
	Age int 
	IsBlocked bool 
	Books []Book
	
}
type Book struct{
	Name string
	Year int
}


func main() {
	var books []Book
	book := Book{
	Name: "Saga of Vindland",
	Year: 2014,
}
books = append(books, book)
	sV := User{
		Name : "Ala",
		Age: 19,
		IsBlocked: true,
	}

			
		
	
boolVar, _ :=	json.Marshal(sV)
fmt.Println(string(boolVar))

}