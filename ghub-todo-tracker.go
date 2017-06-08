package main

import (
	"fmt"
	"github.com/james65535/ghub-todo-tracker/utils"
	"io/ioutil"
	"net/http"
	"flag"
	"github.com/james65535/ghub-todo-tracker/parser"
)

var address = flag.String("address", "localhost:8000", "server address")

func main() {
	flag.Parse()

	http.HandleFunc("/", printRequest)
	http.ListenAndServe(*address, nil)
	// TODO check stuff
	// 2

}

func printRequest(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type","application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Write(response)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	sBody := string(body)
	utils.WebLog(sBody)
	fmt.Printf("request: \n%v", sBody)
	responseParse, err := parser.PushParse(body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Printf("error: %v\n", err)
		panic(err)
	}
	utils.WebLog(string(responseParse))
	// TODO tester

}
