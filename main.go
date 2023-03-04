package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
)

type Article struct {
	ID					int64			`json:"id"`
	Title 			string 		`json:"title"`
	Description string 		`json:"description"`
	Content 		string 		`json:"content"`
	// Created_At 	time.Time `json:"created_at"`
	Likes 			int 			`json:"likes"`
	Dislikes 		int 			`json:"dislikes"`
}

var Articles = []Article {
	Article{
		ID: 1,
		Title: "Hellao", 
		Description: "Article Description", 
		Content: "Article Content", 
		Likes: 1, 
		Dislikes: 5,
	},
	Article{
		ID: 2,
    Title: "Bonjour", 
    Description: "Article Description", 
    Content: "Article Content", 
    Likes: 1, 
    Dislikes: 5,
	},
	Article{
		ID: 3,
    Title: "Ciao", 
    Description: "Article Description", 
    Content: "Article Content", 
    Likes: 1, 
    Dislikes: 5,
	},
}

func GetAllArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Returning all articles")
	json.NewEncoder(w).Encode(Articles)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
  if id == "" {
		w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Invalid article id")
    return
  }
	fmt.Println("Returning a single article")
	fmt.Println(id)
	// json.NewEncoder(w).Encode(Article[id])
}

func Home(w http.ResponseWriter, r *http.Request) {
	var randomString = []string {
		"Hi, nice to meet you!",
		"Great to hear from you!",
		"All hail! We shall have a dinner!",
	}
	
	json.NewEncoder(w).Encode(randomString[rand.Intn(len(randomString))])
}

// func HandlerOver(w http.ResponseWriter, r *http.Request)  {
// 	// The "/" pattern matches everything, so we need to check
// 	// that we're at the root here
// 	if (r.URL.Path !=  "/") {
//     http.NotFound(w, r)
//     return
//   }
// 	fmt.Fprintf(w, "Nothing's on here")
// }

func handleRequests() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/articles", GetAllArticles)
	http.HandleFunc("/articles/{id}", GetArticleById)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main()  {
	handleRequests()
}