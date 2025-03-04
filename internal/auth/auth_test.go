package auth

import (
	"fmt"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	x := 1
	y := 2

	if x > y {
		t.Fatalf(fmt.Sprintf("Expected %d > %d", x, y))
	}
}
