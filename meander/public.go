package meander

// Facade ...
type Facade interface {
	Public() interface{}
}

// Public ...
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
