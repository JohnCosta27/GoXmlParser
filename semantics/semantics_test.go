package semantics

import (
	"testing"

	"johncosta.tech/xmlparse/lexer"
	"johncosta.tech/xmlparse/parser"
)

func TestBasic(t *testing.T) {
  element, err := parser.Parse(lexer.Tokenize("<a>Hello</a>"))
  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  err = SemanticAnalysis(element)
  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }
}

func TestNesting(t *testing.T) {
  element, err := parser.Parse(lexer.Tokenize("<a><b>Hello</b></a>"))
  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  err = SemanticAnalysis(element)
  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }
}

func TestWrongClose(t *testing.T) {
  element, err := parser.Parse(lexer.Tokenize("<a>Hello</b>"))
  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  err = SemanticAnalysis(element)
  if (err == nil) {
    t.Error("Expected to fail")
    t.FailNow()
  }
}
