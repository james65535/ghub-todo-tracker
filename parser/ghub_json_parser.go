package parser

import "github.com/buger/jsonparser"

// Retrieves the ghub commit ID from webhook push payload
func ParseCommit(b *[]byte)(string, error) {
	commitId, _, _, err := jsonparser.Get(
		*b,
		"commits",
		"id")
	// expected format: https://api.github.com/repos/USER/PROJECT/commits{/sha}
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
	patch, _, _, err := jsonparser.Get(
		*s,
		"files",
		"0",
		"patch")

	// TODO parse patch statements to grab TODOs
	result := string(patch)
	return result, err
}
