package main

import (
	"example/rest-service/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController()

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.RemoveUser)
	//
	// r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	u := models.User{
	// 		Name:   "capkum",
	// 		Gender: "male",
	// 		Age:    43,
	// 		ID:     p.ByName("id"),
	// 	}
	//
	// 	uj, _ := json.Marshal(u)
	//
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(200)
	// 	fmt.Fprintf(w, "%s", uj)
	// })
	//
	http.ListenAndServe("localhost:8000", r)
}
