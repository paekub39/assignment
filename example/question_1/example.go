package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	host := "http://localhost:1110"
	fmt.Println("Get XYZ by position")
	getByPosition(&host)
	fmt.Println("")
	fmt.Println("Get XYZ by remove know data")
	getByRemoveKnowData(&host)
}

func getByPosition(host *string) {
	fullPath := fmt.Sprintf("%s/get-xyz-by-position", *host)
	resp, err := http.Get(fullPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := string(b)
	fmt.Printf("%s", result)
}

func getByRemoveKnowData(host *string) {
	fullPath := fmt.Sprintf("%s/get-xyz-by-remove-know-data", *host)
	resp, err := http.Get(fullPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := string(b)
	fmt.Printf("%s", result)
}
