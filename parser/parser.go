package parser

import (
	"errors"

	"johncosta.tech/xmlparse/AST"
	"johncosta.tech/xmlparse/lexer"
)

/*
  Using RDP (Recusive Descent Parsing) to parse this
  simply grammar. I have taken the algirthm from
  RHUL Compilers Course I took in the 3rd year.

  This algorithm is meant to be generated by a computer (compiler compiler), but for this project I wrote it by hand.

  This is why there's a lot of redundant IF statements double checking the same clause.
*/

func Parse(tokenList *lexer.TokenList) bool {
  _, err := parseTag(tokenList)
  return err == nil && tokenList.Index == len(tokenList.Tokens)
}

// Note how we return true if none of the if statements were hit, this means that the tag could be null
func parseTag(tokenList *lexer.TokenList) (AST.Tag, error) {
  ast := AST.Tag{}
  if (!tokenList.HasNext()) {
    return ast, nil
  }

  // First set of OpenTag
  if (tokenList.Current().Token == lexer.LEFT_BRACKET) {
    ast.Type = AST.TagElement

    openTag, err := parseOpenTag(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.OpenTag = openTag

    tag, err := parseTag(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.ChildTag = &tag

    closeTag, err := parseCloseTag(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.CloseTag = closeTag

    tag, err = parseTag(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.SiblingTag = &tag

    return ast, nil
  } else if (tokenList.Current().Token == lexer.TEXT) {
    ast.Type = AST.TagText

    if (tokenList.Current().Token == lexer.TEXT) {
      tokenList.Index += 1
    } else {
      return ast, errors.New("TAG | Expected text")
    }

    tag, err := parseTag(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.SiblingTag = &tag

    return ast, nil

  }

  return ast, nil
}

func parseOpenTag(tokenList *lexer.TokenList) (AST.OpenTag, error) {
  ast := AST.OpenTag{}
  if (tokenList.Current().Token == lexer.LEFT_BRACKET) {

    if (tokenList.Current().Token == lexer.LEFT_BRACKET) {
      tokenList.Index += 1
    } else {
      return AST.OpenTag{}, errors.New("OPEN TAG | Expect left bracket")
    }

    text, err := parseText(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.TagName = text

    if (tokenList.Current().Token == lexer.RIGHT_BRACKET) {
      // Build AST
      tokenList.Index += 1
    } else {
      return ast, errors.New("OPEN TAG | Expected right bracket")
    }

    return ast, nil
  }
  return ast, errors.New("CLOSE TAG | Not allowed to be null")
}

func parseCloseTag(tokenList *lexer.TokenList) (AST.CloseTag, error) {
  ast := AST.CloseTag{}
  if (tokenList.Current().Token == lexer.LEFT_AND_SLASH) {

    if (tokenList.Current().Token == lexer.LEFT_AND_SLASH) {
      // Build AST
      tokenList.Index += 1
    } else {
      return ast, errors.New("CLOSE TAG | Expected left and slash")
    }

    text, err := parseText(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.TagName = text

    if (tokenList.Current().Token == lexer.RIGHT_BRACKET) {
      // Build AST
      tokenList.Index += 1
    } else {
      return ast, errors.New("CLOSE TAG | Expected right bracket")
    }

    return ast, nil
  }
  return ast, errors.New("CLOSE TAG | Not allowed to be null")
}

func parseText(tokenList *lexer.TokenList) (AST.Text, error) {
  ast := AST.Text{}
  if (tokenList.Current().Token == lexer.TEXT) {
    if (tokenList.Current().Token == lexer.TEXT) {
      // Build AST
      ast.Text = tokenList.Current().Token_content
      tokenList.Index += 1
    } else {
      return ast, errors.New("TEXT | Could not find text token")
    }
  }
  return ast, nil
}
