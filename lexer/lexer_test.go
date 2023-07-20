package lexer_test

import (
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

// When you have <a>HEllo World</a>, the space should
// trigger another token to be created.
func TestSpacedTokens(t *testing.T) {
  snaps.MatchSnapshot(t, lexer.Tokenize("<a>Hello    World          Bruh</a>"))
}
