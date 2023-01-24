package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterStoreRoutes(r)
	http.Handle("/",r)
	fmt.Println("Starting Web Server On 8080")
	err := http.ListenAndServe(":8000",r); if err != nil{
		log.Fatal(err)
	}
}