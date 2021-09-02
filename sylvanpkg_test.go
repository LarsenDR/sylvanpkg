package sylvanpkg

import (
	"testing"
)

func TestCreateplot(t *testing.T) {
	result := createplot()
	expected := "hello"
	if result != expected {
		t.Errorf("createplot() test returned an unexpected result: got %v want %v", result, expected)
	}
}
