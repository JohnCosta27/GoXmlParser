package lexer

func Tokenize(input string) []Token {
  tokens := make([]Token, 0)

  currentToken := Token{}

  for (len(input) > 0) {

    switch string(input[0]) {
    case "<":
      if (currentToken.Token == TEXT) {
        tokens = append(tokens, currentToken)
        currentToken = GetTextToken()
      }

      currentToken.Token = LEFT_BRACKET
      tokens = append(tokens, currentToken)
    case ">":
      if (currentToken.Token == TEXT) {
        tokens = append(tokens, currentToken)
        currentToken = GetTextToken()
      }

      currentToken.Token = RIGHT_BRACKET
      tokens = append(tokens, currentToken)
    case "/":
      if (currentToken.Token == TEXT) {
        tokens = append(tokens, currentToken)
        currentToken = GetTextToken()
      }

      currentToken.Token = SLASH
      tokens = append(tokens, currentToken)
    default:
      if (currentToken.Token != TEXT) {
        currentToken = GetTextToken()
      }
      currentToken.Token_content += string(input[0])
    }
    input = input[1:]
  }

  if (currentToken.Token == TEXT) {
    tokens = append(tokens, currentToken)
  }
  return tokens
}
