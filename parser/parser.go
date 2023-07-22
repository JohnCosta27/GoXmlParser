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

func Parse(tokenList lexer.TokenList) (*AST.Element, error) {
  tag, err := parseElement(&tokenList)

  if (err == nil && tokenList.Index == len(tokenList.Tokens)) {
    return tag, nil
  }
  return nil, err
}

func parseElement(tokenList *lexer.TokenList) (*AST.Element, error) {
  ast := &AST.Element{}
  if (!tokenList.HasNext()) {
    return ast, nil
  }

  // First set of OpenTag
  if (tokenList.Current().Token == lexer.LEFT_BRACKET) {
    openTag, err := parseOpenTag(tokenList)
    if (err != nil) {
      return ast, err 
    }
    ast.OpenTag = openTag

    content, err := parseContent(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.Content = content

    closeTag, err := parseCloseTag(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.CloseTag = closeTag

    return ast, nil
  }

  return ast, errors.New("ELEMENT | Cannot be null")
}

func parseContent(tokenList *lexer.TokenList) (*AST.Content, error) {
  ast := &AST.Content{}
  if (tokenList.Current().Token == lexer.LEFT_BRACKET) {
    ast.Type = AST.CONTENT_ELEMENT

    element, err := parseElement(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.Element = element

    content, err := parseContent(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.Content = content

    return ast, nil
  } else if (tokenList.Current().Token == lexer.DATA) {
    ast.Type = AST.CONTENT_DATA

    if (tokenList.Current().Token == lexer.DATA) {
      ast.DATA = tokenList.Current()
      tokenList.Index += 1
    } else {
      return ast, errors.New("CONTENT DATA | Expect data")
    }

    content, err := parseContent(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.Content = content

    return ast, nil
  }

  ast.Type = AST.EPSILLON
  return ast, nil
}

func parseOpenTag(tokenList *lexer.TokenList) (*AST.OpenTag, error) {
  ast := &AST.OpenTag{}
  if (tokenList.Current().Token == lexer.LEFT_BRACKET) {
    if (tokenList.Current().Token == lexer.LEFT_BRACKET) {
      tokenList.Index += 1
    } else {
      return ast, errors.New("OPEN TAG | Expect left bracket")
    }

    if (tokenList.Current().Token == lexer.NAME) {
      ast.NAME = tokenList.Current()
      tokenList.Index += 1
    } else {
      return ast, errors.New("OPEN TAG | Expect name")
    }

    attribute, err := parseAttribute(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.Attribute = attribute

    if (tokenList.Current().Token == lexer.RIGHT_BRACKET) {
      tokenList.Index += 1
    } else {
      return ast, errors.New("OPEN TAG | Expected right bracket")
    }

    return ast, nil
  }
  return ast, errors.New("OPEN TAG | Not allowed to be null")
}

func parseAttribute(tokenList *lexer.TokenList) (*AST.Attribute, error) {
  ast := &AST.Attribute{}
  if (tokenList.Current().Token == lexer.NAME) {
    ast.Type = AST.ATTRIBUTE_ELEMENT

    if (tokenList.Current().Token == lexer.NAME) {
      ast.NAME = tokenList.Current()
      tokenList.Index += 1
    } else {
      return ast, errors.New("ATTRIBUTE | Expected name")
    }

    if (tokenList.Current().Token == lexer.EQUAL) {
      tokenList.Index += 1
    } else {
      return ast, errors.New("ATTRIBUTE | Expected equals")
    }

    if (tokenList.Current().Token == lexer.STRING) {
      tokenList.Index += 1
      ast.STRING = tokenList.Current()
    } else {
      return ast, errors.New("ATTRIBUTE | Expected equals")
    }

    attribute, err := parseAttribute(tokenList)
    if (err != nil) {
      return ast, err
    }
    ast.Attribute = attribute

    return ast, nil
  }

  ast.Type = AST.EPSILLON
  return ast, nil
}

func parseCloseTag(tokenList *lexer.TokenList) (*AST.CloseTag, error) {
  ast := &AST.CloseTag{}
  if (tokenList.Current().Token == lexer.LEFT_AND_SLASH) {
    if (tokenList.Current().Token == lexer.LEFT_AND_SLASH) {
      ast.NAME = tokenList.Current()
      tokenList.Index += 1
    } else {
      return ast, errors.New("CLOSE TAG | Expected left and slash")
    }

    token := tokenList.Current()
    if (token.Token == lexer.NAME) {
      tokenList.Index += 1
      ast.NAME = token
    } else {
      return ast, errors.New("CLOSE TAG | Expected left and slash")
    }

    if (tokenList.Current().Token == lexer.RIGHT_BRACKET) {
      tokenList.Index += 1
    } else {
      return ast, errors.New("CLOSE TAG | Expected right bracket")
    }

    return ast, nil
  }
  return ast, errors.New("CLOSE TAG | Not allowed to be null")
}
