package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Protocol Versin: %s\n", resp.Proto)
}
