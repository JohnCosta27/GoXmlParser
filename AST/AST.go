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
  Print(indent int) string
  Walk(func (node ASTNode), func (node ASTNode))
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

func getIndentString(indent int) string {
  indents := ""
  for i := 0; i < indent; i++ {
    indents += " "
  }
  return indents
}

func (tag *Element) Print(indent int) string {
  indentString := getIndentString(indent)
  startString := indentString + "Tag\n"
  startString += tag.OpenTag.Print(indent + 1)
  startString += tag.ElementSuffix.Print(indent + 1)
  return startString
}

func (tag *Element) Walk(enter func (node ASTNode), exit func (node ASTNode)) {
  enter(tag)

  tag.OpenTag.Walk(enter, exit)
  tag.ElementSuffix.Walk(enter, exit)

  exit(tag)
}

func (attribute *Attribute) Print(indent int) string {
  if (attribute.Type == ATTRIBUTE_ELEMENT) {
    indentString := getIndentString(indent)
    startString := indentString + "Attributes\n"
    startString += indentString + "Name: " + attribute.NAME.TokenContent + "\n"
    startString += indentString + "String: " + attribute.STRING.TokenContent + "\n"
    startString += attribute.Attribute.Print(indent + 1)
    return startString
  }
  
  return ""
}

func (attribute *Attribute) Walk(enter func (node ASTNode), exit func (node ASTNode)) {
  enter(attribute)

  if (attribute.Type != EPSILLON) {
    attribute.Attribute.Walk(enter, exit)
  }

  exit(attribute)
}

func (suffix *ElementSuffix) Print(indent int) string {
  indentString := getIndentString(indent)

  if (suffix.Type == ELEMENT_SUFFIX_OPEN) {
    startString := indentString + "ElementSuffixOpen\n"
    startString += suffix.Content.Print(indent + 1)
    startString += suffix.CloseTag.Print(indent + 1)
    return startString
  }

  return indentString + "ElementSuffixClose\n"
}

func (suffix *ElementSuffix) Walk(enter func (node ASTNode), exit func (node ASTNode)) {
  enter(suffix)

  if (suffix.Type == ELEMENT_SUFFIX_OPEN) {
    suffix.Content.Walk(enter, exit)
    suffix.CloseTag.Walk(enter, exit)
  } else {
    // Nothing to do
  }

  exit(suffix)
}

func (content *Content) Walk(enter func (node ASTNode), exit func (node ASTNode)) {
  enter(content)

  if (content.Type == CONTENT_ELEMENT) {
    content.Element.Walk(enter, exit)
    content.Content.Walk(enter, exit)
  } else if (content.Type == CONTENT_DATA) {
    content.Content.Walk(enter, exit)
  }

  exit(content)
}

func (content *Content) Print(indent int) string {
  indentString := getIndentString(indent)

  if (content.Type == CONTENT_ELEMENT) {
    startString := indentString + "ContentElement\n"
    startString += content.Element.Print(indent + 1)
    startString += content.Content.Print(indent + 1)
    return startString
  } else if (content.Type == CONTENT_DATA) {
    startString := indentString + "ContentData\n"
    startString += indentString + "Data: " + content.DATA.TokenContent + "\n"
    startString += content.Content.Print(indent + 1)
    return startString
  }

  return ""
}

func (openTag *OpenTag) Print(indent int) string {
  indentString := getIndentString(indent)
  startString := indentString + "OpenTag\n"
  startString += indentString + "Name: " + openTag.NAME.TokenContent + "\n"
  startString += openTag.Attribute.Print(indent + 1)
  return startString
}

func (openTag *OpenTag) Walk(enter func (node ASTNode), exit func (node ASTNode)) {
  enter(openTag)
  exit(openTag)
}

func (closeTag *CloseTag) Print(indent int) string {
  indentString := getIndentString(indent)
  startString := indentString + "CloseTag\n"
  startString += indentString + "Name: " + closeTag.NAME.TokenContent + "\n"
  return startString
}

func (closeTag *CloseTag) Walk(enter func (node ASTNode), exit func (node ASTNode)) {
  enter(closeTag)
  exit(closeTag)
}
