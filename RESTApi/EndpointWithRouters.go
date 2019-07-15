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
	Id          string `json:"id"`
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

func getArticleById(w http.ResponseWriter, request *http.Request) {
	requestParams := mux.Vars(request)
	key := requestParams["id"]
	//fmt.Fprintf(w, "Key"+key)
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	fmt.Fprintf(w, "Invalid Article Id provided. ID: "+key)
}

func handleRequests() {
	// Creates a new instance of mux
	endpointRouter := mux.NewRouter().StrictSlash(true)
	endpointRouter.HandleFunc("/", homePage)
	endpointRouter.HandleFunc("/articles", getAllArticles)
	endpointRouter.HandleFunc("/articles/{id}", getArticleById)
	//Pass the newly created router as the second argument
	log.Fatal(http.ListenAndServe(":8080", endpointRouter))
}

func main() {
	articles = []Article{
		Article{Id: "1", Title: "Sapiens", Description: "Homo Sapiens", Content: "Fantastic book about the human evolution"},
		Article{"2", "Homo Dues", "Homo Deus", "Fantastic book about Human desctruction !!"},
	}

	handleRequests()
}
