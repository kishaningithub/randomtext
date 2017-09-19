package randomtext_test

import (
	"testing"

	"github.com/kishaningithub/randomtext"
	"github.com/stretchr/testify/assert"
)

func TestZeroGeneratorShouldGenerateZeros(t *testing.T) {
	assert := assert.New(t)

	generator := randomtext.ZeroGenerator()

	for i := 0; i < 100; i++ {
		assert.Equal("0", generator())
	}
}
