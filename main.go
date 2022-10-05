package main

import (
	"fmt"
	"net/http"
	"time"
)


func main(){

	serveMux := http.NewServeMux()       // creates HTTP Multiplexer, use to match the pattern and call the associated handler
	serveMux.HandleFunc("/",testHandler) //associate the path '/' with the testHandler function 

	const addr = "localhost:8000"		// the server will listen on the following port 
	
	srv := http.Server {
		Handler : serveMux,
		Addr : addr,
		WriteTimeout : 30 * time.Second,
		ReadTimeout : 30 * time.Second,
	}
	err := srv.ListenAndServe() 	//  the server on port 8000 to listen and server the respective handler 
	fmt.Println(err)
}


//handler add sets the content type to app/json , add a success code and an empty JSON body
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json") 	//sets content Type
	w.WriteHeader(200)								   	// status code 200
	w.Write([]byte("{}"))								// empty JSON {}
}

