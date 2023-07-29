package translator

import (
	"fmt"
)

type JSONObject struct {
  Map map[string]JSONValueTypes
}

type JSONStringValue struct {
  Value string
}

type JSONObjectValue struct {
  Value JSONObject
}

type JSONValueTypes interface {
  Print() string
}

func (s JSONStringValue) Print() string {
  return s.Value
}

func (o JSONObjectValue) Print() string {
  return ""
}

func (object *JSONObject) Print() string {
  returnString := ""
  for k, v := range object.Map {
    returnString += fmt.Sprintf(`"%s":"%s",%s`, k, v.Print(), "\n")
  }
  return returnString
}
