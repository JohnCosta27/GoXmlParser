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
  returnString := "{\n"
  for k, v := range object.Map {
    returnString += fmt.Sprintf(`"%s":%s,%s`, k, v.Print(), "\n")
  }
  returnString += "\n}"
  return returnString
}

func (array JSONArrayValue) Print() string {
  returnString := "[\n"
  for _, value := range array.Array {
    returnString += value.Print() + ","
  }
  return returnString
}
