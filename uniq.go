package uniq

import (
	"golang.org/x/exp/slices"
)

func Uniq[T comparable](orgArray []T) (resultArray []T) {
	for _, elm := range orgArray {
		if !slices.Contains(resultArray, elm) {
			resultArray = append(resultArray, elm)
		}
	}
	return
}

func Version() (version string) {
	version = "1.0.0"
	return
}
