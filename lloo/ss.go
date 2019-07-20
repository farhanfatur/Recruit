package main

import (
	"fmt"
	"reflect"
	"sort"
)

func ssa(a []string, b []string) []string {
	var m map[string]bool
	m = make(map[string]bool, len(b))

	for _, s := range b {
		m[s] = false
	}

	var dif []string
	for _, s := range a {
		if _, ok := m[s]; !ok {
			dif = append(dif, s)
			continue
		}
		m[s] = true
	}
	sort.Strings(dif)
	return dif
}

func main() {
	var a = []string{"1", "2"}
	var b = []string{"1", "2", "3", "4"}
	if !reflect.DeepEqual(a, b) {
		fmt.Println(ssa(b, a))
	}

}
