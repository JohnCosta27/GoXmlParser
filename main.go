package main

import (
	"fmt"
	"os"
	"strings"

	"johncosta.tech/xmlparse/lexer"
)

func main() {

  data, err := os.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }

  processedData := strings.ReplaceAll(string(data), "\n", "")
  tokens := lexer.Tokenize(processedData)

  fmt.Printf("Tokens: %+v\n", tokens)

}
