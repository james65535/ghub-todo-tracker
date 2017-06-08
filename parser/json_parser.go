package parser

import "github.com/buger/jsonparser"

func PushParse(b []byte)(string, error) {
	result, _, _, err := jsonparser.Get(
		b,
		"commits",
		"id")
	return string(result), err
}
