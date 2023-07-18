package main

import (
	"os"
	"strings"

	"johncosta.tech/xmlparse/lexer"
	"johncosta.tech/xmlparse/parser"
	"johncosta.tech/xmlparse/semantics"
)

func main() {

  data, err := os.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }

  processedData := strings.ReplaceAll(string(data), "\n", "")
  processedData = strings.ReplaceAll(processedData, " ", "")
  tokenList := lexer.Tokenize(processedData)

  ast, err := parser.Parse(&tokenList)
  if (err != nil) {
    panic(err)
  }

  semantics.SemanticAnalysis(ast)
}
