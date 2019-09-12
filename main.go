package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Bundit",
		Lastname:  "Wisedphanit",
		Code:      001,
		Phone:     "08x-xxx-xxxx",
	}
	json.NewEncoder(w).Encode(addBook)
}

func getAvailableHarddiskSpace(w http.ResponseWriter, r *http.Request) {
	var result string
	var available string
	query := r.URL.Query()
	size := query.Get("size")
	if size != "" {
		realSize, err := strconv.Atoi(size)
		if err == nil {
			available = strconv.Itoa(realSize * 1000 * 1000 * 1000 / 1024 / 1024 / 1024)
		}
	}
	result = fmt.Sprintf("Hardisk Size %s available space : %s", size, available)
	json.NewEncoder(w).Encode(result)
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
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.HandleFunc("/getAvailableHarddiskSpace", getAvailableHarddiskSpace)
	http.ListenAndServe(getPort(), nil)
}

func main() {
	handleRequest()
}
