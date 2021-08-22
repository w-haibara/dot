package v2d

import (
	"fmt"
	"io"
	"reflect"
)

type Dot struct {
}

func NewDot() *Dot {
	return new(Dot)
}

func (d *Dot) Render(w io.Writer, values ...interface{}) error {
	fmt.Fprintln(w, "digraph  {")
	defer fmt.Fprintln(w, "}")

	par := -1
	id := par

	var fn func(int, reflect.Value, string)
	fn = func(par int, v reflect.Value, label string) {
		id++

		t := v.Type()
		if t.Kind() != reflect.Struct {
			addChild(w, par, id, v)
			return
		}

		if label != "" {
			label = ":" + label
		}
		addChild(w, par, id, t.Name()+label)

		id := id
		for _, f := range reflect.VisibleFields(t) {
			fn(id, v.FieldByIndex(f.Index), f.Tag.Get("label"))
		}
	}

	for _, val := range values {
		fn(par, reflect.ValueOf(val), "")
	}

	return nil
}

func addChild(w io.Writer, id1, id2 int, label interface{}) {
	if id1 < 0 {
		fmt.Fprintf(w, `	%d[label="%v"];`+"\n", id2, label)
		return
	}

	fmt.Fprintf(w, `	%d[label="%v"];`+"\n"+`	%d->%d`+"\n", id2, label, id1, id2)
}
