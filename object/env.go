package object

type Environment struct {
	store map[string]string
}

func (e *Environment) Get(name string) (string, bool) {
	obj, ok := e.store[name]
	return obj, ok
}
func (e *Environment) Set(name string, val string) string {
	e.store[name] = val
	return val
}
func NewEnvironment() *Environment {
	s := make(map[string]string)
	return &Environment{store: s}
}
