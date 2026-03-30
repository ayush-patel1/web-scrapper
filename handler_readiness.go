package main

import "net/http"

func handlerReadiness(w http.ResponseWriter,R *http.Response){
	respondWithJSON(w,200,struct{}{})
}