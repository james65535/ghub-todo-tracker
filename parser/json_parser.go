package parser

import "github.com/buger/jsonparser"

func JsonParse(b []byte)(string) {
	result := jsonparser.Get(b, "person", "name", "fullName")
	return string(result)
}
