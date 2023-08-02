package translator

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestTranslateSimple(t *testing.T) {
  json, err := TranslateJson("<a>hello</a>")
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  stringA, ok := json.Map["a"].(JSONStringValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONStringValue")
    t.FailNow()
  }

  if (stringA.Value != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", stringA.Value)
    t.FailNow()
  }

}

func TestNesting(t *testing.T) {
  json, err := TranslateJson("<a><b>hello</b></a>")
  if (err != nil) {
    t.Log("Expected error to be nil")
    t.FailNow()
  }

  jsonObjectA, ok := json.Map["a"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }

  stringB, ok := jsonObjectA.Map["b"].(JSONStringValue)
  if (!ok) {
    t.Log("Expected `v` to be of type JSONStringValue")
    t.FailNow()
  }
  

  if (stringB.Value != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", stringB.Value)
    t.FailNow()
  }

}

func TestDeepNesting(t *testing.T) {
  json, err := TranslateJson("<a><b><c><d>hello</d></c></b></a>")
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  jsonObjectA, ok := json.Map["a"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }

  jsonObjectB, ok := jsonObjectA.Map["b"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }

  jsonObjectC, ok := jsonObjectB.Map["c"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }
  
  stringD, ok := jsonObjectC.Map["d"].(JSONStringValue)
  if (!ok) {
    t.Log("Expected `d` to be of type JSONStringValue")
    t.FailNow()
  }

  if (stringD.Value != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", stringD.Value)
    t.FailNow()
  }

}

func TestSiblingElements(t *testing.T) {
  json, err := TranslateJson("<a><b>hello</b><c>world</c></a>")
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  jsonObjectA, ok := json.Map["a"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }

  if (len(jsonObjectA.Map) != 2) {
    t.Log("Expected `a` to have 2 elements")
    t.FailNow()
  }

  stringB, ok := jsonObjectA.Map["b"].(JSONStringValue)
  if (!ok) {
    t.Log("Expected `b` to be of type JSONStringValue")
    t.FailNow()
  }

  stringC, ok := jsonObjectA.Map["c"].(JSONStringValue)
  if (!ok) {
    t.Log("Expected `c` to be of type JSONStringValue")
    t.FailNow()
  }

  if (stringB.Value != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", json.Map["a"])
    t.FailNow()
  }

  if (stringC.Value != "world") {
    t.Logf(`Expected "world" but found %s` + "\n", json.Map["a"])
    t.FailNow()
  }

}

func TestNestingAndSibling(t *testing.T) {
  json, err := TranslateJson(`<a>
    <b>
      <d>hello</d>
      <e>world</e>
    </b>
    <c>
      World
    </c>
  </a>`)
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  jsonObjectA, ok := json.Map["a"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }

  if (len(jsonObjectA.Map) != 2) {
    t.Log("Expected `a` to have 2 elements")
    t.FailNow()
  }

  jsonObjectB, ok := jsonObjectA.Map["b"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `b` to be of type JSONObjectValue")
    t.FailNow()
  }

  if (len(jsonObjectB.Map) != 2) {
    t.Logf("Expected `b` to have 2 elements, it instead has %d\n", len(jsonObjectB.Map))
    t.FailNow()
  }

  snaps.MatchSnapshot(t, json.Print())
}

func TestSameKeySiblings(t *testing.T) {
  json, err := TranslateJson(`<a>
    <b>Hello</b>
    <b>World</b>
  </a>`)
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, json.Print())

  jsonObjectA, ok := json.Map["a"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }

  arrayB, ok := jsonObjectA.Map["b"].(JSONArrayValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONArrayValue")
    t.FailNow()
  }

  if (len(arrayB.Array) != 2) {
    t.Logf("Expected `a` to have 2 elements, it instead has %d\n", len(arrayB.Array))
    t.FailNow()
  }

  zeroIndex, ok := arrayB.Array[0].(JSONStringValue)
  if (!ok) {
    t.Log("Expected 0th index to be of type JSONStringValue")
    t.FailNow()
  }

  firstIndex, ok := arrayB.Array[1].(JSONStringValue)
  if (!ok) {
    t.Log("Expected 1st index to be of type JSONStringValue")
    t.FailNow()
  }

  if (zeroIndex.Value != "Hello") {
    t.Logf("Expected `Hello`, but got %s", zeroIndex.Value)
    t.FailNow()
  }

  if (firstIndex.Value != "World") {
    t.Logf("Expected `World`, but got %s", zeroIndex.Value)
    t.FailNow()
  }
}

func TestDeepNestingArray(t *testing.T) {
  json, err := TranslateJson(`
    <a>
      <c>Hello</c>
      <b>
        <d>123</d>
        <d>321</d>
      </b>
    </a>
    `)

  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, json.Print())
}

func TestMoreComplex(t *testing.T) {
  json, err := TranslateJson(`<a>
    <b>Hello</b>
    <b>World</b>
    <c>Another different value</c>
    <b>dmsakmdsa</b>
    <b>another</b>
    <d>123</d>
    <d>123</d>
    <d>123</d>
    <d>123</d>
    <e>
      <a>123</a>
      <a>321</a>
    </e>
  </a>`)
  if (err != nil) {
    t.Log(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, json.Print())
}
