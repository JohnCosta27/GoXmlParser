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

func TestEquals(t *testing.T) {
  tokens := lexer.Tokenize("=")
  if (len(tokens.Tokens) != 1) {
    t.Error(fmt.Sprintf("Expected tokens to have length of 1, not %d", len(tokens.Tokens)))
  }

  if (tokens.Tokens[0].Token != lexer.EQUAL) {
    t.Error(fmt.Sprintf("Should have matched EQUAL, instead matched: %s", tokens.Tokens[0].Token))
    t.FailNow()
  }
}

func TestStringLiteral(t *testing.T) {
  tokens := lexer.Tokenize(`"hello"`)
  if (len(tokens.Tokens) != 1) {
    t.Error(fmt.Sprintf("Expected tokens to have length of 1, not %d", len(tokens.Tokens)))
  }

  if (tokens.Tokens[0].Token != lexer.QUOTE) {
    t.Error(fmt.Sprintf("Should have matched STRING_LITERAL, instead matched: %s", tokens.Tokens[0].Token))
    t.FailNow()
  }

  if (tokens.Tokens[0].Token_content != "hello") {
    t.Error(fmt.Sprintf("Should have found 'hello' as token content, instead got: %s", tokens.Tokens[0].Token_content))
    t.FailNow()
  }
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

func TestMultipleLines(t *testing.T) {
  snaps.MatchSnapshot(t, lexer.Tokenize(`
    <hello>
      <world>Something here <a> </a> ??????          




    dsnmkadsmakkmlsdakmldsa

      </world></hello>
    `))
}

func TestAttributes(t *testing.T) {
  snaps.MatchSnapshot(t, lexer.Tokenize(`<a hello="world">`))
}
