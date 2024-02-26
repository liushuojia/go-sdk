package utils

import (
	"fmt"
	"testing"
)

// go test -run TestContentType
func TestContentType(t *testing.T) {
	i, _ := GetContentType("1.jpg")
	fmt.Println(i)
	fmt.Println(ContentType2Extension(i))

	a, _ := GetContentType("content-type.txt")
	fmt.Println(a)

}
