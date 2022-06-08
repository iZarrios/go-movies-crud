package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

//for json parsing
//for formatting strings

//for logging
//for random number generation
//for server

//for string conversion

// I am trying to learn go so I am not going to use  any external database
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie //our mock database

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
}

func createMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(req.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, newMovie)
	json.NewEncoder(res).Encode(newMovie)

}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "applications/json")

	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var editedMovie Movie
			_ = json.NewDecoder(req.Body).Decode(&editedMovie)
			editedMovie.ID = params["id"]
			movies = append(movies, editedMovie)
			json.NewEncoder(res).Encode(editedMovie)
			return

		}
	}
}

func main() {

	movies = append(movies,
		Movie{
			ID:    "1",
			Isbn:  "438227",
			Title: "Movie 1",
			Director: &Director{
				FirstName: "John",
				LastName:  "Doe"}})
	movies = append(movies,
		Movie{
			ID:    "1",
			Isbn:  "438228",
			Title: "Movie ",
			Director: &Director{
				FirstName: "John2",
				LastName:  "Doe2"}})

	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id]", updateMovie).Methods("PUT")
	router.HandleFunc("movie/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Starting Server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

	os.Exit(0)

}
