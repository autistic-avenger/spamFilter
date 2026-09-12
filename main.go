package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, err:= os.ReadFile("./data/enron1/ham/0113.2000-01-04.farmer.ham.txt")
	if err!=nil{
		panic(err)
	}
	tokens := strings.Fields(string(data))
	if len(tokens)==0{
		panic("Nothing to Tokenize!")
	}
	
	freq := map[string]int{}

	for _,tk := range tokens{
		freq[tk] += 1
	}

	var total int 
	
	for tk := range freq{
		total += freq[tk]
	}

	for tk := range freq{
		fmt.Printf("TOKEN: %-20s | FREQ:%d \n",tk,freq[tk]/total)
	}
}