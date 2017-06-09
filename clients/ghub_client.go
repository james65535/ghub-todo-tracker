package clients

import (
	"io/ioutil"
	"net/http"
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
	return body, err
}
