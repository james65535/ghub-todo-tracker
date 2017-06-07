package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/james65535/ghub-todo-tracker/utils"
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
		sBody := string(body)
		utils.WebLog(sBody)
		fmt.Printf("request: \n%v", sBody)
	}
}
