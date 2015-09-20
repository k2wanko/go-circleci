package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/k2wanko/go-circleci"
)

func main() {
	c := circleci.NewHTTPClient("API Token")

	res, err := c.Get("https://circleci.com/api/v1/me")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)
}
