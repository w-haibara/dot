package main

import (
	"os"

	"github.com/w-haibara/v2d"
)

type T1 struct {
	V1 int
	V2 bool
}

type T2 struct {
	V1 T1 `label:"v1"`
	V2 T1 `label:"v2"`
	V3 string
}

func main() {
	v := T2{
		V1: T1{
			V1: 0,
			V2: false,
		},
		V2: T1{
			V1: 1,
			V2: true,
		},
		V3: "foo",
	}

	d := v2d.NewDot()

	if err := d.Render(os.Stdout, v); err != nil {
		panic(err.Error())
	}
}
