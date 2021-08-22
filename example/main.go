package main

import (
	"os"

	"github.com/w-haibara/v2d"
)

type t1 struct {
	v1 int
	v2 bool
}

type t2 struct {
	v1 t1
	v2 t1
	v3 string
}

func main() {
	v := t2{
		v1: t1{
			v1: 0,
			v2: false,
		},
		v2: t1{
			v1: 1,
			v2: true,
		},
		v3: "foo",
	}

	d := v2d.NewDot()

	if err := d.Render(os.Stdout, v); err != nil {
		panic(err.Error())
	}
}
