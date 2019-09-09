package ast

import "first/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// このノードは、構文解析器が生成する全てのASTのルートノードになる
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name *Identifier // 束縛の識別子を保持するために必要
	Value string // 値を生成する式を保持するために必要
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) exporessionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}