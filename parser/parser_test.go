package parser

import (
	"log"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"johncosta.tech/xmlparse/lexer"
)

// ------------- INDIVIDUAL PARSES ------------- //

func TestSimple(t *testing.T) {
  log.Printf("%+v\n", lexer.Tokenize("<a>hello</a>"))
  ast, err := Parse(lexer.Tokenize("<a>hello</a>"))

  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print())
}

func TestNesting(t *testing.T) {
  ast, err := Parse(lexer.Tokenize("<a><b>Something</b></a>"))

  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print())
}
