package AST

const (
  TagElement string = "tag_element"
  TagText string = "tag_text"
)

type Tag struct {
  // To be used with enums above
  Type string

  OpenTag OpenTag
  ChildTag *Tag
  CloseTag CloseTag
  
  Text Text

  SiblingTag *Tag
}

type OpenTag struct {
  TagName Text
}

type CloseTag struct {
  TagName Text
}

type Text struct {
  Text string
}
