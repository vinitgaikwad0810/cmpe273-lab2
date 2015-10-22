package main

import (  
    // Standard library packages
	"encoding/json"
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)
type JSONRequest struct {
	   Name   string `json:"name"`
      
}

type JSONResponse struct {
	Greeting string `json:"greeting"`	
}

func postRequestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
   
    request:= JSONRequest{}
    response := JSONResponse{}
    json.NewDecoder(r.Body).Decode(&request)
    
	response.Greeting = "Hello, "+request.Name
    uj, _ := json.Marshal(response)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", uj)
}



func main() {  
    r := httprouter.New()


	r.POST("/hello", postRequestHandler)
	 server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: r,
    }
    server.ListenAndServe()


}
// curl -H "Content-Type: application/json" -X POST -d '{"username":"xyz","password":"xyz"}' http://localhost:8080 