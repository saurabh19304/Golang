package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string
	Age     int
	Address string
}

func main() {

	//we are defining the a struct instance main
	human1 := Human{"ankit", 23, "new delhi"}

	//encoding human1 struct
	// print error
	human_enc, err := json.Marshal(human1)

	if err != nil {
		//if error in not nil
		// print error
		fmt.Println(err)
	}

	//as human_enc is in a byte array
	// format, it needs to be
	// converted into the string
	fmt.Println(string(human_enc))

	//converting slices from
	// golang to JSON format
	//
	// defining an array
	// of struct instance
	human2 := []Human{
		{Name: "rahul", Age: 2, Address: "new delhi"},
		{Name: "priyansh", Age: 34, Address: "pune"},
		{Name: "shivam", Age: 54, Address: "bangalore"},
	}

	//encoding into JSON format
	human2_enc, err := json.Marshal(human2)

	if err != nil {
		//if error is not nil
		// print error
		fmt.Println(err)
	}

	//printing encoded array
	fmt.Println()
	fmt.Println(string(human2_enc))

}
