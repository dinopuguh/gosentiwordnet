package gosentiwordnet

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"

	"github.com/dinopuguh/gosentiwordnet/data"
	"github.com/dinopuguh/gosentiwordnet/helpers"
)

const sentiWordnetAssetName = "rawdata/SentiWordNet_3.0.0.txt"

// SentimentAnalyzer represent the sentiment analyzer with sentiwordnet lexicon
type SentimentAnalyzer struct {
	Lexicon map[string]Sentiment
}

// Sentiment reprensent sentiment score for each word
// containing positive, negative and objective
type Sentiment struct {
	Positive  float64
	Negative  float64
	Objective float64
}

func (sa *SentimentAnalyzer) generateLexicon() {
	sa.Lexicon = make(map[string]Sentiment)

	asset, err := data.Asset(sentiWordnetAssetName)
	if err != nil {
		log.Panic(err.Error())
	}

	file := bytes.NewReader(asset)
	scanner := bufio.NewScanner(file)

	line := 1
	for scanner.Scan() {
		if line > 26 {
			thisRawLine := scanner.Text()
			thisSplitLine := strings.Split(thisRawLine, "\t")

			posTag := thisSplitLine[0]
			positive, _ := strconv.ParseFloat(thisSplitLine[2], 64)
			negative, _ := strconv.ParseFloat(thisSplitLine[3], 64)
			synsetTerms := thisSplitLine[4]

			key := "(" + posTag + ") " + synsetTerms + " "
			var sentiment Sentiment
			sentiment.Positive = positive
			sentiment.Negative = negative
			sentiment.Objective = 1 - (positive + negative)

			sa.Lexicon[key] = sentiment
		}

		line++
	}
}

// GetSentimentScore count the sentiment score of word based on POS tag and the word usage.
// POS tag: part-of-speech tag of word
// Word usage: 1 for most common usage and a higher number would indicate lesser common usages
func (sa *SentimentAnalyzer) GetSentimentScore(word string, posTag string, usage string) (Sentiment, bool) {
	var result Sentiment

	posTag = "(" + posTag + ")"
	synset := " " + word + "#" + usage + " "

	match := false

	for key, value := range sa.Lexicon {
		if helpers.CheckSubstrings(key, posTag, synset) {
			result = value
			match = true
			break
		}
	}

	return result, match
}

// New generate the sentiment analyzer with sentiwordnet lexicon
func New() *SentimentAnalyzer {
	var sa SentimentAnalyzer

	sa.generateLexicon()

	return &sa
}
