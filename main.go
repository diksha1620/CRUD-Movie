package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "github.com/gorilla/mux"
)

type Movie struct {
	ID     string `json:"id"`
	Title  string `json:"title" gorm:"unique"`
	Editor string `json:"editor"`
	// Director     *Director `json:"director"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// type Director struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Generate the serial id
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")

	// Find the maximum ID value
	maxID := 0
	for _, item := range movies {
		movieID, err := strconv.Atoi(item.ID)
		if err == nil && movieID > maxID {
			maxID = movieID
		}
	}

	// Generate the next ID
	nextID := maxID + 1

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(nextID)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

// Generate new random id
//
//	func createMovie(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content Type", "application/json")
//		var movie Movie
//		_ = json.NewDecoder(r.Body).Decode(&movie)
//		movie.ID = strconv.Itoa(rand.Intn(10000000))
//		movies = append(movies, movie)
//		json.NewEncoder(w).Encode(movies)
//	}
func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func main() {
	movies = append(movies, Movie{ID: "1", Title: "ZNMD", Editor: "Rakesh", CategoryId: "32", CategoryName: "Comedy-Drama"})
	movies = append(movies, Movie{ID: "2", Title: "YJHD", Editor: "Harsh", CategoryId: "22", CategoryName: "Rom-Com"})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("starting server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
