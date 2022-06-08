package main

import "fmt"

import (
	"encoding/json" //for json parsing
	"fmt"           //for formatting strings
	"github.com/gorilla/mux"
	"log"       //for logging
	"math/rand" //for random number generation
	"net/http"  //for server
	"os"
	"strconv" //for string conversion
)

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
	router.HandleFunc("/movies/{id]", updateMovies).Methods("PUT")
	router.HandleFunc("movie/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))

	os.Exit(0)

}
