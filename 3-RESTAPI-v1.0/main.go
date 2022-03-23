package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Id		string	`json:"Id"`
	Title		string	`json:"Title"`
	Desc		string	`json:"desc"`
	Content		string	`json:"content"`
}

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage")
	fmt.Println("endpoint hit: Homepage")
}

func returnAllArticle(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Articles)
	fmt.Println("endpoint hit: article")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for _, artic := range Articles {
		if artic.Id == key {
			json.NewEncoder(w).Encode(artic)
		}
	}

	fmt.Println("endpoint hit: single article")
}

func createNewArticle(w http.ResponseWriter, r *http.Request){
    	reqBody, _ := ioutil.ReadAll(r.Body)
    	fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func handleFunction(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", returnAllArticle)

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8880", myRouter))
}

func main(){
	fmt.Println("RESTAPI v2.0")
    Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }

	handleFunction()
}
