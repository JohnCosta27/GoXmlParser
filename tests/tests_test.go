package tests

import (
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"johncosta.tech/xmlparse/lexer"
	"johncosta.tech/xmlparse/parser"
	"johncosta.tech/xmlparse/semantics"
	"johncosta.tech/xmlparse/translator"
)

func TestBigSample(t *testing.T) {

  file, err := os.ReadFile("big_xml.txt")
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  ast, err := parser.Parse(lexer.Tokenize(string(file)))
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  err = semantics.SemanticAnalysis(ast)
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast)

  json, err := translator.TranslateJson(string(file))
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, json.Print())
}
