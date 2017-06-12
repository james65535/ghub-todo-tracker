package clients

import (
	"io/ioutil"
	"net/http"
	"github.com/james65535/ghub-todo-tracker/utils"
)
/*
func IssuesClient(s *string)(error) {
	// TODO Post info to ghub issues board
	return err
} */

func CommitsClient(s *string)([]byte, error) {
	resp, err := http.Get(*s)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	utils.WebLog("Client log: " + string(body))
	return body, err
}
