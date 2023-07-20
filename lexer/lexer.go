package lexer

import (
	"regexp"
)

/*
 * @note match important ones first
 */
func Tokenize(input string) TokenList {
  tokens := make([]Token, 0)
  tokenList := TokenList{
    Index: 0,
    Tokens: tokens,
  }

  // In order or priority.
  leftAndSlashRegex, _ := regexp.Compile("^</")
  leftBracketRegex, _ := regexp.Compile("^<")
  rightBracketRegex, _ := regexp.Compile("^>")
  equalRegex, _ := regexp.Compile("^=")
  stringLiteralRegex, _ := regexp.Compile(`^"[^"]*"`)
  whitespaceRegex, _ := regexp.Compile("^\\s+")
  anyCharacterRegex, _ := regexp.Compile(`^[^><"= ]+`)

  for (len(input) > 0) {
    {
      leftAndSlashMatch := leftAndSlashRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (leftAndSlashMatch != nil && leftAndSlashMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: LEFT_AND_SLASH,
        })
        input = input[leftAndSlashMatch[1]:]
        continue
      }
    }

    {
      leftBracketMatch := leftBracketRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (leftBracketMatch != nil && leftBracketMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: LEFT_BRACKET,
        })
        input = input[leftBracketMatch[1]:]
        continue
      }
    }

    {
      rightBracketMatch := rightBracketRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (rightBracketMatch != nil && rightBracketMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: RIGHT_BRACKET,
        })
        input = input[rightBracketMatch[1]:]
        continue
      }
    }

    {
      equalMatch := equalRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (equalMatch != nil && equalMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: EQUAL,
        })
        input = input[equalMatch[1]:]
        continue
      }
    }

    {
      stringLiteralMatch := stringLiteralRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (stringLiteralMatch != nil && stringLiteralMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: STRING_LITERAL,
          Token_content: input[1:stringLiteralMatch[1] - 1],
        })
        input = input[stringLiteralMatch[1]:]
        continue
      }
    }

    {
      whitespaceMatch := whitespaceRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (whitespaceMatch != nil && whitespaceMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: WHITESPACE,
          Token_content: input[0:whitespaceMatch[1]],
        })
        input = input[whitespaceMatch[1]:]
        continue
      }
    }

    {
      anyCharacterMatch := anyCharacterRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (anyCharacterMatch != nil && anyCharacterMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: TEXT,
          Token_content: input[0:anyCharacterMatch[1]],
        })
        input = input[anyCharacterMatch[1]:]
        continue
      }
    }

    panic("Could not match character!")
  }

  return tokenList
}
