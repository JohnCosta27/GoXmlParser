package AST

const (
  ContentElement string = "content-element"
  ContentData string = "content-data"
)

type ASTNode interface {
  Print() string
  Walk(func (node ASTNode))
}

type Tag struct {
  OpenTag *OpenTag
  Content *Content
  CloseTag *CloseTag
}

type OpenTag struct {
  TagName *Text
}

type CloseTag struct {
  TagName *Text
}

type Content struct {
  // Do be used with the elements above
  Type string

  Element *Tag
  Data *Data
  Content *Content
}

// ------ Literals ------

type Data struct {
  Text string
}

type Name struct {
  Text string
}

type String struct {
  Text string
}

func (tag *Tag) Print() string {
  startString := ""
  if (tag.Type == TagElement) {
    startString += "( Tag Element "
    startString += tag.OpenTag.Print()
    startString += tag.ChildTag.Print()
    startString += tag.CloseTag.Print()
    startString += tag.SiblingTag.Print()
    startString += ") "
  } else if (tag.Type == TagText) {
    startString += "( Tag Text "
    startString += tag.Text.Print()
    startString += tag.SiblingTag.Print()
    startString += ") "
  } else {
    startString += "( TAG EPSILLON ) "
  }
  return startString
}

func (tag *Tag) Walk(callback func (node ASTNode)) {
  callback(tag)

  if (tag.Type == TagElement) {
    tag.OpenTag.Walk(callback)
    if (tag.ChildTag != nil) {
      tag.ChildTag.Walk(callback)
    }
    tag.CloseTag.Walk(callback)
    if (tag.SiblingTag != nil) {
      tag.SiblingTag.Walk(callback)
    }
  } else {
    tag.Text.Walk(callback)
    if (tag.SiblingTag != nil) {
      tag.SiblingTag.Walk(callback)
    }
  }
}

func (openTag *OpenTag) Print() string {
  str := "( OPEN TAG "
  str += openTag.TagName.Print()
  str += ") "
  return str
}

func (openTag *OpenTag) Walk(callback func (node ASTNode)) {
  callback(openTag)

  openTag.TagName.Walk(callback)
}

func (closeTag *CloseTag) Print() string {
  str := "( CLOSE TAG "
  str += closeTag.TagName.Print()
  str += ") "
  return str
}

func (closeTag *CloseTag) Walk(callback func (node ASTNode)) {
  callback(closeTag)

  closeTag.TagName.Walk(callback)
}

func (text *Text) Print() string {
  return "( TEXT ( " + text.Text + " ) "
}

func (text *Text) Walk(callback func (node ASTNode)) {
  callback(text)
}
