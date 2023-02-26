package main

import (
	"fmt"
	"log"
	"net/http"
	"errors"
)

func hello(name string) (string, error)  {
	if (name == "") {
		return "", errors.New("Empty name parameter")
	}
	
	message := fmt.Sprintf("Message to %s", name)

	return message, nil
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to Main Go!")
	fmt.Println("Endpoint hit: homePage")
	fmt.Println(hello("Firhan"))
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main()  {
	handleRequests()
}