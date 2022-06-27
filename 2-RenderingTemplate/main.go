package main

//import necesssary modul
import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func index(w http.ResponseWriter, f *http.Request){
	var filepath = path.Join("views","index.html")
	var tmpl, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	var data = map[string]interface{}{
		"title" : "Main Page",
		"name"	: "lintang",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Server started at localhost:8090")
	http.ListenAndServe(":8090", nil)
}
