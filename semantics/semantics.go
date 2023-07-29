package semantics

import (
	"errors"

	"johncosta.tech/xmlparse/AST"
	"johncosta.tech/xmlparse/utils"
)

func SemanticAnalysis(ast *AST.Element) error {
  
  stack := utils.StackFactory[string]()

  var err error = nil

  ast.Walk(func (node AST.ASTNode) {
    switch node.(type) {
    case *AST.OpenTag:
      stack.Push(node.(*AST.OpenTag).NAME.TokenContent)
    case *AST.CloseTag:
      popped := stack.Pop()
      if (node.(*AST.CloseTag).NAME.TokenContent != popped) {
        err = errors.New("Semantics | Expected: " + popped + " and got: " + node.(*AST.CloseTag).NAME.TokenContent)
      }
    }
  }, func(node AST.ASTNode) {})

  return err
}
