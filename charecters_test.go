package randomtext_test

import (
	"strings"
	"testing"

	"github.com/kishaningithub/randomtext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCharGeneratorShouldGenerateRandomCharecters(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	wordsByteArr, err := randomtext.Asset("assets/charecters.list")
	require.Nil(err)
	expectedCharecters := strings.Split(string(wordsByteArr), "\n")

	generator := randomtext.CharGenerator()
	assert.Contains(expectedCharecters, generator())
}
