package clients

import (
	"io/ioutil"
	"net/http"
	"github.com/james65535/ghub-todo-tracker/util"
	"encoding/json"
	"bytes"
)

func AuthClient (username, password string) error {

}

func IssuesClient(s *string)(error) {

	type ghubIssues struct {
		title, body, assignee string
		milestone int
		labels, assignees []string
	}

	var issue = ghubIssues{"todo", *s, "", 0, nil, nil}
	body, err := json.Marshal(issue)
	if err != nil {
		panic(err)
	}
	url := "https://api.github.com/repos/james65535/ghub-todo-tracker/issues"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	util.WebLog("Issues Client log: " + string(respBody))

	// TODO Post info to ghub issues board
	return err
}

func CommitsClient(s *string)([]byte, error) {
	resp, err := http.Get(*s)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	util.WebLog("Commit Client log: " + string(body))
	return body, err
}
