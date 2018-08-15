package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	// "strings"
)

type words []string

func (p words) Len() int           { return len(p) }
func (p words) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p words) Less(i, j int) bool { return p[i] < p[j] }

// func (p words) lower() {
// 	for _, v := range p {
// 		fmt.Println(strings.ToLower(v))
// 	}
// }

func main() {
	var collection words

	file, err := os.Open("./pkg.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		collection = append(collection, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(collection)
	for _, v := range collection {
		fmt.Println(v)
	}
}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface
