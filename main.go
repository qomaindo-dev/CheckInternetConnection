package main

import (
	"fmt"
	"net/http"
)

func main() {
	if !CheckConnection() {
		fmt.Println("Please, check your internet connection")
	} else {
		fmt.Println("Connection success!")
	}
}

func CheckConnection() (ok bool) {
	_, err := http.Get("https://www.google.com/")
	if err != nil {
		return false
	}
	return true
}
