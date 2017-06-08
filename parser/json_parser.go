package parser

import "github.com/buger/jsonparser"

func PushParse(b []byte)(string) {
	result := jsonparser.Get(b, "person", "name", "fullName")
	return string(result)
}
