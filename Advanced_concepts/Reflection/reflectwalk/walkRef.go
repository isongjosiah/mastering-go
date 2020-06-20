package main

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/reflectwalk"
)

// Values ...
type Values struct {
	Extra map[string]string
}

// WalkMap ...
type WalkMap struct {
	MapVal reflect.Value
	Keys   map[string]bool
	Values map[string]bool
}

// Map ...
func (t *WalkMap) Map(m reflect.Value) error {
	t.MapVal = m
	return nil
}

// MapElem ...
func (t *WalkMap) MapElem(m, k, v reflect.Value) error {
	if t.Keys == nil {
		t.Keys = make(map[string]bool)
		t.Values = make(map[string]bool)
	}

	t.Keys[k.Interface().(string)] = true
	t.Values[v.Interface().(string)] = true
	return nil
}

func main() {
	w := new(WalkMap)

	type S struct {
		Map map[string]string
	}

	data := &S{
		Map: map[string]string{
			"V1": "v1v",
			"V2": "v2v",
			"V3": "v3v",
			"V4": "v4v",
		},
	}

	err := reflectwalk.Walk(data, w)
	if err != nil {
		fmt.Println(err)
		return
	}
}
