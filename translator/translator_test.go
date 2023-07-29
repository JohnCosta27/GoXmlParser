package translator

import (
	"testing"
)

func TestTranslateSimple(t *testing.T) {
  json, err := TranslateJson("<a>hello</a>")
  if (err != nil) {
    t.Log("Expected error to be nil")
    t.FailNow()
  }

  if (json.Map["a"].Print() != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", json.Map["a"])
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
  

  if (jsonObjectA.Map["b"].Print() != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", json.Map["a"])
    t.FailNow()
  }

}

func TestDeepNesting(t *testing.T) {
  json, err := TranslateJson("<a><b><c><d>hello</d></d></b></a>")
  if (err != nil) {
    t.Log("Expected error to be nil")
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
  

  if (jsonObjectC.Map["d"].Print() != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", json.Map["a"])
    t.FailNow()
  }

}

func TestSiblingElements(t *testing.T) {
  json, err := TranslateJson("<a><b>hello</b><c>world</c></a>")
  if (err != nil) {
    t.Log("Expected error to be nil")
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

  if (jsonObjectA.Map["b"].Print() != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", json.Map["a"])
    t.FailNow()
  }

  if (jsonObjectA.Map["c"].Print() != "world") {
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
    t.Log("Expected error to be nil")
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
    t.Log("Expected `c` to be of type JSONObjectValue")
    t.FailNow()
  }

  if (len(jsonObjectB.Map) != 2) {
    t.Logf("Expected `c` to have 2 elements, it instead has %d\n", len(jsonObjectB.Map))
    t.FailNow()
  }
}
