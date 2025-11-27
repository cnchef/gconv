package gconv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

/* ------------------------------
   基础类型转换函数
--------------------------------*/

func ToString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case json.Number:
		return val.String()
	case float64, float32, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, bool:
		return fmt.Sprintf("%v", val)
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}

func ToInt(v any) int {
	switch val := v.(type) {
	case int:
		return val
	case int64:
		return int(val)
	case float64:
		return int(val)
	case json.Number:
		i, _ := val.Int64()
		return int(i)
	case string:
		i, _ := strconv.Atoi(val)
		return i
	default:
		return 0
	}
}

func ToFloat(v any) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case json.Number:
		f, _ := val.Float64()
		return f
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	default:
		return 0
	}
}

func ToBool(v any) bool {
	switch val := v.(type) {
	case bool:
		return val
	case string:
		b, _ := strconv.ParseBool(val)
		return b
	case int:
		return val != 0
	case float64:
		return val != 0
	case json.Number:
		f, _ := val.Float64()
		return f != 0
	default:
		return false
	}
}

func ToMap(v any) map[string]any {
	if m, ok := v.(map[string]any); ok {
		return m
	}
	return map[string]any{}
}

func ToSlice(v any) []any {
	if arr, ok := v.([]any); ok {
		return arr
	}
	return []any{}
}

/* ------------------------------
   ⭐ 泛型万能转换 Cast[T]
--------------------------------*/

func Cast[T any](v any) T {
	var zero T
	t := any(zero)

	switch t.(type) {

	case string:
		return any(ToString(v)).(T)

	case int:
		return any(ToInt(v)).(T)

	case int64:
		return any(int64(ToInt(v))).(T)

	case float64:
		return any(ToFloat(v)).(T)

	case bool:
		return any(ToBool(v)).(T)

	case map[string]any:
		return any(ToMap(v)).(T)

	case []any:
		return any(ToSlice(v)).(T)

	default:
		// 如果 T 是 struct，尝试 json 转换
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Map || rv.Kind() == reflect.Slice || rv.Kind() == reflect.Struct {
			b, _ := json.Marshal(v)
			_ = json.Unmarshal(b, &zero)
			return zero
		}

		// 其他情况返回零值
		return zero
	}
}
