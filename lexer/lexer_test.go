package lexer_test

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"johncosta.tech/xmlparse/lexer"
)

func TestBasicTokens(t *testing.T) {
  snaps.MatchSnapshot(t, lexer.Tokenize("<a>Tag</a>"))
}
