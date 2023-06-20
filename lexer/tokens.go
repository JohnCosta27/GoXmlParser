package lexer

const (
  // "<"
  LEFT_BRACKET int = 0
  // ">"
  RIGHT_BRACKET int = 1
  // "/"
  SLASH int = 2
  // any string
  TEXT int = 3
)

type Token struct {
  // Do be used with enums above.
  token int
  token_content string
}

func GetTextToken() Token {
  return Token{
    token: TEXT,
    token_content: "",
  }
}
