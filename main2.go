package main

import (
	"fmt"
	//"net/http"
)

func tmain() {
	fmt.Println("Hello World")
	b, res := run("GET", "http://quotes.rest/qod.json")
	fmt.Printf("%v", res)
	fmt.Println("")
	fmt.Printf(string(b))
	//GetMedia("http://quotes.rest/qod.json")
}
