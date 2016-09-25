package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Filename string
		Expected string
	}{
		{"", ""},
	}

	for _, tt := range tests {
		actual := CleanName(tt.Filename)
		assert.Equal(t, tt.Expected, actual, "Cleaned name should be equal")
	}
}
