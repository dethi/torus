package util

import "testing"

func TestSplitFilename(t *testing.T) {
	tests := []struct {
		Filename string
		Name     string
		Ext      string
	}{
		{"file.go", "file", ".go"},
		{"foo", "foo", ""},
		{"foo.bar.baz", "foo.bar", ".baz"},
		{"", "", ""},
		{"empty.", "empty", "."},
	}

	for _, tt := range tests {
		name, ext := SplitFilename(tt.Filename)
		if name != tt.Name || ext != tt.Ext {
			t.Errorf("SplitFilename(%s) => (%s, %s), want (%s, %s)",
				tt.Filename, name, ext, tt.Name, tt.Ext)
		}
	}
}
