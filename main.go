package main

import (
	"fmt"
	"os"

	"johncosta.tech/xmlparse/lexer"
	"johncosta.tech/xmlparse/parser"
	"johncosta.tech/xmlparse/semantics"
)

func main() {

  data, err := os.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }

  processedData := string(data);

  tokenList := lexer.Tokenize(processedData)

  ast, err := parser.Parse(&tokenList)
  if (err != nil) {
    panic(err)
  }

  err = semantics.SemanticAnalysis(ast)
  if (err != nil) {
    panic(err)
  }

  fmt.Println("Successfully parsed and semantically checked!")
}
