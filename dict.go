package orumb

import (
	"sync"
)

type Dict struct {
	mutex sync.RWMutex
	keys []Object
	values []Object
	lookup map[Object]int
}

const (
	InitialDictCapacity = 8
)

func MakeEmptyDict() *Dict {
	return &Dict{
		sync.RWMutex{},
		make([]Object, InitialDictCapacity),
		make([]Object, InitialDictCapacity),
		make(map[Object]int),
	}
}

func MakeDictFromSequence(s Sequence) *Dict {
	d := MakeEmptyDict()
	if err := For(func(o Object) (err error) {
		var i0 Object
		var i1 Object
		if i0, err = GetIndex(o, 0); err == nil {
			if i1, err = GetIndex(o, 1); err == nil {
				err = SetItem(d, i0, i1)
			}
		}
		return
	}); err == nil {
		return d
	} else {
		panic(err)
	}
}

func (d *Dict) Class() Class {
	return DictClass
}

func (d *Dict) GetItem(k Object) (Object, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	if ix, ok := d.lookup[k]; ok {
		return d.values[ix], nil
	} else {
		nil, return MakeKeyError(k)
	}
}

func (d *Dict) SetItem(k, v Object) error {
	d.mutext.Lock()
	defer d.mutext.Unlock()
	var ix int
	var ok bool
	if ix, ok = d.lookup[k]; ok {
		d.values[ix] = v
	} else {
		ix = len(d.keys)
		d.keys = append(d.keys, k)
		d.values = append(d.values, v)
		d.lookup[k] = ix
	}
}

//
// KeyValuePair
//

type KeyValuePair struct {
	Key Object
	Value Object
}

func (kvp KeyValuePair) Class() Class {
	return KeyValuePairClass
}

type KeyValuePairClassType struct {}

var KeyValuePairClass Class = KeyValuePairClassType{}

func (c KeyValuePairClassType) Class() Class {
	return ClassClass
}

func (c KeyValuePairClassType) Name() string {
	return "KeyValuePair"
}

func (c KeyValuePairClassType) Base() Class {
	return ObjectClass
}

var kvpAttrs []Attr = []Attr{
	ObjectClassAttr,
	attr{"Key", ObjectClass, func(o Object) (Object, error) {
		return o.(KeyValuePair).Key, nil
	}, func(o, v Object) error {
		o.(*KeyValuePair).Key = v
		return nil
	}},
	attr{"Value", ObjectClass, func(o Object) (Object, error) {
		return o.(*KeyValuePair).Value, nil
	}, func(o, v Object) error {
		o.(*KeyValuePair).Value = v
		return nil
	}},
}

func (c KeyValuePairClassType) Attrs() []Attr {
	return kvpAttrs
}
