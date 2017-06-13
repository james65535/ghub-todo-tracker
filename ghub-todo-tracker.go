package main

import (
	"io/ioutil"
	"net/http"
	"flag"
	"fmt"
	"github.com/james65535/ghub-todo-tracker/util"
	"github.com/james65535/ghub-todo-tracker/clients"
	"github.com/james65535/ghub-todo-tracker/parser"
)

var address = flag.String("address", "localhost:8000", "server address")

func main() {
	// TODO temp to test ghub issue
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
	// util.WebLog()
}

func todoGenerator(b *[]byte) {

	// Parse commit URL from ghub webhook JSON push
	/* FIXME add error handling: github.com/james65535/ghub-todo-tracker/parser.ParseCommit(0xc420041c80, 0x0, 0x0, 0xc4200f0000, 0x1871)
	ghub-todo-tracker/parser/ghub_json_parser.go:24 +0x1d0
	*/

	commitUrl, err := parser.ParseCommit(b)
	if err != nil {
		detail := fmt.Errorf("Error getting commit url: %v", err)
		fmt.Println(detail)
	}
	util.WebLog(commitUrl)

	// Get Patch JSON from Ghub
	patch, err := clients.CommitsClient(&commitUrl)
	if err != nil {
		detail := fmt.Errorf("Error getting patch: %v", err)
		fmt.Println(detail)
	}

	// Get to-do statement(s) from JSON
	issue, err := parser.ParsePatch(&patch)
	if err != nil {
		detail := fmt.Errorf("Error getting issue: %v", err)
		fmt.Println(detail)
	}
	for i, v := range issue {
		if v[1] == "+" {
			issueLog := fmt.Sprintf("Issue %v: %v\n", i, v[2])
			util.WebLog(issueLog)
			var issue = clients.GhubIssue{}
			issue.SetIssue("todo", &v[2])
			issue.SubmitIssue()
		}
	}
}