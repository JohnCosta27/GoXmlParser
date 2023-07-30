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

type JSONArrayValue struct {
  Array []JSONValueTypes
}

type JSONValueTypes interface {
  Print() string
}

func (s JSONStringValue) Print() string {
  return fmt.Sprintf(`"%s"`, s.Value)
}

func (object JSONObjectValue) Print() string {
  returnString := "{"
  for k, v := range object.Map {
    returnString += fmt.Sprintf(`"%s":%s,`, k, v.Print())
  }
  returnString = returnString[:len(returnString) - 1]
  returnString += "}"
  return returnString
}

func (array JSONArrayValue) Print() string {
  returnString := "["
  for _, value := range array.Array {
    returnString += value.Print() + ","
  }
  returnString = returnString[:len(returnString) - 1]
  returnString += "]"
  return returnString
}
