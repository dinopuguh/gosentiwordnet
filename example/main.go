package main

import (
	"fmt"

	goswn "github.com/dinopuguh/gosentiwordnet"
)

// Token represent the required parameter for using gosentiwordnet
type Token struct {
	Word   string // the word want to process
	PosTag string // part-of-speech tag of word
	Usage  string // word usage (1 for most common usage and a higher number would indicate lesser common usages)
}

func main() {
	sa := goswn.New()

	tokens := []Token{
		{Word: "love", PosTag: "v", Usage: "2"},
		{Word: "neat", PosTag: "a", Usage: "4"},
		{Word: "overacting", PosTag: "n", Usage: "1"},
	}

	for _, token := range tokens {
		scores, exist := sa.GetSentimentScore(token.Word, token.PosTag, token.Usage)
		if exist {
			fmt.Printf("💬 Sentiment score of %s: %v\n", token.Word, scores)
			// 💬 Sentiment score: {positive_score negative_score objective_score}
		}
	}
}
