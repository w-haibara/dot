package main

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
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

	var fn func(int, reflect.Value)
	fn = func(par int, v reflect.Value) {
		id++

		addChildlen(w, par, id, v.String())

		t := v.Type()
		if t.Kind() != reflect.Struct {
			return
		}

		id := id
		for _, f := range reflect.VisibleFields(t) {
			fn(id, v.FieldByIndex(f.Index))
		}
	}

	for _, val := range values {
		fn(par, reflect.ValueOf(val))
	}

	return nil
}

func addChildlen(w io.Writer, id1, id2 int, label string) {
	if id1 < 0 {
		fmt.Fprintln(w, "	"+strconv.Itoa(id2)+`[label="`+label+`"];`)
		return
	}

	fmt.Fprintln(w, "	"+strconv.Itoa(id2)+`[label="`+label+`"];`+"\n	"+strconv.Itoa(id1)+"->"+strconv.Itoa(id2)+";")
}
