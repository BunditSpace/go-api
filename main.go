package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func getAddresBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Bundit",
		Lastname:  "Wisedphanit",
		Code:      1988,
		Phone:     "0851072564",
	}
	json.NewEncoder(w).Encode(addBook)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddres", getAddresBookAll)
	http.ListenAndServe(getPort(), nil)
}

func main() {
	handleRequest()
}
