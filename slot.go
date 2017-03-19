package orumb

type Slot struct {
	c Class
	n string
	i int
}

func (s *Slot) Class() Class {
	return s.c
}

func (s *Slot) Name() string {
	return s.n
}

func (s *Slot) GetValue(o Object) (Object, error) {
	if so, ok := o.(*SlotObject); ok {
		return so.slots[s.i], nil
	} else {
		return nil, fmt.Errorf("Expected SlotObject, not %T", o)
	}
}

func (s *Slot) SetValue(o, v Object) error {
	if so, ok := o.(*SlotObject); ok {
		so.slots[s.i] = v
		return nil
	} else {
		return fmt.Errorf("Expected SlotObject, not %T", o)
	}
}

//
// SlotObject
//

type SlotObject struct {
	c Class
	slots []Object
}

func (o *SlotObject) Class() Class {
	return ObjectClass
}

func getSlotAttrCount(attrs []Attr) int {
	if attrs == nil {
		return 0
	} else {
		var c int = 0
		for _, a := range attrs {
			if _, ok := a.(Slot) {
				c += 1
			}
		}
		return c
	}
}

func MakeSlotObject(c Class) *SlotObject {
	return &SlotObject{c, make([]Object, getSlotAttrCount(c.Attrs())}
}
