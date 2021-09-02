package sylvanpkg

import (
	"testing"
)

func TestCreateplot(t *testing.T) {
	result := createplot("testproject")
	expected := "testproject"
	if result != expected {
		t.Errorf("createplot() test returned an unexpected result: got %v want %v", result, expected)
	}
}

func TestWriteplot(t *testing.T) {
	result := writegenerate()
	expected := `{"Projectname":"testproject","Ptype":"plantation","Ppattern":["wo"],"Species":["wo"],"Spdensity":[100],"Hardcore":10}`
	if result != expected {
		t.Errorf("createplot() test returned an unexpected result: got %v want %v", result, expected)
	}
}
