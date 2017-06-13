package parser

import (
	"github.com/buger/jsonparser"
	//"bytes"
	//"fmt"
	"regexp"
)

type issue struct {
	file, message string
}

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

func ParsePatch(s *[]byte)([][]string, error) {
	// Retrieve patch content from JSON response
	// TODO switch JSON parser to get all and create array of issue
	patch, _, _, err := jsonparser.Get(
		*s,
		"files",
		"[0]",
		"patch")
	// Extract to do element(s)
	// Enhanced regex for neutral, subtractions, and additions - for now we want the latter
	addsRegex := regexp.MustCompile(`(?m)\\n([ ,\+,\-])[\\t, , \/,]+TODO *([\w, ]*)`)
	match := addsRegex.FindAllStringSubmatch(string(patch), -1)
	return match, err
}