package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
