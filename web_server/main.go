package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Settings struct {
	SearchEngine []string `json:"searchEngine"`
	Model        []string `json:"model"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, settings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var settings Settings

func main() {
	searchEngineData, err := ioutil.ReadFile("settings/searchEngine.json")
	if err != nil {
		panic(err)
	}
	modelData, err := ioutil.ReadFile("settings/model.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(searchEngineData, &settings.SearchEngine); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(modelData, &settings.Model); err != nil {
		panic(err)
	}

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8888", nil)
}
