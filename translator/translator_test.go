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
  

  if (jsonObjectA.Value.Map["b"].Print() != "hello") {
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

  jsonObjectB, ok := jsonObjectA.Value.Map["b"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }

  jsonObjectC, ok := jsonObjectB.Value.Map["c"].(JSONObjectValue)
  if (!ok) {
    t.Log("Expected `a` to be of type JSONObjectValue")
    t.FailNow()
  }
  

  if (jsonObjectC.Value.Map["d"].Print() != "hello") {
    t.Logf(`Expected "hello" but found %s` + "\n", json.Map["a"])
    t.FailNow()
  }

}

/*
func TestTranslateNesting(t *testing.T) {
  json, err := TranslateJson("<a><b>Hello</b></a>")
  if (err != nil) {
    t.Log("Expected error to be nil")
    t.FailNow()
  }

  if (json != `{"a":{"b":"Hello"}}`) {
    t.Logf("Incorrect json output, got %s", json)
    t.FailNow()
  }
}


func TestTranslateArray(t *testing.T) {
  json, err := TranslateJson("<a><b>Hello</b><b>bruh</b></a>")
  if (err != nil) {
    t.Log("Expected error to be nil")
    t.FailNow()
  }

  if (json != `{"a":{"b":"Hello"}}`) {
    t.Logf("Incorrect json output, got %s", json)
    t.FailNow()
  }
}
*/
