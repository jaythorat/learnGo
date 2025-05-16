package main

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message11"`
	Name string 
}

func Handler(w http.ResponseWriter, r *http.Request){
	resp := HelloResponse{Message: "HI JAY",Name: "jay"}
	json.NewEncoder(w).Encode(resp)
}



func main() {
	http.HandleFunc("/hello",Handler)
	http.ListenAndServe(":8080",nil)

}
