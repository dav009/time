package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	SecretKey string
	Address   string
}

var config = Config{
	SecretKey: os.Getenv("SECRET_KEY"),
	Address:   os.Getenv("ADDRESS"),
}

func say_hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("Hello endpoint got called by %s", req.Referer())
	msg := []byte(fmt.Sprintf("<h1>hello secret food is: %s </h1>", config.SecretKey))
	w.Write(msg)
}

func main() {
	log.Printf("Starting Server listening on %s\n", config.Address)
	server := http.NewServeMux()
	server.HandleFunc("/", say_hello)
	err := http.ListenAndServe(config.Address, server)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
