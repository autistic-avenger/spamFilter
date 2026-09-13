package tokenize

import "strings"

func Tokenize(data []byte) []string {
	stData := strings.Fields(string(data))
	var returnData []string
	for _,st := range stData{
		returnData = append(returnData, strings.ToUpper(st))
	}
	return returnData
}