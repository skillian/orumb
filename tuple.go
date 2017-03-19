package orumb

type Tuple struct {
	items []Object
}

func MakeTuple(items... Object) Tuple {
	return Tuple{items}
}

func (t Tuple) Class() Class {
	return TupleClass
}

type TupleClassType struct {}

var TupleClass Class = TupleClassType{}

func (c TupleClassType) Class() Class {
	return ClassClass
}

func (c TupleClassType) Name() string {
	return "Tuple"
}

var tupleAttrs []Attr = []Attr{
	ObjectClassAttr,
	ObjectLenAttr,
}

func (c TupleClassType) Attrs() []Attr {
	return tupleAttrs
}

//
// 
//
