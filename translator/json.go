package translator

import (
	"fmt"
)

type JSONStringValue struct {
  Value string
}

type JSONObjectValue struct {
  Map map[string]JSONValueTypes
}

type JSONValueTypes interface {
  Print() string
}

func (s JSONStringValue) Print() string {
  return s.Value
}

func (object JSONObjectValue) Print() string {
  returnString := ""
  for k, v := range object.Map {
    returnString += fmt.Sprintf(`"%s":"%s",%s`, k, v.Print(), "\n")
  }
  return returnString
}
