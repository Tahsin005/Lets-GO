package main

import (
	"fmt"
)

type jekono interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Area[T jekono](r T) T {
	switch v := any(r).(type) {
	case string:
		return any(v + v).(T)
	case int:
		return any(v * v).(T)
	case int8:
		return any(v * v).(T)
	case int16:
		return any(v * v).(T)
	case int32:
		return any(v * v).(T)
	case int64:
		return any(v * v).(T)
	case float32:
		return any(v * v).(T)
	case float64:
		return any(v * v).(T)
	default:
		var zero T
		return zero
	}
}

func main() {
	fmt.Println(Area(10))      // 100
	fmt.Println(Area(10.5))    // 110.25
	fmt.Println(Area("Go"))    // GoGo
}
