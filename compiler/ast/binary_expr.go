package ast

import (
	"github.com/d5/tengo/compiler/source"
	"github.com/d5/tengo/compiler/token"
)

// BinaryExpr represents a binary operator expression.
type BinaryExpr struct {
	LHS      Expr        `json:"lhs"`
	RHS      Expr        `json:"rhs"`
	Token    token.Token `json:"token"`
	TokenPos source.Pos  `json:"token_pos"`
}

func (e *BinaryExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BinaryExpr) Pos() source.Pos {
	return e.LHS.Pos()
}

// End returns the position of first character immediately after the node.
func (e *BinaryExpr) End() source.Pos {
	return e.RHS.End()
}

func (e *BinaryExpr) String() string {
	return "(" + e.LHS.String() + " " + e.Token.String() + " " + e.RHS.String() + ")"
}
