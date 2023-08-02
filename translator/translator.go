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

  returnObject := JSONObjectValue{
    Map: make(map[string]JSONValueTypes),
  }

  returnObject.Map[ast.OpenTag.NAME.TokenContent] = jsonObject

  return returnObject, nil
}

func translateElement(element *AST.Element) JSONValueTypes {
  
  c := element.ElementSuffix.Content

  // Base case: <a>Hello</a>
  // Note that this case catches when there is ONLY data inside.
  // Fails on: <a>Hello<b>World</b></a>
  if (c.Type == AST.CONTENT_DATA && c.Content.Type == AST.EPSILLON) {
    return JSONStringValue{
      Value: c.DATA.TokenContent,
    }
  }

  jsonValue := JSONObjectValue{
    Map: make(map[string]JSONValueTypes),
  }

  for (c.Type != AST.EPSILLON) {

    // ---- Logic for loose text elements ----
    if (c.Type == AST.CONTENT_DATA) {
      value, exists := jsonValue.Map["#text"]

      if (exists) {
        stringValue, ok := value.(JSONStringValue)
        if (!ok) {
          panic("Should always be string value")
        }
        stringValue.Value += c.DATA.TokenContent
        jsonValue.Map["#text"] = stringValue
      } else {
        jsonValue.Map["#text"] = JSONStringValue{
          Value: c.DATA.TokenContent,
        }
      }
    // -- End Logic for loose text elements --
    } else {
      value, exists := jsonValue.Map[c.Element.OpenTag.NAME.TokenContent]

      if (exists) {
        // Exists, turn check if array already, or turn into array

        arrayValue, isArray := value.(JSONArrayValue)
        if (isArray) {
          // Array, append 
          arrayValue.Array = append(arrayValue.Array, translateElement(c.Element))
          jsonValue.Map[c.Element.OpenTag.NAME.TokenContent] = arrayValue
        } else {
          // Not array, turn into array
          newArray := JSONArrayValue{
            Array: make([]JSONValueTypes, 0),
          }

          newArray.Array = append(newArray.Array, jsonValue.Map[c.Element.OpenTag.NAME.TokenContent], translateElement(c.Element))
          jsonValue.Map[c.Element.OpenTag.NAME.TokenContent] = newArray
        }

      } else {
        // Does not exist
        jsonValue.Map[c.Element.OpenTag.NAME.TokenContent] = translateElement(c.Element)
      }
    }

    c = c.Content
  }

  return jsonValue
}
