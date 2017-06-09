package main

import (
	"io/ioutil"
	"net/http"
	"flag"
	"fmt"
	"github.com/james65535/ghub-todo-tracker/utils"
	"github.com/james65535/ghub-todo-tracker/clients"
	"github.com/james65535/ghub-todo-tracker/parser"
)

var address = flag.String("address", "localhost:8000", "server address")

func main() {
	flag.Parse()

	http.HandleFunc("/", receivePush)
	http.ListenAndServe(*address, nil)
}

// Receives JSON payload from webhook push
func receivePush(w http.ResponseWriter, r *http.Request) {

	// TODO Check error and provide response
	/* if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Printf("error: %v\n", err)
		panic(err)
	}*/
	// TODO write 200 ok response
	// w.Write(response)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	todoGenerator(&body)
	// TODO determine what to log
	// utils.WebLog()
}

func todoGenerator(b *[]byte) {
	// Parse commit URL from ghub webhook JSON push
	commitUrl, err := parser.ParseCommit(b)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		panic(err)
	}
	utils.WebLog(commitUrl)

	patch, err := clients.CommitsClient(&commitUrl)
	issue, err := parser.ParsePatch(&patch)
	utils.WebLog(issue)
	// err := clients.IssuesClient(&issue)

}