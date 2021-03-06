package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func createArticle(responseWriter http.ResponseWriter, request *http.Request) {
	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Fprintf(responseWriter, "Error during creation")
	}
	fmt.Fprintf(responseWriter, "Creating new article.. \n")
	var article Article
	json.Unmarshal(reqBody, &article)
	fmt.Println(article)
	//Update the articles collection with the latest
	articles = append(articles, article)
	json.NewEncoder(responseWriter).Encode(article)
}

func deleteArticleById(responseWriter http.ResponseWriter, request *http.Request) {
	requestParams := mux.Vars(request)
	id := requestParams["id"]

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
			fmt.Fprintf(responseWriter, "Article successfully deleted \n")
			break
		}
	}
}

func updateArticleById(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint hit: Update Article")
	requestParams := mux.Vars(request)
	reqBody, err := ioutil.ReadAll(request.Body)
	id := requestParams["id"]
	if err != nil {
		fmt.Fprintf(responseWriter, "Error during updation \n")
	}

	var updatedArticle Article
	json.Unmarshal(reqBody, &updatedArticle)

	for index, article := range articles {
		if article.Id == id {
			fmt.Println("Deleting existing article")
			articles = append(articles[:index], articles[index+1:]...)
			fmt.Println("Updating article: ", article.Id)
			article.Id = updatedArticle.Id
			article.Title = updatedArticle.Title
			article.Description = updatedArticle.Description
			article.Content = updatedArticle.Content
			articles = append(articles, article)
			json.NewEncoder(responseWriter).Encode(article)
			break
		}
	}
}

func handleRequests() {
	// Creates a new instance of mux
	endpointRouter := mux.NewRouter().StrictSlash(true)
	//Somehow the router order matters here
	endpointRouter.HandleFunc("/", homePage)
	endpointRouter.HandleFunc("/articles", getAllArticles)
	endpointRouter.HandleFunc("/article", createArticle).Methods("POST")
	endpointRouter.HandleFunc("/article/{id}", deleteArticleById).Methods("DELETE")
	endpointRouter.HandleFunc("/article/{id}", updateArticleById).Methods("PUT")
	endpointRouter.HandleFunc("/article/{id}", getArticleById)

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
