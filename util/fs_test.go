package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitFilename(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Filename string
		Name     string
		Ext      string
	}{
		{"", "", ""},
		{"file.go", "file", ".go"},
		{"dir/file.go", "file", ".go"},
		{"noext", "noext", ""},
		{"dir/noext", "noext", ""},
		{"foo.bar.baz", "foo.bar", ".baz"},
		{"dir/foo.bar.baz", "foo.bar", ".baz"},
		{"empty.", "empty", "."},
		{"dir/empty.", "empty", "."},
		{".ignore", ".ignore", ""},
		{"dir/.ignore", ".ignore", ""},
		{".config.db", ".config", ".db"},
		{"dir/.config.db", ".config", ".db"},
	}

	for _, tt := range tests {
		name, ext := SplitFilename(tt.Filename)
		assert.Equal(t, tt.Name, name, "Name should be equal")
		assert.Equal(t, tt.Ext, ext, "Extension should be equal")
	}
}
