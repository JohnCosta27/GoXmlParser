package lexer

const (
  // "<"
  LEFT_BRACKET string = "<"
  // ">"
  RIGHT_BRACKET string = ">"
  // "</"
  LEFT_AND_SLASH string = "</"
  // any string
  TEXT string = "TEXT"
)

type Token struct {
  // Do be used with enums above.
  Token string
  Token_content string
}

func GetTextToken() Token {
  return Token{
    Token: TEXT,
    Token_content: "",
  }
}
