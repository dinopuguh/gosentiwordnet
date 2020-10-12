# 💬 GoSentiwordnet

[![Build Status](https://travis-ci.com/dinopuguh/gosentiwordnet.svg?branch=master)](https://travis-ci.com/dinopuguh/gosentiwordnet) [![Go Report Card](https://goreportcard.com/badge/github.com/dinopuguh/gosentiwordnet)](https://goreportcard.com/report/github.com/dinopuguh/gosentiwordnet) [![codecov](https://codecov.io/gh/dinopuguh/gosentiwordnet/branch/master/graph/badge.svg)](https://codecov.io/gh/dinopuguh/gosentiwordnet)  

Sentiment analyzer using [sentiwordnet](https://github.com/aesuli/SentiWordNet) lexicon in Go. This library produce sentiment score for each word, including positive, negative, and objective score. 

## ⚙ Installation

First of all, [download](https://golang.org/dl/) and install Go `1.14` or higher is required.

Install this library using the [`go get`](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

```bash
$ go get github.com/dinopuguh/gosentiwordnet
```



## ⚡ Quickstart

```go
package main

import (
    "fmt"
    
    goswn "github.com/dinopuguh/gosentiwordnet"
)

func main() {
    sa := goswn.New()
    
    scores, exist := sa.GetSentimentScore("love", "v", "2")
    if exist {
        fmt.Println("💬 Sentiment score:", scores) // => 💬 Sentiment score: {1 0 0}
    }
}
```

The `GetSentimentScore` required 3 parameters(word, pos-tag, and word usage):

1. **Word**: the word want to process
2. **POS tag**: part-of-speech tag of the word
3. **Word usage**: 1 for most common usage and a higher number would indicate lesser common usages



## 👍 Contributing

If you want to say **thank you** and/or support the active development of `Gosentiwordnet`:

1. Add a [GitHub Star](https://github.com/dinopuguh/gosentiwordnet/stargazers) to the project.
2. Write a review or tutorial on [Medium](https://medium.com/), [Dev.to](https://dev.to/) or personal blog.
3. Be a part of our [sponsors](https://github.com/sponsors/dinopuguh) to support this project.



## 💻 Contributors

- Dino Puguh (initial works)

Open for any pull requests to develop this project.