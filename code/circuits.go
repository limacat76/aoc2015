package code

type registry interface {
	fetch(name string) gate
}

type circuit struct {
	dir map[string]gate
}

func (r circuit) fetch(name string) gate {
	return r.dir[name]
}

// gates
type gate interface {
	output(r registry) uint16
}

// Unary Gates
type input struct {
	value uint16
}

func (g input) output(r registry) uint16 {
	return g.value
}

type not struct {
	x string
}

func (g not) output(r registry) uint16 {
	xg := (r.fetch(g.x))
	return ^xg.output(r)
}

// Binary Gates
type and struct {
	x, y string
}

func (g and) output(r registry) uint16 {
	xg := (r.fetch(g.x))
	yg := (r.fetch(g.y))
	return xg.output(r) & yg.output(r)
}

type or struct {
	x, y string
}

func (g or) output(r registry) uint16 {
	xg := (r.fetch(g.x))
	yg := (r.fetch(g.y))
	return xg.output(r) | yg.output(r)
}

// Shift Operations
type lshift struct {
	x    string
	bits uint
}

func (g lshift) output(r registry) uint16 {
	xg := (r.fetch(g.x))
	return xg.output(r) << g.bits
}

type rshift struct {
	x    string
	bits uint
}

func (g rshift) output(r registry) uint16 {
	xg := (r.fetch(g.x))
	return xg.output(r) >> g.bits
}
