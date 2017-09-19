package randomtext_test

import (
	"strings"
	"testing"

	"github.com/kishaningithub/randomtext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWordGeneratorShouldGenerateRandomWords(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	wordsByteArr, err := randomtext.Asset("assets/words.list")
	require.Nil(err)
	expectedWords := strings.Split(string(wordsByteArr), "\n")

	generator := randomtext.WordGenerator()
	assert.Contains(expectedWords, generator())
}
