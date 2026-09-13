package main

import (
	"emailSpam/parsing"
	spamclassify "emailSpam/spamClassify"
	"fmt"
)

type TokenMap map[string]int

func main() {
	fmt.Printf("EMAIL SPAM DETECTION USING NAIVE BAYES THEOREM!\n\n\n")

	fmt.Printf("Training...")

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


	fmt.Print("\rTraining Complete!\n")

	testPath := "./test.txt"
	fmt.Println("Classifying :",testPath)

	hamProb ,spamProb ,err := spamclassify.Classify(testPath,spam,SpamCount,ham,HamCount,TotalCount)
	if err!=nil{
		fmt.Println("Error Classifying")
		panic(err)
	}

	if hamProb>spamProb{
		fmt.Println("\n\n\033[32mIt's likely HAM !\033[0m")
	}else{
		fmt.Println("\n\n\033[31mIt's likely SPAM !\033[0m")
	}
}