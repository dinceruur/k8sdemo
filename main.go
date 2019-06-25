package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

const cfgPath = "/config/demo.yaml"

var config struct {
	Message string
}

func init() {
	cfgData, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(cfgData, &config); err != nil {
		log.Fatalf("failed to parse config file %q: %v", cfgPath, err)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, _ *http.Request) {
	message := fmt.Sprintf("Hello, %v", config.Message)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Pong"))
	if err != nil {
		log.Fatal(err)
	}
}
