package main

import (
	"emailSpam/parsing"
	spamclassify "emailSpam/spamClassify"
	"fmt"
)

type TokenMap map[string]int

func main() {
	fmt.Println("EMAIL SPAM DETECTION USING NAIVE BAYES THEOREM!")

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



	
	testPath := "./test.txt"

	hamProb ,spamProp ,err := spamclassify.Classify(testPath,spam,SpamCount,ham,HamCount,TotalCount)
	if err!=nil{
		fmt.Println("Error Classifying")
		panic(err)
	}

	if hamProb>spamProp{
		fmt.Println("\n\n\nIt's likely HAM !")
	}else{
		fmt.Println("\n\n\nIt's likely SPAM !")
	}
}