package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string   `json:"id"`
	ISBN     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range movies {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	fmt.Printf("Added id %s .....\n", movie.ID)
	json.NewEncoder(w).Encode(movies)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	for index, item := range movies {
		if param["id"] == item.ID {
			fmt.Printf("Updating id %s.....\n", item.ID)
			movieUpdateHelper(movie, index)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func movieUpdateHelper(movie Movie, index int) {
	movies[index].Title = movie.Title
	movies[index].ISBN = movie.ISBN
	movies[index].Title = movie.Title
	movies[index].Director = movie.Director
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.ID == param["id"] {
			fmt.Printf("Deleting id %s...\n", item.ID)
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func main() {
	movies = append(movies, Movie{ID: "1", ISBN: "438227", Title: "Movie One", Director: Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", ISBN: "454553", Title: "Movie Two", Director: Director{Firstname: "Steve", Lastname: "Smith"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Starting Server on Port 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}
