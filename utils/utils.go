package utils

func If(condition bool, valueIfTrue interface{}, valueIfFalse interface{}) interface{} {
	if condition {
		return valueIfTrue
	}
	return valueIfFalse
}
