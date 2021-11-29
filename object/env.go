//Object 是用于给所有语句的执行结构构造一个鸭子类型的返回值
package object

//Environment 是给每个用户一个独立的执行环境
type Environment struct {
	store map[string]string
}

//Get 用于从执行环境中获取值
func (e *Environment) Get(name string) (string, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

//Set 用于初始化执行环境
func (e *Environment) Set(name string, val string) string {
	e.store[name] = val
	return val
}

//New 用于新建一个执行环境，
func NewEnvironment() *Environment {
	s := make(map[string]string)
	return &Environment{store: s}
}
