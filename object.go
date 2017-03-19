package orumb

import (
    "errors"
)

var AttrReadOnly error = errors.New("Attr is read only")

type ObjectClassType struct {}

var ObjectClass Class = ObjectClassType{}

func (c ObjectClassType) Class() Class {
    return ClassClass
}

func (c ObjectClassType) Name() string {
    return "Object"
}

func (c ObjectClassType) Base() Class {
    return nil
}

var objAttrs []Attr = []Attr{ObjectClassAttr}

func (c ObjectClassType) Attrs() []Attr {
    return 
}

//
// ObjectClassAttr
//

type ObjectClassAttrType struct {}

var ObjectClassAttr Attr = ObjectClassAttrType{}

func (a ObjectClassAttr) Class() Class {
    return ClassClass
}

func (a ObjectClassAttrType) Name() string {
    return "Class"
}

func (a ObjectClassAttrType) GetValue(o Object) (Object, error) {
    return o.Class(), nil
}

func (a ObjectClassAttrType) SetValue(o, v Object) error {
    return AttrReadOnly
}

//
// ObjectNameAttr
//

type Namer interface {
    Name() string
}

type ObjectNameAttrType struct {}

var ObjectNameAttr Attr = ObjectNameAttrType{}

func (a ObjectNameAttrType) Class() Class {
    return StringClass
}

func (a ObjectNameAttrType) Name() string {
    return "Name"
}

func (a ObjectNameAttrType) GetValue(o Object) (Object, error) {
    if n, ok := o.(Namer); ok {
        return MakeString(n.Name()), nil
    } else {
        return nil, MakeTypeErrorString("Namer", o)
    }
}

func (a ObjectNameAttrType) SetValue(o, v Object) error {
    return AttrReadOnly
}
