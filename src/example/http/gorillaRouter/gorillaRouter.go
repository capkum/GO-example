package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	vars1 := vars["vars1"]
	vars2 := vars["vars2"]

	fmt.Println("vars1: " + vars1)
	fmt.Println("vars2: " + vars2)
	fmt.Fprintf(w, "Key: "+key)
	// fmt.Fprintf(w, "returns a specific article")
	fmt.Println("Endpoint Hit: returnArticle")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	article := Articles{
		Article{
			Title:   "Hello",
			Desc:    "Article Description",
			Content: "Article Content",
		},
		Article{
			Title:   "Hello 1",
			Desc:    "Article Description 1",
			Content: "Article Content 1",
		},
	}
	// fmt.Fprintf(w, "All Articles")
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(article)
}

func addArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Adds an article to list of articles")
	fmt.Println("Endpoint Hit: addArticle")
}

func delArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deletes a specific article")
	fmt.Println("Endpoint Hit: delArticle")
}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnAllArticles)
	router.HandleFunc("/article/{key}/{vars1}/{vars2}/", returnArticle)
	router.HandleFunc("/delete", delArticle)
	router.HandleFunc("/add", addArticle)
	log.Fatal(http.ListenAndServe(":8081", router))
	// http.ListenAndServe(":8081", nil)
}

func main() {
	fmt.Println("######## server start #########")
	handleRequest()
}
