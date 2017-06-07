package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main () {
	http.HandleFunc("/", printRequest)
	http.ListenAndServe("localhost:8000", nil)
}

func printRequest(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type","application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Write(response)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("request:")
		fmt.Println(string(body))
	}
}
