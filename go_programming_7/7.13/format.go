package main

import "fmt"

func (v Var) String(env Env) string {
	return string(v)
}

func (l literal) String(env Env) string {
	return fmt.Sprintf("%f", l)
}

func (l logic) String(env Env) string {
	return fmt.Sprintf("%f%s%f", l.x.Eval(env), l.op, l.y.Eval(env))
}

func (b binary) String(env Env) string {
	return fmt.Sprintf("%f%s%f", b.x.Eval(env), string(b.op), b.y.Eval(env))
}

func (u unary) String(env Env) string {
	return fmt.Sprintf("%s%f", string(u.op), u.x.Eval(env))
}

func (c call) String(env Env) string {
	return fmt.Sprintf("%s%f", c.fn, c.args[0].Eval(env))
}
