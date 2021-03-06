package acousticid
import (

	"strings"
)

func FuzzyMatchString(source string, target string) float64 {

	//check how many words of the source (acousticid) is available in the target (file name)
	sourcewords := strings.Fields(source)
	targetwords := strings.Fields(target)

	//Highly likely the words would be sorted, so the efficiency of the double loop should be ok-ish
	count := 0
	for _, word := range sourcewords {
		for _, word2 := range targetwords {
			if strings.EqualFold(word, word2) {
				count ++
				break
			}
		}
	}

	//calculate percentage match
	return float64(count) / float64(len(sourcewords))
}
