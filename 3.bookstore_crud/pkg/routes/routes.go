package routes

import(
	"github.com/gorilla/mux"
	"example.com/bookstore/pkg/controller"
)

func RegisterStoreRoutes (router *mux.Router){
	router.HandleFunc("/book/",controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}",controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controller.DeleteBook).Methods("DELETE")
}