package types

const KeyUnknown = "unknown"

type Unknown struct {
	X string `json:"x"`
}

var _ Type = (*Unknown)(nil)

func (x *Unknown) Key() string {
	return KeyUnknown
}

func (x *Unknown) String() string {
	return x.Key() + "(" + x.X + ")"
}

func (x *Unknown) Sortable() bool {
	return false
}

func (x *Unknown) From(v interface{}) interface{} {
	return v
}

func (x *Unknown) Default(string) interface{} {
	return x.X
}

func NewUnknown(x string) *Wrapped {
	return Wrap(&Unknown{X: x})
}
