package lexer_test

import (
	"fmt"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"johncosta.tech/xmlparse/lexer"
)

func TestText(t *testing.T) {
  snaps.MatchSnapshot(t, lexer.Tokenize("Just Text"))
}

func TestBasicTokens(t *testing.T) {
  snaps.MatchSnapshot(t, lexer.Tokenize("<a>Tag</a>"))
}

func TestSpacedTokens(t *testing.T) {
  snaps.MatchSnapshot(t, lexer.Tokenize("<a>Hello    World          Bruh</a>"))
}

func TestNesting(t *testing.T) {
  snaps.MatchSnapshot(t, lexer.Tokenize("<hello><world>Something And Something</world></hello>"))
}

func TestMatchingPriority(t *testing.T) {
  tokens := lexer.Tokenize("</")

  if (len(tokens.Tokens) != 1) {
    t.Error("Should only have 1 token")
    t.FailNow()
  }

  if (tokens.Tokens[0].Token != lexer.LEFT_AND_SLASH) {
    t.Error(fmt.Sprintf("Should have matched LEFT_AND_SLASH, instead matched: %s", tokens.Tokens[0].Token))
    t.FailNow()
  }
  
}
