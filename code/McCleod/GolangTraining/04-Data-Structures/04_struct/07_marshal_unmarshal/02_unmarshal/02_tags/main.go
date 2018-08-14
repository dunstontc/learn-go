package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 person
	fmt.Println(p1.First) //
	fmt.Println(p1.Last)  //
	fmt.Println(p1.Age)   // 0

	byteSlice := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	json.Unmarshal(byteSlice, &p1)

	fmt.Println("--------------") // --------------
	fmt.Println(p1.First)         // James
	fmt.Println(p1.Last)          // Bond
	fmt.Println(p1.Age)           // 20
	fmt.Printf("%T \n", p1)       // main.person
}
