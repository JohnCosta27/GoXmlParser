package parser

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"johncosta.tech/xmlparse/AST"
	"johncosta.tech/xmlparse/lexer"
)

// ------------- INDIVIDUAL PARSES ------------- //

func TestSimple(t *testing.T) {
  ast, err := Parse(lexer.Tokenize("<a>hello</a>"))

  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print(0))
}

func TestNesting(t *testing.T) {
  ast, err := Parse(lexer.Tokenize("<a><b>Something</b></a>"))

  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print(0))
}

func TestAttributes(t *testing.T) {
  ast, err := Parse(lexer.Tokenize(`<a hello="world"></a>`))

  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print(0))
}

func TestSelfClosing(t *testing.T) {
  ast, err := Parse(lexer.Tokenize(`<atag bruh="hello-world" />`))

  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print(0))
}

func TestSibling(t *testing.T) {
  ast, err := Parse(lexer.Tokenize("<a><b>Hello</b><c>World</c></a>"))

  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  if (ast.ElementSuffix.Content.Type != AST.CONTENT_ELEMENT) {
    t.Log("Expected content to be element")
    t.FailNow()
  }

  el := ast.ElementSuffix.Content.Element
  if (el.OpenTag.NAME.TokenContent != "b") {
    t.Log("Expected first sibling to be `b`")
    t.FailNow()
  }

  otherEl := ast.ElementSuffix.Content.Content
  if (otherEl.Type != AST.CONTENT_ELEMENT) {
    t.Log("Expected content to be element")
    t.FailNow()
  }

  if (otherEl.Element.OpenTag.NAME.TokenContent != "c") {
    t.Log("Expected first sibling to be `c`")
    t.FailNow()
  }

}

func TestMedium(t *testing.T) {
  ast, err := Parse(lexer.Tokenize("<a><b>hello</b><c><d>world</d><e>world2</e><f>world3</f></c></a>"))

  if err != nil {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print(0))
}

func TestBigObject(t *testing.T) {
  ast, err := Parse(lexer.Tokenize(`<Tests xmlns="http://www.adatum.com">
  <Test TestId="0001" TestType="CMD">
    <Name>Convert number to string</Name>
    <CommandLine>Examp1.EXE</CommandLine>
    <Input>1</Input>
    <Output>One</Output>
  </Test>
  <Test TestId="0002" TestType="CMD">
    <Name>Find succeeding characters</Name>
    <CommandLine>Examp2.EXE</CommandLine>
    <Input>abc</Input>
    <Output>def</Output>
  </Test>
  <Test TestId="0003" TestType="GUI">
    <Name>Convert multiple numbers to strings</Name>
    <CommandLine>Examp2.EXE /Verbose</CommandLine>
    <Input>123</Input>
    <Output>One Two Three</Output>
  </Test>
  <Test TestId="0004" TestType="GUI">
    <Name>Find correlated key</Name>
    <CommandLine>Examp3.EXE</CommandLine>
    <Input>a1</Input>
    <Output>b1</Output>
  </Test>
  <Test TestId="0005" TestType="GUI">
    <Name>Count characters</Name>
    <CommandLine>FinalExamp.EXE</CommandLine>
    <Input>This is a test</Input>
    <Output>14</Output>
  </Test>
  <Test TestId="0006" TestType="GUI">
    <Name>Another Test</Name>
    <CommandLine>Examp2.EXE</CommandLine>
    <Input>Test Input</Input>
    <Output>10</Output>
  </Test>
</Tests>`))

  if (err != nil) {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print(0))
}

func TestDoubleTesting(t *testing.T) {
  ast, err := Parse(lexer.Tokenize(`
    <a>
      <c>Hello</c>
      <b>
        <d>123</d>
        <d>321</d>
      </b>
    </a>
  `))

  if err != nil {
    t.Error(err)
    t.FailNow()
  }

  snaps.MatchSnapshot(t, ast.Print(0))
}
