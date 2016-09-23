package torus

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestConvertInfoHash(t *testing.T) {
	var b [20]byte
	if len, _ := rand.Read(b[:]); len != 20 {
		t.Fail()
	}

	res := convertInfoHash(b)
	if res != fmt.Sprintf("%x", b) {
		t.Fail()
	}
}
