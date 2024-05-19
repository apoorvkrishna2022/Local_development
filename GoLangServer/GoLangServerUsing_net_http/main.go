package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Movie struct represents the data structure for a movie
type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
}

// MovieStore interface defines the CRUD operations for movies
type MovieStore interface {
	GetMovie(id int) (*Movie, error)
	GetMovies() ([]Movie, error)
	AddMovie(movie Movie) error
	ModifyMovie(id int, updatedMovie Movie) error
	DeleteMovie(id int) error
}

// InMemoryMovieStore is an in-memory implementation of MovieStore
type InMemoryMovieStore struct {
	movies map[int]Movie
}

// NewInMemoryMovieStore creates a new instance of InMemoryMovieStore
func NewInMemoryMovieStore() *InMemoryMovieStore {
	return &InMemoryMovieStore{
		movies: make(map[int]Movie),
	}
}

func (s *InMemoryMovieStore) GetMovie(id int) (*Movie, error) {
	movie, ok := s.movies[id]
	if !ok {
		return nil, fmt.Errorf("Movie with ID %d not found", id)
	}
	return &movie, nil
}

func (s *InMemoryMovieStore) GetMovies() ([]Movie, error) {
	fmt.Println("inside get movies function")
	var movieList []Movie
	for _, movie := range s.movies {
		movieList = append(movieList, movie)
	}
	return movieList, nil
}

func (s *InMemoryMovieStore) AddMovie(movie Movie) error {
	s.movies[movie.ID] = movie
	return nil
}

func (s *InMemoryMovieStore) ModifyMovie(id int, updatedMovie Movie) error {
	_, ok := s.movies[id]
	if !ok {
		return fmt.Errorf("Movie with ID %d not found", id)
	}
	s.movies[id] = updatedMovie
	return nil
}

func (s *InMemoryMovieStore) DeleteMovie(id int) error {
	_, ok := s.movies[id]
	if !ok {
		return fmt.Errorf("Movie with ID %d not found", id)
	}
	delete(s.movies, id)
	return nil
}

// HandleGetMovie handles the "/getMovie" endpoint
func HandleGetMovie(w http.ResponseWriter, r *http.Request, store MovieStore) {
	id := 1 // You can extract the ID from the request URL or request parameters
	movie, err := store.GetMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}

// HandleGetMovies handles the "/getMovies" endpoint
func HandleGetMovies(w http.ResponseWriter, r *http.Request, store MovieStore) {
	movies, err := store.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, movies)
}

// HandleAddMovie handles the "/addMovie" endpoint
func HandleAddMovie(w http.ResponseWriter, r *http.Request, store MovieStore) {
	fmt.Println("inside handle add movie ", r.Method, r.Body)
	var newMovie Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = store.AddMovie(newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// HandleModifyMovie handles the "/modifyMovie" endpoint
func HandleModifyMovie(w http.ResponseWriter, r *http.Request, store MovieStore) {
	fmt.Println("inside handle modify movie")
	id := 1 // You can extract the ID from the request URL or request parameters
	var updatedMovie Movie
	err := json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = store.ModifyMovie(id, updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleDeleteMovie handles the "/deleteMovie" endpoint
func HandleDeleteMovie(w http.ResponseWriter, r *http.Request, store MovieStore) {
	fmt.Println("handle Delete movie")
	id := 1 // You can extract the ID from the request URL or request parameters
	err := store.DeleteMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func main() {
	store := NewInMemoryMovieStore()

	http.HandleFunc("/getMovie", func(w http.ResponseWriter, r *http.Request) {
		HandleGetMovie(w, r, store)
	})

	http.HandleFunc("/getMovies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle get movies")
		HandleGetMovies(w, r, store)
	})

	http.HandleFunc("/addMovie", func(w http.ResponseWriter, r *http.Request) {
		HandleAddMovie(w, r, store)
	})

	http.HandleFunc("/modifyMovie", func(w http.ResponseWriter, r *http.Request) {
		HandleModifyMovie(w, r, store)
	})

	http.HandleFunc("/deleteMovie", func(w http.ResponseWriter, r *http.Request) {
		HandleDeleteMovie(w, r, store)
	})

	fmt.Println("Server is running on :8080 and yes and i know i am the best..")
	http.ListenAndServe(":8080", nil)
}
