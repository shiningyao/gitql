package expression

import (
	"fmt"

	"github.com/gitql/gitql/sql"
)

type Comparison struct {
	BinaryExpression
	ChildType sql.Type
}

func (*Comparison) Type() sql.Type {
	return sql.Boolean
}

func (*Comparison) Name() string {
	return ""
}

type Equals struct {
	Comparison
}

func NewEquals(left sql.Expression, right sql.Expression) *Equals {
	// FIXME: enable this again
	// checkEqualTypes(left, right)
	return &Equals{Comparison{BinaryExpression{left, right}, left.Type()}}
}

func (e Equals) Eval(row sql.Row) interface{} {
	a := e.Left.Eval(row)
	b := e.Right.Eval(row)
	return e.ChildType.Compare(a, b) == 0
}

func (c *Equals) TransformUp(f func(sql.Expression) sql.Expression) sql.Expression {
	lc := c.BinaryExpression.Left.TransformUp(f)
	rc := c.BinaryExpression.Right.TransformUp(f)

	return f(NewEquals(lc, rc))
}

type GreaterThan struct {
	Comparison
}

func NewGreaterThan(left sql.Expression, right sql.Expression) *GreaterThan {
	// FIXME: enable this again
	// checkEqualTypes(left, right)
	return &GreaterThan{Comparison{BinaryExpression{left, right}, left.Type()}}
}

func (e GreaterThan) Eval(row sql.Row) interface{} {
	a := e.Left.Eval(row)
	b := e.Right.Eval(row)
	return e.ChildType.Compare(a, b) == 1
}

func (c *GreaterThan) TransformUp(f func(sql.Expression) sql.Expression) sql.Expression {
	lc := c.BinaryExpression.Left.TransformUp(f)
	rc := c.BinaryExpression.Right.TransformUp(f)

	return f(NewGreaterThan(lc, rc))
}

type LessThan struct {
	Comparison
}

func NewLessThan(left sql.Expression, right sql.Expression) *LessThan {
	// FIXME: enable this again
	// checkEqualTypes(left, right)
	return &LessThan{Comparison{BinaryExpression{left, right}, left.Type()}}
}

func (e LessThan) Eval(row sql.Row) interface{} {
	a := e.Left.Eval(row)
	b := e.Right.Eval(row)
	return e.ChildType.Compare(a, b) == -1
}

func (c *LessThan) TransformUp(f func(sql.Expression) sql.Expression) sql.Expression {
	lc := c.BinaryExpression.Left.TransformUp(f)
	rc := c.BinaryExpression.Right.TransformUp(f)

	return f(NewLessThan(lc, rc))
}

type GreaterThanOrEqual struct {
	Comparison
}

func NewGreaterThanOrEqual(left sql.Expression, right sql.Expression) *GreaterThanOrEqual {
	// FIXME: enable this again
	// checkEqualTypes(left, right)
	return &GreaterThanOrEqual{Comparison{BinaryExpression{left, right}, left.Type()}}
}

func (e GreaterThanOrEqual) Eval(row sql.Row) interface{} {
	a := e.Left.Eval(row)
	b := e.Right.Eval(row)
	return e.ChildType.Compare(a, b) > -1
}

func (c *GreaterThanOrEqual) TransformUp(f func(sql.Expression) sql.Expression) sql.Expression {
	lc := c.BinaryExpression.Left.TransformUp(f)
	rc := c.BinaryExpression.Right.TransformUp(f)

	return f(NewGreaterThanOrEqual(lc, rc))
}

type LessThanOrEqual struct {
	Comparison
}

func NewLessThanOrEqual(left sql.Expression, right sql.Expression) *LessThanOrEqual {
	// FIXME: enable this again
	// checkEqualTypes(left, right)
	return &LessThanOrEqual{Comparison{BinaryExpression{left, right}, left.Type()}}
}

func (e LessThanOrEqual) Eval(row sql.Row) interface{} {
	a := e.Left.Eval(row)
	b := e.Right.Eval(row)
	return e.ChildType.Compare(a, b) < 1
}

func (c *LessThanOrEqual) TransformUp(f func(sql.Expression) sql.Expression) sql.Expression {
	lc := c.BinaryExpression.Left.TransformUp(f)
	rc := c.BinaryExpression.Right.TransformUp(f)

	return f(NewLessThanOrEqual(lc, rc))
}

func checkEqualTypes(a sql.Expression, b sql.Expression) {
	if a.Resolved() && b.Resolved() && a.Type() != b.Type() {
		panic(fmt.Errorf("both types should be equal: %v and %v\n", a, b))
	}
}

func (e Equals) Name() string {
	return e.Left.Name() + "==" + e.Right.Name()
}
