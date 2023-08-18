package utils

import "github.com/samber/lo"

func Ptr[T any](v T) *T {
	return &v
}

func PtrString(v string) *string {
	if v == "" {
		return nil
	}

	return &v
}

func EmptyPtr[T comparable](v T) *T {
	if lo.IsEmpty(v) {
		return nil
	}

	return &v
}

func ParseEmptyPtr[T any](v *T) T {
	if v == nil {
		return lo.Empty[T]()
	}

	return *v
}
