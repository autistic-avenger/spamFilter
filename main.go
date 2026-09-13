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

	for tk := range spam{
		SpamCount += spam[tk]
	}



	HamCount := 0 
	ham := TokenMap{}
	parsing.AddDir("./data/enron1/ham/",ham,&HamCount)
		for tk := range ham{
		HamCount += ham[tk]
	}

	TotalCount := HamCount+ SpamCount


	fmt.Println(HamCount)
	fmt.Println(SpamCount)
	fmt.Println(TotalCount)
}