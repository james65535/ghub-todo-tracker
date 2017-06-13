package clients

import (
	"io/ioutil"
	"net/http"
	"github.com/james65535/ghub-todo-tracker/util"
	"encoding/json"
	"bytes"
	"fmt"
)

type GhubIssue struct {
	Title string `json:"title"`
	Body *string `json:"body"`
	Assignee string `json:"assignee,omitempty"`
	Milestone int `json:"milestone,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

/*
func AuthClient (username, password string) error {

}
*/

func (ghI *GhubIssue)SetIssue(title string, body *string){
	ghI.Title = "todo"
	ghI.Body = body
}

func (ghI *GhubIssue)SubmitIssue()(error) {

	body, err := json.Marshal(*ghI)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Issue JSON: %v\n", string(body))
	// TODO remove personal link
	url := "https://api.github.com/repos/james65535/ghub-todo-tracker/issues"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	// TODO remove token
	request.Header.Set("Authorization", "token d5adfeed8896e9a852d69d318c779695d678a6ed")

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

// TODO this may have to go elsewhere
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
