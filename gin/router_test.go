package GIN

import (
	"testing"
)

func TestRouter(t *testing.T) {
	r := Router(true)
	r.Run(":9999")
}
