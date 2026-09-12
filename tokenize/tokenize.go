package tokenize

import "strings"

func Tokenize(data []byte) []string {
	return strings.Fields(string(data))
}