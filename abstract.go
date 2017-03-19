package orumb

// Simpe Attr implementation for quickly defining attributes
type attr struct {
    cls Class
    name string
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
        *err = e
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
