package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title		string	`json:"Title"`
	Desc		string	`json:"desc"`
	Content		string	`json:"content"`
}

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage")
	fmt.Println("endpoint hit: Homepage")
}

func article(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Articles)
	fmt.Println("endpoint hit: article")
}

func main(){
	Articles = []Article{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	http.HandleFunc("/", homepage)
	http.HandleFunc("/article", article)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
