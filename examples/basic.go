package main

import (
	"fmt"
	"github.com/cnchef/gconv"
)

func main() {
	var v any = "123"

	fmt.Println(gconv.Cast[int](v))       // 123
	fmt.Println(gconv.Cast[float64](v))   // 123.0
	fmt.Println(gconv.Cast[string](v))    // "123"
	fmt.Println(gconv.Cast[bool]("true")) // true

	m := map[string]any{"a": 1}
	fmt.Println(gconv.Cast[map[string]any](m)) // map[a:1]
}
