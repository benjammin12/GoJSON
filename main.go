package main

import (
	"encoding/json"
	"log"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}


func main() {
	/* error using os.open because it needed the exact number of bytes, ioutil is better because it just returns
	 ** a slice of bytes instead of a pointer to a file


	jsonFile, err := os.Open("books.json") //returns an error and a pointer to a file
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}

	data := make([]byte,100) // if no error, we can read the file into a slice of bytes

	count, err := jsonFile.Read(data)

	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	fmt.Printf("number of byes is %d\n", count) // <-slice of bytes
	*/

	data, err := ioutil.ReadFile("books.json") //returns an error and bytes
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	fmt.Println(data) //bytes
	fmt.Println(string(data))

	//for multiple values, store the unmarshall in a list of struct types
	var b []Book
	err = json.Unmarshal(data, &b)
	if err != nil {
		log.Fatalln("Error unmarshalling file",err)
	}
	for _,val := range b {
		fmt.Println(val.Title, "written by:", val.Author)
	}

	//marshalling returns the bytes
	data, err = json.Marshal(b)
	if err != nil {
		log.Fatalln("Error marshalling")
	}
	fmt.Println(data)

}


