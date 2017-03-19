package orumb

type Object interface {
	Class() Class
}

type Class interface {
	Object
	Name() string
	Base() Class
	Attrs() []Attr
}

type Attr interface {
	Object
	Name() string
	GetValue(o Object) (Object, error)
	SetValue(o, v Object) error
}

type Sequence interface {
	GetIndex(i int) (Object, error)
	SetIndex(i int, v Object) error
}

type Mapping interface {
	GetItem(k Object) (Object, error)
	SetItem(k, v Object) error
}
