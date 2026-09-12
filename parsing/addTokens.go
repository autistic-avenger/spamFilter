package parsing

import (
	"emailSpam/tokenize"
	"os"
)

func AddTokens(path string, mainTokenMap map[string]int) error {
	data ,err:= os.ReadFile(path)
	if err!=nil{
		return err
	}

	tokens := tokenize.Tokenize(data)
	for _,tk := range tokens{
		mainTokenMap[tk]++
	}
	

	return nil
}