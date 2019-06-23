package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, _ *http.Request){
	_, err := w.Write([]byte("Hello"))
	if err != nil{
		log.Fatal(err)
	}
}

func ping(w http.ResponseWriter, _ *http.Request){
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Pong"))
	if err != nil{
		log.Fatal(err)
	}
}