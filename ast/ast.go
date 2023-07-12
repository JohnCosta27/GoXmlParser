package parser

type Tag struct {
  OpenTag OpenTag
  ChildTag *Tag
  CloseTag CloseTag
  SiblingTag *Tag
}

type OpenTag struct {
  TagName string
}

type CloseTag struct {
  TagName string
}

type Text struct {
  Text string
}
