package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestReplaceDollar(t *testing.T) {
	n1 := Replacer("a")
	assert.Equal(t, "a", n1)

	n2 := Replacer("a $$( abc ) b")
	assert.Equal(t, "a HA( abc ) b", n2)

}

func TestClosingBrace(t *testing.T) {
	s1 := []rune("hallo ) hallo")
	i := FindClosingBrace(s1, 0)
	assert.Equal(t, 6, i)

	s2 := []rune("hallo hallo")
	i2 := FindClosingBrace(s2, 0)
	assert.Equal(t, -1, i2)

	s3 := []rune("hallo ( ) hallo ) hallo")
	i3 := FindClosingBrace(s3, 0)
	assert.Equal(t, 16, i3)

}
