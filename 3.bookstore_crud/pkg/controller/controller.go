package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/bookstore/pkg/models"
	"example.com/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Create := &models.Book{}
	utils.ParseBody(r, Create)
	b := Create.CreateBook()
	res,_ := json.Marshal(b)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBook()
	res, _ := json.Marshal(NewBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["BookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Println(err)
	}
	BookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(BookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	BookId := vars["BookId"]
	ID,err := strconv.ParseInt(BookId,0,0); if err != nil{
		log.Println("Error While Parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["BookId"]
	ID,err := strconv.ParseInt(BookId,0,0);if err != nil{
		log.Fatal(err)
	} 
	BookDetails := models.DeleteBook(ID)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	res,_ := json.Marshal(BookDetails)
	w.Write(res)
}
