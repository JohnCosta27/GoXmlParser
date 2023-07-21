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
  // String literal "Like this"
  STRING string = "string"
  // A contiguous string
  NAME string = "name"
  // Data is surrounded by >DATA<
  DATA string = "data"
)

type Token struct {
  // Do be used with enums above.
  Token string
  TokenContent string
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
