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
	data := []byte(`{
  		"person": {
		    "name": {
		      "first": "Leonid",
		      "last": "Bugaev",
		      "fullName": "Leonid Bugaev"
		    },
		    "github": {
		      "handle": "buger",
		      "followers": 109
		    },
		    "avatars": [
		      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
		    ]
		    },
		    "company": {
		      "name": "Acme"
  		}
	}`)
	responseParse := PushParse(data)
	utils.WebLog(string(responseParse))


}
