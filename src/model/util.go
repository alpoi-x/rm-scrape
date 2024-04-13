package model

func AsPointer[T any](val T) *T {
	return &val
}
