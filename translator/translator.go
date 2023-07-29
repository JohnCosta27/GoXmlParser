package translator

import (
	"log"

	"johncosta.tech/xmlparse/AST"
	"johncosta.tech/xmlparse/lexer"
	"johncosta.tech/xmlparse/parser"
)

func TranslateJson(xmlString string) (JSONObjectValue, error) {
  ast, err := parser.Parse(lexer.Tokenize(xmlString))
  if (err != nil) {
    return JSONObjectValue{}, err
  }

  jsonObject := translateElement(ast)

  return jsonObject, nil
}

func translateElement(element *AST.Element) JSONObjectValue {
  jsonObject := JSONObjectValue{
    Map: make(map[string]JSONValueTypes),
  }

  c := translateContent(element.ElementSuffix.Content)
  jsonObject.Map[element.OpenTag.NAME.TokenContent] = c

  if (element.ElementSuffix.Content.Type != AST.CONTENT_ELEMENT) {
    return jsonObject
  }

  content := element.ElementSuffix.Content
  for (content.Type == AST.CONTENT_ELEMENT) {
    v, ok := c.(JSONObjectValue)
    if (!ok) {
      panic("Should always be json object value here")
    }
    v.Map[content.Element.OpenTag.NAME.TokenContent] = translateContent(content.Element.ElementSuffix.Content)
    content = content.Content
  }

  return jsonObject
}

func translateContent(content *AST.Content) JSONValueTypes {
  if (content.Type == AST.CONTENT_DATA) {
    return JSONStringValue{
      Value: content.DATA.TokenContent,
    }
  } else if (content.Type == AST.CONTENT_ELEMENT) {
    return translateElement(content.Element)
  }
  return JSONStringValue{
    Value: "",
  }
}

