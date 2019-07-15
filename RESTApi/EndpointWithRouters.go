package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	// Export a field by starting it with an uppercase letter. Otherwise it wont be exported
	Title       string `json:"title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}

var articles []Article

func homePage(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Welcome to the Home page")
	fmt.Println("Endpoint hit: Home Page")
}

func getAllArticles(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint hit: Get All articles")
	//Encode articles into a JSON string
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	// Creates a new instance of mux
	endpointRouter := mux.NewRouter().StrictSlash(true)
	endpointRouter.HandleFunc("/", homePage)
	endpointRouter.HandleFunc("/articles", getAllArticles)
	//Pass the newly created router as the second argument
	log.Fatal(http.ListenAndServe(":8080", endpointRouter))
}

func main() {
	articles = []Article{
		Article{Title: "Sapiens", Description: "Homo Sapiens", Content: "Fantastic book about the human evolution"},
		Article{"Homo Dues", "Homo Deus", "Fantastic book about Human desctruction !!"},
	}

	handleRequests()
}
