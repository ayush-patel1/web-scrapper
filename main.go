package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
    "github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in env")
	}

	router:=chi.NewRouter()

	srv:=&http.Server{
		Handler: router,
		Addr: ":"+ portString,
	}
    log.Printf("server starting on port %v",portString)
	err:=srv.ListenAndServe()
    if(err!=nil){
		log.Fatal(err)
	}
	// fmt.Println("Port:", portString)
}
