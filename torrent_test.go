package torus

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInfoHash(t *testing.T) {
	t.Parallel()

	var b [20]byte
	if len, _ := rand.Read(b[:]); len != 20 {
		t.Fail()
	}

	expected := fmt.Sprintf("%x", b)
	actual := convertInfoHash(b)
	assert.Equal(t, expected, actual, "Hex representation should be equal")
}
