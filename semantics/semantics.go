package semantics

import (
	"johncosta.tech/xmlparse/AST"
	"johncosta.tech/xmlparse/utils"
)

func SemanticAnalysis(ast *AST.Tag) error {
  
  stack := utils.StackFactory[string]()

  ast.Walk(func (node AST.ASTNode) {
    switch node.(type) {
    case *AST.OpenTag:
      stack.Push(node.(*AST.OpenTag).TagName.Text)
    case *AST.CloseTag:
      popped := stack.Pop()
      if (node.(*AST.CloseTag).TagName.Text != popped) {
        panic("I expected: " + popped + " and got: " + node.(*AST.CloseTag).TagName.Text)
      }
    }
  })

  return nil
}
