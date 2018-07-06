package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	values := appendValues(nil, t)
	var buf bytes.Buffer
	buf.WriteString("{")
	if len(values) > 0 {
		buf.WriteString(fmt.Sprint(values[0]))
	}
	for _, v := range values[1:] {
		buf.WriteString(fmt.Sprintf(", %d", v))
	}
	buf.WriteString("}")
	return buf.String()
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	var root *tree
	for _, v := range []int{1, 3, 7, 2, 5, 4} {
		root = add(root, v)
	}
	fmt.Println(root.String())
}
