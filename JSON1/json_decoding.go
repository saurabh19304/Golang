package main

//golang program to illustrate the
// concept of the decoding using JSON

import (
	"encoding/json"
	"fmt"
)

// declaring a struct
type Human struct {

	//sefining struct varaible
	Name    string
	Address string
	Age     int
}

// main function is written here also it can be one in a single file
func main() {
	//defining a struct instance
	var human1 Human

	//data in JSON format which is to be decoded
	Data := []byte(`{
	"Name":"deeksha",
	 "Address:":"hydrabad",
	  "Age": 21
	  }`)

	//decoding human1 struct
	// from json format
	err := json.Unmarshal(Data, &human1)

	if err != nil {
		//if error is not nil
		// print error
		fmt.Println(err)
	}

	//printing the details of decoding data
	fmt.Println("Struct is:", human1)
	fmt.Printf("%s lives in %s. \n", human1.Name, human1.Address)

	//unmarshing a json array
	// to array type in Golang

	//defining an array instance
	// of struct type
	var human2 []Human

	//JSON array decoded to an array main

	Data2 := []byte(`[
	  	{"Name": "Vani", "Address": "Delhi", "Age": 21},
	  	{"Name": "rishi", "Address": "noida", "Age": 43},
	  	{"Name":"rohit", "Address": "pune", "Age": 56}
	  ]`)

	//decoding JSON arary to human array
	err2 := json.Unmarshal(Data2, &human2)

	if err2 != nil {

		//if error is not nil
		// print error
		fmt.Println(err2)
	}

	//printing the decoded array
	// value one by one
	for i := range human2 {
		fmt.Println(human2[i])
	}
}
