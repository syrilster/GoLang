package main

//Building a basic REST Api without routers to display all articles stored in a global array
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnAllArticles(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint hit: Get All articles")
	//fmt.Println(articles)
	//Encode articles into a JSON string
	json.NewEncoder(w).Encode(articles)
}

func main() {
	articles = []Article{
		Article{Title: "Sapiens", Description: "Homo Sapiens", Content: "Fantastic book about the human evolution"},
		Article{"Homo Dues", "Homo Deus", "Fantastic book about Human desctruction !!"},
	}

	handleRequests()
}
