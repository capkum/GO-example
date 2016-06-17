package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

func returnArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	vars1 := ps.ByName("vars1")
	vars2 := ps.ByName("vars2")

	fmt.Fprintln(w, "Key: "+key)
	fmt.Fprintln(w, "vars2: "+vars1)
	fmt.Fprintln(w, "vars2: "+vars2)
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
	w.Header().Set("Content-Type", "application/json")
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
	router := httprouter.New()
	router.HandlerFunc("GET", "/", homePage)
	router.HandlerFunc("GET", "/all", returnAllArticles)
	router.GET("/article/:key/:vars1/:vars2", returnArticle)
	router.HandlerFunc("GET", "/delete", delArticle)
	router.HandlerFunc("GET", "/add", addArticle)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	fmt.Println("######## server start #########")
	go handleRequest()
	fmt.Scanln()
}
