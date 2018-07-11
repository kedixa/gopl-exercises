package eval

// Variables get all variables in e
func Variables(e Expr, m map[Var]float64) {
	switch e := e.(type) {
	case literal:
		return
	case Var:
		m[Var(e)] = 0.0
	case unary:
		Variables(e.x, m)
	case binary:
		Variables(e.x, m)
		Variables(e.y, m)
	case call:
		for _, arg := range e.args {
			Variables(arg, m)
		}
	}
}
