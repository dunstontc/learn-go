package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."
	rdr := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr)
	check(err)

	rdr2 := bytes.NewBuffer([]byte(msg))
	_, err = io.Copy(os.Stdout, rdr2)
	check(err)

	res, err := http.Get("http://www.mcleods.com")
	check(err)

	io.Copy(os.Stdout, res.Body)
	if err := res.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
