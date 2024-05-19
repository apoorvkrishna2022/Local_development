package main

import (
	"fmt"
	"reflect"
)

type test1 struct {
	a bool
	b bool
}

type test2 struct {
	temp1 test1
}

type test3 struct {
	temp2 test2
}

type test struct {
	x test1
	y test2
	z test3
}

func main() {
	apoorv := new(test)
	*apoorv = test{
		x: test1{
			a: false,
			b: true,
		},
		y: test2{
			temp1: test1{
				a: true,
				b: false,
			},
		},
		z: test3{
			temp2: test2{
				temp1: test1{
					a: true,
					b: false,
				},
			},
		},
	}

	v := reflect.ValueOf(*apoorv)
	fmt.Println("value ", v.NumField())

	fmt.Println(reflect.ValueOf(v.Field(0).Interface()))
}
