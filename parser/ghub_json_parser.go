package parser

import (
	"github.com/buger/jsonparser"
	//"bytes"
	//"fmt"
	"regexp"
	"fmt"
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

	return patchExtract(patch), err
}

func patchExtract (p []byte) [][]string {
	// Extract to do element(s)
	fmt.Printf("original body: %v\n", string(p))
	/* Enhanced regex for neutral, addition, and subtractions*/
	addsRegex := regexp.MustCompile(`(?m)\\n([ ,\+,\-])[\\t, , \/,]+TODO *([\w, ]*)`)

	ret := addsRegex.FindAllStringSubmatch(string(p), -1)
	return ret
}