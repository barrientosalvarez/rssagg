package main

import (
    "encoding/json"
    "log"
    "net/http"
)


func respondWithSON(w http.ResponseWriter, code int, payload interface{}){
    dat, err := json.Marshal(payload)
    if err != nil{
        log.Println("Failed to marshal JSON response: %v", payload)
        w.WriteHeader(500)
        return 
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(dat)
}
