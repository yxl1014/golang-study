package object

// 环境

// 新建一个子作用域
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object // 本作用域的map
	outer *Environment      // 父作用域
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	// 如果自己这没有，那就看看父作用域
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
