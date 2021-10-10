package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Config()
	r := router.Create()
	fmt.Println("Listen on port 3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
