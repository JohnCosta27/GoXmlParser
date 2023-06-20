package lexer

func Tokenize(input string) []Token {
  tokens := make([]Token, 0)

  currentToken := Token{
    token: TEXT,
    token_content: "",
  }

  for (len(input) > 0) {

    switch string(input[0]) {
    case "<":
      currentToken.token = LEFT_BRACKET
      tokens = append(tokens, currentToken)
      currentToken = GetTextToken()
    case ">":
      currentToken.token = RIGHT_BRACKET
      tokens = append(tokens, currentToken)
      currentToken = GetTextToken()
    case "/":
      currentToken.token = SLASH
      tokens = append(tokens, currentToken)
      currentToken = GetTextToken()
    default:
      currentToken.token_content += string(input[0])
    }
    input = input[1:]
  }

  return tokens
}
