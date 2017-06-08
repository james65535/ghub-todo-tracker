package parser

import "github.com/buger/jsonparser"

// Retrieves the ghub commit ID from webhook push payload
func GetCommitUrl(b []byte)(string, error) {
	commitId, _, _, err := jsonparser.Get(
		b,
		"commits",
		"id")
	commitUrl, _, _, err := jsonparser.Get(
		b,
		"repository",
		"commits_url")

	result := string(commitId) + string(commitUrl)
	return string(result), err
}
