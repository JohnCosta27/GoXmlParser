package AST

import "fmt"

const (
  TagElement string = "tag_element"
  TagText string = "tag_text"
)

type Tag struct {
  // To be used with enums above
  Type string

  OpenTag *OpenTag
  ChildTag *Tag
  CloseTag *CloseTag
  
  Text *Text

  SiblingTag *Tag
}

type OpenTag struct {
  TagName *Text
}

type CloseTag struct {
  TagName *Text
}

type Text struct {
  Text string
}

func getIndentString(indent int) string {
  tabs := ""
  for i := 0; i < indent; i++ {
    tabs += "\t"
  }
  return tabs
}

func (tag *Tag) Print(indent int) {
  if (tag.Type == TagElement) {
    fmt.Printf("%sTAG ELEMENT\n", getIndentString(indent))
    tag.OpenTag.Print(indent + 1)
    tag.ChildTag.Print(indent + 1)
    tag.CloseTag.Print(indent + 1)
    tag.SiblingTag.Print(indent + 1)
  } else if (tag.Type == TagText) {
    fmt.Printf("%sTAG TEXT\n", getIndentString(indent))
    tag.Text.Print(indent + 1)
    tag.SiblingTag.Print(indent + 1)
  } else {
    fmt.Printf("%sTAG EPSILLON\n", getIndentString(indent))
  }
}

func (openTag *OpenTag) Print(indent int) {
  fmt.Printf("%sOPEN TAG\n", getIndentString(indent))
  openTag.TagName.Print(indent + 1)
}

func (closeTag *CloseTag) Print(indent int) {
  fmt.Printf("%sCLOSE TAG\n", getIndentString(indent))
  closeTag.TagName.Print(indent + 1)
}

func (text *Text) Print(indent int) {
  fmt.Printf("%sTEXT | %s\n", getIndentString(indent), text.Text)
}
