package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type book struct{
	Name string `json:"name"`
	Author string `json:"author"`
	OriginalLanguage string `json:"originalLanguage"`
	ApproximateSales float32 `json:"approximateSales"`
	FirstPublished uint `json:"firstPublished"`
}

type Books struct{
	Books []book `json:"books"`
}


func main() {

	// open the books json file
	jsonFile, err := os.Open("books.json")
	if err != nil {
    fmt.Println(err)
	}
	// close json file when the function ends
	defer jsonFile.Close()

	// read opened json file as byte array to use in further unmarshal method
	bytevalue, _ := ioutil.ReadAll(jsonFile)
	var books Books
	json.Unmarshal(bytevalue, &books)

		
	// check if an arguement is given by the user or not
		if len(os.Args) > 1 {
			//if an arguement is given, execute respective function
		if os.Args[1] == "list" {	
			list(books)
		} else if os.Args[1] == "search"{
			bookName := strings.Join(os.Args[2:], " ")
			search(books, bookName)
		} else {
			err := errors.New("This arguement does not exist. Please try again.")
			fmt.Println(err)
		} 
		} else {
			err := errors.New("Please type one of the 'list' or 'search <bookname>' arguements")
			fmt.Println(err)
		}
}

// list : takes input Books struct (which contains an array consisting of book structs) prints the names of all the books in the book list
func list(books Books){
	for _, v := range books.Books {
		fmt.Println(v.Name)
	}
}

// search : searches books by the name input given by the user and prints the name of the book if it is included in the book list
// the function is case insensitive
func search(books Books, bookName string){
	isListed := false
		for _, v := range books.Books {
			if strings.ToLower(v.Name) == strings.ToLower(bookName){
			isListed = true
			fmt.Println(v.Name)
			}
		}
		if (isListed == false) {
			err := errors.New("This book is not on the list.")
			fmt.Println(err)
		}
}