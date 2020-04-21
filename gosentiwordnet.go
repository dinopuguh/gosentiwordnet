package go-sentiwordnet

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"

	"github.com/dinopuguh/go-sentiwordnet/data"
	"github.com/dinopuguh/go-sentiwordnet/helpers"
)

const sentiWordnetAssetName = "rawdata/SentiWordNet_3.0.0.txt"

type SentimentAnalyzer struct {
	Lexicon map[string]Sentiment
}

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}
}

func (sa *SentimentAnalyzer) GetSentimentScore(word string, posTag string, usage string) (bool, Sentiment) {
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

	return match, result
}

func NewGoSentiwordnet() *SentimentAnalyzer {
	var sa SentimentAnalyzer

	sa.generateLexicon()

	return &sa
}
