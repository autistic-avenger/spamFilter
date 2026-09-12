package main

import (
	"emailSpam/parsing"
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	tokenCount := 1 
	mainTokenMap :=  map[string]int{}
	filepath.WalkDir("./data/enron1",func(path string, d fs.DirEntry, err error) error {
		if d.IsDir(){
			return nil
		}
		err = parsing.AddTokens(path,mainTokenMap)
		if err!=nil{
			panic(err)
		}
		for tk,cont := range mainTokenMap{
			fmt.Printf("TOKEN :%-15s | COUNT: %4d | TOKENCOUNT: %7d \n",tk,cont,tokenCount)
			tokenCount++
		}
		return nil
	})
}