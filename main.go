package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"
)

// Movie struct represents a movie object with ID, Isbn, Title and Director fields
type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

// Director struct represents a director object with Firstname and Lastname fields
type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie //create movies slice

// getMovies function returns all movies in the movies slice
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set json content type
	json.NewEncoder(w).Encode(movies) //encode movies slice to json and return it
}

// getMovie function returns a single movie with the given ID
func getMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set json content type
	params := mux.Vars(r) //get url params

	for _, item := range movies {
		if item.ID == params["id"] { //check if ID of movie matches the ID in the params
			json.NewEncoder(w).Encode(item) //encode movie object
			return 
		}
	}
}

// createMovie function creates a new movie object and appends it to the movies slice
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set json content type
	var movie Movie //create movie variable
	_ = json.NewDecoder(r.Body).Decode(&movie) //decode json object from request body into struct
	movie.ID = strconv.Itoa((rand.Intn(1000000000))) //generate random ID for movie
	movies = append(movies, movie) //append movie to movies slice
	json.NewEncoder(w).Encode(movie) //encode movie object and return it
}

// updateMovie function updates an existing movie object with the given ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set json content type
	params := mux.Vars(r) // get url params

	//loop through all movies in slice
	for index, item := range movies {
		if item.ID == params["id"] { //check if ID of movie matches the ID in the params
			movies = append(movies[:index], movies[index+1:]...) //delete movie from slice
			var movie Movie //create movie
			_ = json.NewDecoder(r.Body).Decode(&movie) //decode json object from request body into struct
			movie.ID = params["id"] //set ID of movie to ID from params
			movies = append(movies, movie) //append movie to movies slice
			json.NewEncoder(w).Encode(movie) //encode movie object
			return
		}
	}
}

// deleteMovie function deletes an existing movie object with the given ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set json content type
	params := mux.Vars(r) // get url params

	//loop through all movies in slice
	for index, item := range movies {
		if item.ID == params["id"] { //check if ID of movie matches the ID in the params
			movies = append(movies[:index], movies[index+1:]...) //delete movie from slice
			break //break from loop
		}
	}
}

func main () {
	r := mux.NewRouter() //create new router

	//add some movies
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454555", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	//route handlers
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) //listen and serve on port 8000
}