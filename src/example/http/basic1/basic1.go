package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello " + html.EscapeString(req.URL.Path)))
		fmt.Println("Hello ", html.EscapeString(req.URL.Path))
	})

	http.HandleFunc("/hi", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hi"))
		fmt.Println("Hi")
	})

	http.ListenAndServe(":8000", nil)
	fmt.Println("######## Server Start############")

}
