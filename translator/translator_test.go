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
