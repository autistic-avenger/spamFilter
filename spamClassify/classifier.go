package spamclassify

import (
	"emailSpam/tokenize"
	"math"
	"os"
)

func Classify(path string, spamMap map[string]int, spamCount int, hamMap map[string]int, hamCount int, totalCount int) (float64, float64, error) {

	data,err := os.ReadFile(path)
	if err!=nil{
		return 0,0,err
	}

	contentTokens := tokenize.Tokenize(data)

	dProbability := 0.0
	hamProbability := 0.0
	spamProbability := 0.0

	for _,tk := range contentTokens{
		if spamMap[tk] != 0 {
			spamProbability += math.Log(float64(spamMap[tk])/float64(spamCount))
		}
		if hamMap[tk] != 0 {
			hamProbability += math.Log(float64(hamMap[tk])/float64(hamCount))
		}
		
		ct := hamMap[tk] + spamMap[tk]
		if ct!=0{
			dProbability += math.Log(float64(ct)/float64(spamCount))
		}

	}
	defaultHam := float64(hamCount)/float64(totalCount)
	defaultSpam := float64(spamCount)/float64(totalCount)

	finalHam := hamProbability+ defaultHam - dProbability
	finalSpam := spamProbability +defaultSpam - dProbability
	return finalHam,finalSpam,nil
}