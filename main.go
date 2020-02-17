package main

import (
	"log"
	"net/http"
)

func main() {
	Init()
	go Update()
	go Save()

	http.HandleFunc("/img/", HandleImage)
	http.HandleFunc("/res/", HandleResult)
	http.HandleFunc("/ranking/", HandleRanking)
	http.HandleFunc("/ranking_info/", HandleRankingInfo)
	http.HandleFunc("/register/", HandleRegister)
	http.HandleFunc("/", HandleMain)

	log.Println("Server Started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
