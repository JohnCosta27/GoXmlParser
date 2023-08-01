package main

import (
	"fmt"
	"syscall/js"

	"johncosta.tech/xmlparse/translator"
)

func XmlToJson() js.Func {
  return js.FuncOf(func(this js.Value, p []js.Value) interface{} {
    obj, err := translator.TranslateJson(p[0].String())
    if (err != nil) {
      return err.Error()
    }

    return obj.Print()
  })
}

func main() {
  c := make(chan bool)
  fmt.Println("Hello Web Assembly from Go!")

  js.Global().Set("XmlToJson", XmlToJson())
  <-c
}
