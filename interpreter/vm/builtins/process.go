package builtins

type processClass struct {
	valueStub
}

func NewProcessClass() Class {
	f := &processClass{}
	f.initialize()
	f.class = NewClassValue().(Class)
	return f
}

type processValue struct {
	valueStub
}

func (class *processClass) New() Value {
	p := &processClass{}
	p.initialize()
	p.class = class

	return p
}