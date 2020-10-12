package gosentiwordnet_test

import (
	"testing"

	"github.com/dinopuguh/gosentiwordnet"
	"github.com/stretchr/testify/assert"
)

type SentimentTestCase struct {
	Word   string
	PosTag string
	Usage  string
	Scores gosentiwordnet.Sentiment
}

func generateTestCases() []SentimentTestCase {
	return []SentimentTestCase{
		SentimentTestCase{Word: "love", PosTag: "v", Usage: "2", Scores: gosentiwordnet.Sentiment{Positive: 1, Negative: 0, Objective: 0}},
		SentimentTestCase{Word: "neat", PosTag: "a", Usage: "4", Scores: gosentiwordnet.Sentiment{Positive: 0.625, Negative: 0, Objective: 0.375}},
		SentimentTestCase{Word: "overacting", PosTag: "n", Usage: "1", Scores: gosentiwordnet.Sentiment{Positive: 0, Negative: 0.875, Objective: 0.125}},
		SentimentTestCase{Word: "finely", PosTag: "r", Usage: "2", Scores: gosentiwordnet.Sentiment{Positive: 0.625, Negative: 0, Objective: 0.375}},
	}
}

func TestSentimentAnalysis(t *testing.T) {
	sa := gosentiwordnet.New()
	for _, testCase := range generateTestCases() {
		scores, match := sa.GetSentimentScore(testCase.Word, testCase.PosTag, testCase.Usage)
		if match {
			assert.Equalf(t, testCase.Scores, scores, testCase.Word)
		}
	}
}
