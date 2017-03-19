package orumb

import (
    "errors"
    "fmt"
)

// Simpe Attr implementation for quickly defining attributes
type attr struct {
    name string
    cls Class
    get func(o Object) (Object, error)
    set func(o, v Object) error
}

func (a attr) Class() Class {
    return a.cls
}

func (a attr) Name() string {
    return a.name
}

func RecoverToError(err *error) {
    e := recover()
    if *err == nil && e != nil {
        *err = errors.New(fmt.Sprint(e))
    }
}

func (a attr) GetValue(o Object) (v Object, err error) {
    defer RecoverToError(&err)
    return a.get(o)
}

func (a attr) SetValue(o, v Object) (err error) {
  defer RecoverToError(&err)
  return a.set(o, v)
}
