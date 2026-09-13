package main

import (
	"emailSpam/parsing"
	"fmt"
)

type TokenMap map[string]int

func main() {
	SpamCount := 0 
	spam :=  TokenMap{}
	parsing.AddDir("./data/enron1/spam/",spam,&SpamCount)



	fmt.Println(SpamCount)
}