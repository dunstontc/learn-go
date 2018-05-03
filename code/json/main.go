package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Jeffail/gabs"
)

func main() {

	fmt.Println("go!")
	packFile := "./modules.json"
	jsonBytes := getBytes(packFile)
	jsonParsed, _ := gabs.ParseJSON(jsonBytes)

	children, _ := jsonParsed.S("dependencies").ChildrenMap()
	for key, child := range children {
		var pack Package
		err := json.Unmarshal(child.Bytes(), &pack)
		if err != nil {
			panic(err)
		}
		pack.Name = key
		fmt.Println(pack)
	}

}

// Package struct
type Package struct {
	Name     string `json:"name,omitempty"`
	Version  string `json:"version,omitempty"`
	From     string `json:"from,omitempty"`
	Resolved string `json:"resolved,omitempty"`
}

// Implement the Stringer interface for printing
func (p Package) String() string {
	return fmt.Sprintf("{Package: %s, %s, %s, %s", p.Name, p.From, p.Version, p.Resolved)
}

// Gets a slice of bytes from a file.
func getBytes(file string) []byte {
	theBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Demonic Invasion In Progress: %s", err.Error())
		return []byte{}
	}
	return theBytes
}

// // Unmarshal using a generic interface
// func makeFace(jsonBytes []byte) map[string]interface{} {
// 	var f interface{}
// 	err := json.Unmarshal(jsonBytes, &f) // Unmarshal using a generic interface
// 	if err != nil {
// 		log.Println("Error parsing JSON: ", err)
// 	}
//
// 	itemsMap := f.(map[string]interface{}) // JSON object parses into a map with string keys
// 	return itemsMap
// }

// func parseFace(itemsMap map[string]interface{}) {
// 	for _, value := range itemsMap { // Loop through the Items; we're not interested in the key, just the values
// 		switch jsonObj := value.(type) { // Use type assertions to ensure that the value is a JSON object
// 		case interface{}: // The value is an Item, represented as a generic interface
// 			var pack Package
// 			for itemKey, itemValue := range jsonObj.(map[string]interface{}) { // Access the values in the JSON object and place them in an Item
// 				switch itemKey {
// 				case "version":
// 					switch itemValue := itemValue.(type) {
// 					case string:
// 						pack.Version = itemValue
// 					default:
// 						fmt.Printf("Incorrect type for %s: %v\n", itemKey, itemValue)
// 					}
// 				case "from":
// 					switch itemValue := itemValue.(type) {
// 					case string:
// 						pack.From = itemValue
// 					default:
// 						fmt.Printf("Incorrect type for %s: %v\n", itemKey, itemValue)
// 					}
// 				case "resolved":
// 					switch itemValue := itemValue.(type) {
// 					case string:
// 						pack.Resolved = itemValue
// 					default:
// 						fmt.Printf("Incorrect type for %s: %v\n", itemKey, itemValue)
// 					}
// 				default:
// 					fmt.Println("Unknown key for Item found in JSON")
// 				}
// 			}
// 			fmt.Println(pack)
// 		default:
// 			fmt.Println("Expecting a JSON object; got something else") // Not a JSON object; handle the error
// 		}
// 	}
// }
