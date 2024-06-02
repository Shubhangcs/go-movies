package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	IMDB string `json:"imdb"`
	TITLE string `json:"title"`
	DIRECTOR *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies =  []Movie{
	{ID: "1" , IMDB: "8.9" , TITLE: "The Movie 1" , DIRECTOR: &Director{FirstName: "Director" , LastName: "1"}},
	{ID: "2" , IMDB: "4.9" , TITLE: "The Movie 2" , DIRECTOR: &Director{FirstName: "Director" , LastName: "2"}},
}

func getMovies(res http.ResponseWriter , req *http.Request){
	res.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(res).Encode(movies)
}

func getMoviesById(res http.ResponseWriter , req *http.Request){
	res.Header().Set("Content-Type","application/json")
	parameters := mux.Vars(req)

	for _ , movie := range movies{
		if movie.ID == parameters["id"] {
			json.NewEncoder(res).Encode(movie)
			break
		}
	}
}

func deleteMovies(res http.ResponseWriter , req *http.Request){
	res.Header().Set("Content-Type","application/json")
	parameters := mux.Vars(req)

	for index , movie := range movies{
		if movie.ID == parameters["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}

func createMovie(res http.ResponseWriter , req *http.Request){
	res.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movies)
}

func updateMovies(res http.ResponseWriter , req *http.Request){
	res.Header().Set("Content-Type","application/json")
	parameters := mux.Vars(req)
	for index , movie := range movies{
		if movie.ID == parameters["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movies = append(movies, movie)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}


func main(){
	var PORT = "8000"
	r := mux.NewRouter()

	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMoviesById).Methods("GET")
	r.HandleFunc("/movies/{id}" , deleteMovies).Methods("DELETE")
	r.HandleFunc("/movies" , createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}" , updateMovies).Methods("PUT")

	log.Printf("Server Is Running at Port %v" , PORT)

	log.Fatal(http.ListenAndServe(":"+PORT , r))
}