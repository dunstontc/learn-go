package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.geekwiseacademy.com/")
	check(err)
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	check(err)
	fmt.Printf("%s", page)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}