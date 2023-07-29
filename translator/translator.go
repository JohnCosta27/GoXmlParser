package translator

import (
	"log"

	"johncosta.tech/xmlparse/AST"
	"johncosta.tech/xmlparse/lexer"
	"johncosta.tech/xmlparse/parser"
)

func TranslateJson(xmlString string) (JSONObject, error) {
  ast, err := parser.Parse(lexer.Tokenize(xmlString))
  if (err != nil) {
    return JSONObject{}, err
  }

  jsonObject := translateElement(ast)

  return jsonObject, nil
}

func translateElement(element *AST.Element) JSONObject {
  jsonObject := JSONObject{
    Map: make(map[string]JSONValueTypes),
  }

  if (element.ElementSuffix.Content.Type == AST.CONTENT_DATA) {

    jsonObject.Map[element.OpenTag.NAME.TokenContent] = JSONStringValue{
      Value: element.ElementSuffix.Content.DATA.TokenContent,
    }

    log.Println(element.ElementSuffix.Content.Content.Type)

    if (element.ElementSuffix.Content.Content.Type == AST.CONTENT_ELEMENT) {
      jsonObject.Map[element.ElementSuffix.Content.Element.OpenTag.NAME.TokenContent] = JSONStringValue{
        Value: element.ElementSuffix.Content.Element.ElementSuffix.Content.DATA.TokenContent,
      }
    }

  } else if (element.ElementSuffix.Content.Type == AST.CONTENT_ELEMENT) {
    jsonObject.Map[element.OpenTag.NAME.TokenContent] = JSONObjectValue{
      Value: translateElement(element.ElementSuffix.Content.Element),
    }
  }

  return jsonObject
}
