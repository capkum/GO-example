package main

import (
	"fmt"
	"html"
	"net/http"
	"sync"
)

var count int
var mutex sync.Mutex

func echoString(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello " + html.EscapeString(req.URL.Path)))
	fmt.Println("Hello ", html.EscapeString(req.URL.Path))
}

func counter(res http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	count++
	fmt.Fprintf(res, "Count %d\n", count)
	mutex.Unlock()
}
func hi(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hi"))
	fmt.Println("Hi")
}

func main() {
	http.HandleFunc("/", echoString)
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/hi", hi)

	fmt.Println("######## Server Start############")
	http.ListenAndServe(":8000", nil)

}
