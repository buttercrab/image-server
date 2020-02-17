package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Id    int
	Value int
	Name  string
}

type Profile struct {
	Name   string
	Secret string
}

func send404(w http.ResponseWriter) {
	sendFile(w, "public/404.html")
}

func sendFile(w http.ResponseWriter, file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error: %s", err.Error())
		return
	}
	w.WriteHeader(200)
	_, _ = w.Write(content)
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		send404(w)
		return
	}
	sendFile(w, "public/index.html")
}

func HandleImage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/img/" || r.Method != "GET" {
		send404(w)
		return
	}
	w.WriteHeader(200)
	id, url := GetNew()
	_, _ = w.Write([]byte(fmt.Sprintf(`{"id":%d,"url":"%s"}`, id, url)))
}

func HandleResult(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/res/" || r.Method != "POST" {
		send404(w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var t Result
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error: %s", err.Error())
		return
	}
	SetValue(t.Id, t.Value, t.Name)
	w.WriteHeader(200)
}

func HandleRanking(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ranking/" || r.Method != "GET" {
		send404(w)
		return
	}
	sendFile(w, "public/ranking.html")
}

func HandleRankingInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ranking_info/" || r.Method != "GET" {
		send404(w)
		return
	}

	str, err := json.Marshal(rank)

	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error: %s", err.Error())
		return
	}

	w.WriteHeader(200)
	_, _ = w.Write(str)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register/" || r.Method != "POST" {
		send404(w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var t Profile
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error: %s", err.Error())
		return
	}
	var res = Register(t.Name, t.Secret)
	w.WriteHeader(200)
	_, _ = w.Write([]byte(fmt.Sprintf("%t", res)))
}
