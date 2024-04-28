package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Post struct {
	ID      string `json:"id"`
	Isbn    string `json:"isbn"`
	Content string `json:"content"`
	User    *User  `json:"user"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

var posts []Post

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Gin setup
	rGin := gin.Default()
	rGin.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Gorilla Mux setup
	rMux := mux.NewRouter()
	rMux.HandleFunc("/posts", getPosts).Methods("GET")
	rMux.HandleFunc("/posts/{id}", getPost).Methods("GET")
	rMux.HandleFunc("/posts", createPost).Methods("POST")
	rMux.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	rMux.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	// Start Gin server
	go func() {
		fmt.Printf("Starting Gin server at 8080\n")
		if err := rGin.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	}()

	// Start Gorilla Mux server
	fmt.Printf("Starting Gorilla Mux server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", rMux))
}

// Gin Handlers
func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

// Gorilla Mux Handlers
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}
	if err := json.NewEncoder(w).Encode(&Post{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			post.ID = params["id"]
			posts = append(posts, post)
			if err := json.NewEncoder(w).Encode(post); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
