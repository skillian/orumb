package orumb

import (
	"errors"
	"fmt"
 	"sync"
)

var (
	clsMapName map[string]Class = make(map[string]Class)
	clsMapLock sync.RWMutex = sync.RWMutex{}
)

func GetClassByName(name string) (c Class, err error) {
	clsMapLock.RLock()
	defer clsMapLock.RUnlock()
	var ok bool
	if c, ok = clsMapName[name]; ok {
		err = nil
	} else {
		err = fmt.Errorf("Class not found: %s", name)
	}
}

func RegisterClass(c Class) error {
	clsMapLock.Lock()
	defer clsMapLock.Unlock()
	if _, ok := clsMapName[c.Name()]; ok {
		return fmt.Errorf("Redefinition of class %s", c.Name())
	} else {
		clsMapName[c.Name()] = c
		return nil
	}
}

type ClassClassType struct {}

var ClassClass Class = ClassClassType{}

func (c ClassClassType) Class() Class {
	return ClassClass
}

func (c ClassClassType) Name() string {
	return "Class"
}

func (c ClassClassType) Base() Class {
	return ObjectClass
}

var clsAttrs []Attr = []Attr{ObjectClassAttr, ObjectNameAttr, ClassBaseAttr, ClassAttrsAttr}

func (c ClassClassType) Attrs() []Attr {
	return clsAttrs
}

//
// ClassBaseAttr
//

type ClassBaseAttrType struct {}

var ClassBaseAttr Attr = ClassBaseAttrType{}

func (a ClassBaseAttrType) Class() Class {
	return ClassClass
}

func (a ClassBaseAttrType) Name() string {
	return "Base"
}

func (a ClassBaseAttrType) GetValue(o Object) (Object, error) {
	if cls, ok := o.(Class); ok {
		return cls.Base(), nil
	} else {
		return fmt.Errorf("Class.Base requires a class as its context")
	}
}

var cannotSetClassBase error = errors.New("Cannot set Class.Base")

func (a ClassBaseAttrType) SetValue(o, v Object) error {
	return cannotSetClassBase
}

//
// ClassAttrsAttr
//

type ClassAttrsAttrType struct {}

var ClassAttrsAttr Attr = ClassAttrsAttrType{}

func (a ClassAttrsAttrType) Class() Class {
	return TupleClass
}

func (a ClassAttrsAttrType) Name() string {
	return "Attrs"
}

func (a ClassAttrsAttrType) GetValue(o Object) (Object, error) {
	if cls, ok := o.(Class); ok {
		return MakeTuple(cls.Attrs()...), nil
	} else {
		return fmt.Errorf("Class.Attrs requires Class as context, not %T", o)
	}
}

var cannotSetClassAttrs error = errors.New("Cannot set Class.Attrs")

func (a ClassAttrsAttrType) SetValue(o, v Object) error {
	return cannotSetClassAttrs
}
