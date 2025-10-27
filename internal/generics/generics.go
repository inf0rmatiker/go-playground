package generics

import "cmp"

func Compare[K cmp.Ordered](a, b K, operator string) bool {
	switch operator {
	case "gt":
		return a > b
	case "ge":
		return a >= b
	case "lt":
		return a < b
	case "le":
		return a <= b
	case "eq":
		return a == b
	case "ne":
		return a != b
	default:
		return false
	}
}
