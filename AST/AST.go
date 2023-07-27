package AST

import (
	"johncosta.tech/xmlparse/lexer"
)

const (
  ELEMENT_SUFFIX_CLOSE string = "element-suffix-close"
  ELEMENT_SUFFIX_OPEN string = "element-suffix-open"

  ATTRIBUTE_ELEMENT string = "attribute-element"
  
  CONTENT_ELEMENT string = "content-element"
  CONTENT_DATA string = "content-data"

  EPSILLON string = "epsillon"
)

type ASTNode interface {
  Print() string
  Walk(func (node ASTNode))
}

type Element struct {
  OpenTag *OpenTag
  ElementSuffix *ElementSuffix
}

type OpenTag struct {
  Type string

  NAME lexer.Token
  Attribute *Attribute
}

type CloseTag struct {
  NAME lexer.Token
}

type Attribute struct {
  // Do be used with the elements above
  Type string

  NAME lexer.Token
  STRING lexer.Token

  Attribute *Attribute
}

type ElementSuffix struct {
  Type string

  Content *Content
  CloseTag *CloseTag
}

type Content struct {
  // Do be used with the elements above
  Type string

  Element *Element
  DATA lexer.Token

  Content *Content
}

func (tag *Element) Print() string {
  startString := ""
  startString += "( Tag "
  startString += tag.OpenTag.Print()
  startString += tag.ElementSuffix.Print()
  startString += ") "
  return startString
}

func (tag *Element) Walk(callback func (node ASTNode)) {
  callback(tag)

  tag.OpenTag.Walk(callback)
  tag.ElementSuffix.Walk(callback)
}

func (attribute *Attribute) Print() string {
  startString := ""
  if (attribute.Type == EPSILLON) {
    startString += "( ATTRIBUTE EPSILLON ) "
  } else {
    startString += "( ATTRIBUTE ELEMENT "
    startString += "( NAME ( " + attribute.NAME.TokenContent + " ) "
    startString += "STRING ( " + attribute.STRING.TokenContent + " ) "
    startString += attribute.Attribute.Print()
    startString += ") "
  }
  
  return startString
}

func (attribute *Attribute) Walk(callback func(node ASTNode)) {
  callback(attribute)

  if (attribute.Type == EPSILLON) {
    return
  }

  attribute.Attribute.Walk(callback)
}

func (suffix *ElementSuffix) Print() string {
  if (suffix.Type == ELEMENT_SUFFIX_OPEN) {
    startString := "( ELEMTN_SUFFIX_OPEN "
    startString += suffix.Content.Print()
    startString += suffix.CloseTag.Print()
    startString += ") "
    return startString
  }

  return "( ELEMENT_SUFFIX_CLOSE ) "
}

func (suffix *ElementSuffix) Walk(callback func (node ASTNode)) {
  callback(suffix)

  if (suffix.Type == ELEMENT_SUFFIX_OPEN) {
    suffix.Content.Walk(callback)
    suffix.CloseTag.Walk(callback)
  } else {
    // Nothing to do
  }
}

func (content *Content) Walk(callback func(node ASTNode)) {
  callback(content)

  if (content.Type == CONTENT_ELEMENT) {
    content.Element.Walk(callback)
    content.Content.Walk(callback)
  } else if (content.Type == CONTENT_DATA) {
    content.Content.Walk(callback)
  }
}

func (content *Content) Print() string {
  startString := ""

  if (content.Type == EPSILLON) {
    startString += "( CONTENT EPSILLON ) "
  } else if (content.Type == CONTENT_ELEMENT) {
    startString += "( CONTENT ELEMENT "
    startString += content.Element.Print()
    startString += content.Content.Print()
    startString += ") "
  } else if (content.Type == CONTENT_DATA) {
    startString += "( CONTENT ELEMENT "
    startString += "( DATA ( "+ content.DATA.TokenContent + " ) "
    startString += content.Content.Print()
    startString += ") "
  }

  return startString
}

func (openTag *OpenTag) Print() string {
  startString := "( OPEN TAG "
  startString += "( NAME ( " + openTag.NAME.TokenContent + " ) "
  startString += openTag.Attribute.Print()
  startString += ") ) "
  return startString
}

func (openTag *OpenTag) Walk(callback func (node ASTNode)) {
  callback(openTag)
}

func (closeTag *CloseTag) Print() string {
  startString := "( CLOSE TAG "
  startString += "( NAME ( " + closeTag.NAME.TokenContent + " ) ) "
  startString += ") "
  return startString
}

func (closeTag *CloseTag) Walk(callback func (node ASTNode)) {
  callback(closeTag)
}
