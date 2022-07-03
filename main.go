package main

import (
	"WordleServer/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
