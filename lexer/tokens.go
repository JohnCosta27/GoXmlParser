package lexer

const (
  // "<"
  LEFT_BRACKET string = "<"
  // ">"
  RIGHT_BRACKET string = ">"
  // "</"
  LEFT_AND_SLASH string = "</"
  // "=" to be used for attributes.
  EQUAL string = "="
  // Whitespace characters (at least one, but could have more).
  WHITESPACE string = " "
  // any string
  TEXT string = "TEXT"
)

type Token struct {
  // Do be used with enums above.
  Token string
  Token_content string
}

type TokenList struct {
  Index int
  Tokens []Token
}

func (tl *TokenList) HasNext() bool {
  return len(tl.Tokens) > tl.Index
}

func (tl *TokenList) Current() Token {
  return tl.Tokens[tl.Index]
}
