package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := []Article{
		Article{Title:"Test Title", Desc: "Test Description", Content: "Test Content"},
		Article{Title:"Test Title2", Desc: "Test Description2", Content: "Test Content2"},
	}
	fmt.Println("Endpoint Hit: All Articels")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequest()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	Articles = []Article{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	handleRequest()
}