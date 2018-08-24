package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	SecretFood string
	Flavor     string
}

var config = Config{
	SecretFood: os.Getenv("SECRETFOOD"),
	Flavor:     os.Getenv("FLAVOR"),
}

func say_hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("Hello endpoint got called by %s", req.Referer())
	msg := []byte(fmt.Sprintf("<h1>hello secret food is: %s . flavour: %s</h1>", config.SecretFood, config.Flavor))
	w.Write(msg)
}

func main() {
	log.Printf("Starting Server listening on 0.0.0.0:8500")
	server := http.NewServeMux()
	server.HandleFunc("/", say_hello)
	err := http.ListenAndServe("0.0.0.0:8500", server)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
