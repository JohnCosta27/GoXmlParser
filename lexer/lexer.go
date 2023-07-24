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

  leftAndSlashRegex := regexp.MustCompile("</")
  leftBracketRegex := regexp.MustCompile("<")
  slashBracketRegex := regexp.MustCompile("/>")
  rightBracketRegex := regexp.MustCompile(">")

  equalRegex := regexp.MustCompile("=")
  dataRegex := regexp.MustCompile("[^<>]*<")

  stringLiteralRegex := regexp.MustCompile(`"[^"]*"`)
  nameRegex := regexp.MustCompile("[A-z]+")

  whitespaceRegex := regexp.MustCompile("[\\s\n]+")

  for (len(input) > 0) {

    {
      whitespaceMatch := whitespaceRegex.FindStringIndex(input)
      if (whitespaceMatch != nil && whitespaceMatch[0] == 0) {
        input = input[whitespaceMatch[1]:]
        continue
      }
    }

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
      slashBracketMatch := slashBracketRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (slashBracketMatch != nil && slashBracketMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: SLASH_AND_RIGHT,
        })
        input = input[slashBracketMatch[1]:]
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
      dataMatch := dataRegex.FindStringIndex(input)
      if (dataMatch != nil && dataMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: DATA,
          TokenContent: input[:dataMatch[1] - 1],
        })
        input = input[dataMatch[1] - 1:]
        continue
      }
    }

    {
      stringLiteralMatch := stringLiteralRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (stringLiteralMatch != nil && stringLiteralMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: STRING,
          TokenContent: input[1:stringLiteralMatch[1] - 1],
        })
        input = input[stringLiteralMatch[1]:]
        continue
      }
    }

    {
      nameMatch := nameRegex.FindStringIndex(input)
      // Found a match right at the beginning of the string.
      if (nameMatch != nil && nameMatch[0] == 0) {
        tokenList.Tokens = append(tokenList.Tokens, Token{
          Token: NAME,
          TokenContent: input[:nameMatch[1]],
        })
        input = input[nameMatch[1]:]
        continue
      }
    }

    panic("Could not match character! " + input)
  }

  return tokenList
}
