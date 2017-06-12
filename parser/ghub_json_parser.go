package parser

import (
	"github.com/buger/jsonparser"
	"bytes"
)

// Retrieves the ghub commit ID and URL from webhook push payload
func ParseCommit(b *[]byte)(string, error) {
	// Get commit ID from JSON
	commitId, _, _, err := jsonparser.Get(
		*b,
		"commits",
		"id")

	// Get commit URL from JSON - expected format: https://api.github.com/repos/USER/PROJECT/commits{/sha}
	commitUrl, _, _, err := jsonparser.Get(
		*b,
		"repository",
		"commits_url")

	// Remove {/sha} from URL
	commitUrl = commitUrl[:len(commitUrl)-6]

	result := string(commitUrl) + "/" + string(commitId)
	return result, err
}

func ParsePatch(s *[]byte)(string, error) {
	// Retrieve patch content from JSON response
	patch, _, _, err := jsonparser.Get(
		*s,
		"files",
		"[0]",
		"patch")
	return patchExtract(patch), err
}

func patchExtract (p []byte) string{
	// Extract to do element(s)
	// TODO extra multiple todos
	extractLeft := "+\t// TODO"
	extractRight:= "\n"
	p = bytes.TrimLeft(p, extractLeft)
	return string(bytes.TrimRight(p,extractRight))
}