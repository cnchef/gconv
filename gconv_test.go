package gconv

import (
	"encoding/json"
	"testing"
)

// TestToString 测试字符串转换
func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{"string", "hello", "hello"},
		{"int", 123, "123"},
		{"float64", 123.45, "123.45"},
		{"bool_true", true, "true"},
		{"bool_false", false, "false"},
		{"json.Number", json.Number("999"), "999"},
		{"map", map[string]any{"key": "value"}, `{"key":"value"}`},
		{"slice", []any{1, 2, 3}, "[1,2,3]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToString(tt.input)
			if result != tt.expected {
				t.Errorf("ToString(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestToInt 测试整数转换
func TestToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected int
	}{
		{"int", 123, 123},
		{"int64", int64(456), 456},
		{"float64", 789.99, 789},
		{"string_valid", "123", 123},
		{"string_invalid", "abc", 0},
		{"json.Number", json.Number("999"), 999},
		{"bool", true, 0},
		{"nil", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToInt(tt.input)
			if result != tt.expected {
				t.Errorf("ToInt(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestToFloat 测试浮点数转换
func TestToFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"float64", 123.45, 123.45},
		{"float32", float32(99.9), 99.9000015258789},
		{"int", 100, 100.0},
		{"int64", int64(200), 200.0},
		{"string_valid", "123.45", 123.45},
		{"string_invalid", "abc", 0.0},
		{"json.Number", json.Number("99.99"), 99.99},
		{"bool", false, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToFloat(tt.input)
			if result != tt.expected {
				t.Errorf("ToFloat(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestToBool 测试布尔值转换
func TestToBool(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"bool_true", true, true},
		{"bool_false", false, false},
		{"string_true", "true", true},
		{"string_false", "false", false},
		{"string_1", "1", true},
		{"string_0", "0", false},
		{"int_nonzero", 123, true},
		{"int_zero", 0, false},
		{"float_nonzero", 1.5, true},
		{"float_zero", 0.0, false},
		{"json.Number", json.Number("1"), true},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToBool(tt.input)
			if result != tt.expected {
				t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestToMap 测试映射转换
func TestToMap(t *testing.T) {
	validMap := map[string]any{"key": "value", "num": 123}

	tests := []struct {
		name     string
		input    any
		expected map[string]any
	}{
		{"valid_map", validMap, validMap},
		{"string", "not a map", map[string]any{}},
		{"int", 123, map[string]any{}},
		{"nil", nil, map[string]any{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMap(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("ToMap(%v) length = %v, want %v", tt.input, len(result), len(tt.expected))
			}
		})
	}
}

// TestToSlice 测试切片转换
func TestToSlice(t *testing.T) {
	validSlice := []any{1, "two", 3.0}

	tests := []struct {
		name     string
		input    any
		expected []any
	}{
		{"valid_slice", validSlice, validSlice},
		{"string", "not a slice", []any{}},
		{"int", 123, []any{}},
		{"nil", nil, []any{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSlice(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("ToSlice(%v) length = %v, want %v", tt.input, len(result), len(tt.expected))
			}
		})
	}
}

// TestCastString 测试泛型转换到字符串
func TestCastString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{"string", "hello", "hello"},
		{"int", 123, "123"},
		{"float", 99.9, "99.9"},
		{"bool", true, "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Cast[string](tt.input)
			if result != tt.expected {
				t.Errorf("Cast[string](%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCastInt 测试泛型转换到整数
func TestCastInt(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected int
	}{
		{"string", "456", 456},
		{"int", 123, 123},
		{"float", 789.99, 789},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Cast[int](tt.input)
			if result != tt.expected {
				t.Errorf("Cast[int](%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCastInt64 测试泛型转换到int64
func TestCastInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected int64
	}{
		{"string", "789", 789},
		{"int", 123, 123},
		{"float", 456.78, 456},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Cast[int64](tt.input)
			if result != tt.expected {
				t.Errorf("Cast[int64](%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCastFloat64 测试泛型转换到float64
func TestCastFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"string", "123.45", 123.45},
		{"int", 100, 100.0},
		{"float", 99.99, 99.99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Cast[float64](tt.input)
			if result != tt.expected {
				t.Errorf("Cast[float64](%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCastBool 测试泛型转换到布尔值
func TestCastBool(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"string_true", "true", true},
		{"string_false", "false", false},
		{"int_1", 1, true},
		{"int_0", 0, false},
		{"bool", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Cast[bool](tt.input)
			if result != tt.expected {
				t.Errorf("Cast[bool](%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCastMap 测试泛型转换到map
func TestCastMap(t *testing.T) {
	validMap := map[string]any{"key": "value"}
	result := Cast[map[string]any](validMap)

	if len(result) != 1 || result["key"] != "value" {
		t.Errorf("Cast[map[string]any] failed")
	}
}

// TestCastSlice 测试泛型转换到slice
func TestCastSlice(t *testing.T) {
	validSlice := []any{1, 2, 3}
	result := Cast[[]any](validSlice)

	if len(result) != 3 {
		t.Errorf("Cast[[]any] failed, got length %d", len(result))
	}
}

// TestCastStruct 测试泛型转换到struct
func TestCastStruct(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	input := map[string]any{
		"name": "Alice",
		"age":  30,
	}

	result := Cast[Person](input)

	if result.Name != "Alice" || result.Age != 30 {
		t.Errorf("Cast[Person] failed, got %+v", result)
	}
}

// BenchmarkCastInt 性能测试：整数转换
func BenchmarkCastInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cast[int]("123")
	}
}

// BenchmarkCastString 性能测试：字符串转换
func BenchmarkCastString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cast[string](123)
	}
}

// BenchmarkCastFloat 性能测试：浮点数转换
func BenchmarkCastFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cast[float64]("123.45")
	}
}

// BenchmarkCastStruct 性能测试：结构体转换
func BenchmarkCastStruct(b *testing.B) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	input := map[string]any{
		"name": "Alice",
		"age":  30,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cast[Person](input)
	}
}
