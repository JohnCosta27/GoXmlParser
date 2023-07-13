package lexer

func Tokenize(input string) TokenList {
  tokens := make([]Token, 0)
  tokenList := TokenList{
    Index: 0,
    Tokens: tokens,
  }

  currentToken := Token{}

  for (len(input) > 0) {

    switch string(input[0]) {
    case "<":
      if (currentToken.Token == TEXT) {
        tokenList.Tokens = append(tokenList.Tokens, currentToken)
        currentToken = GetTextToken()
      }

      if (len(input) > 1 && string(input[1]) == "/") {
        input = input[1:]
        currentToken.Token = LEFT_AND_SLASH
        tokenList.Tokens = append(tokenList.Tokens, currentToken)
      } else {
        currentToken.Token = LEFT_BRACKET
        tokenList.Tokens = append(tokenList.Tokens, currentToken)
      }
    case ">":
      if (currentToken.Token == TEXT) {
        tokenList.Tokens = append(tokenList.Tokens, currentToken)
        currentToken = GetTextToken()
      }

      currentToken.Token = RIGHT_BRACKET
      tokenList.Tokens = append(tokenList.Tokens, currentToken)
    default:
      if (currentToken.Token != TEXT) {
        currentToken = GetTextToken()
      }
      currentToken.Token_content += string(input[0])
    }
    input = input[1:]
  }

  if (currentToken.Token == TEXT) {
    tokenList.Tokens = append(tokenList.Tokens, currentToken)
  }
  return tokenList
}
