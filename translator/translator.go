package translator

import (
	"johncosta.tech/xmlparse/AST"
	"johncosta.tech/xmlparse/lexer"
	"johncosta.tech/xmlparse/parser"
	"johncosta.tech/xmlparse/semantics"
)

func TranslateJson(xmlString string) (JSONObjectValue, error) {
  ast, err := parser.Parse(lexer.Tokenize(xmlString))
  if (err != nil) {
    return JSONObjectValue{}, err
  }

  err = semantics.SemanticAnalysis(ast)
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

  content := element.ElementSuffix.Content.Content
  for (content.Type == AST.CONTENT_ELEMENT) {
    v, ok := c.(JSONObjectValue)
    if (!ok) {
      panic("Should always be json object value here")
    }

    existingValue, exists := v.Map[content.Element.OpenTag.NAME.TokenContent]
    if (exists) {
      array := JSONArrayValue{
        Array: make([]JSONValueTypes, 0),
      }
      array.Array = append(array.Array, existingValue)
      array.Array = append(array.Array, translateContent(content.Element.ElementSuffix.Content))
      v.Map[content.Element.OpenTag.NAME.TokenContent] = array
    } else {
      v.Map[content.Element.OpenTag.NAME.TokenContent] = translateContent(content.Element.ElementSuffix.Content)
    }

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

