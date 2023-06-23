package main

import (
	"fmt"
	"os"
	"strings"

	"johncosta.tech/xmlparse/lexer"
	"johncosta.tech/xmlparse/parser"
)

func main() {

  data, err := os.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }

  processedData := strings.ReplaceAll(string(data), "\n", "")
  processedData = strings.ReplaceAll(processedData, " ", "")
  tokens := lexer.Tokenize(processedData)

  hasParsed := parser.Parse(tokens)

  /*
  for _, v := range tokens {
    fmt.Printf("Token: %+v\n", v)
  }
  */

  fmt.Printf("Has Parsed: %v\n", hasParsed)
}
