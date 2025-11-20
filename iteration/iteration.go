package iteration

import "strings"

// The standard library provides the strings.BuilderstringsBuilder type which minimizes memory copying.
// It implements a WriteString method which we can use to concatenate strings

func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// Initial version without strings.Builder
//func Repeat(character string) string {
//	var repeated string
//	for i := 0; i < repeatCount; i++ {
//		repeated += character
//	}
//	return repeated
//}
