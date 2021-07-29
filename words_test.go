package randomtext_test

import (
	"strings"
	"testing"

	"github.com/kishaningithub/randomtext"
	"github.com/stretchr/testify/require"
)

func TestWordGeneratorShouldGenerateRandomWords(t *testing.T) {
	require := require.New(t)
	expectedWords := strings.Split(randomtext.Words, "\n")

	generator := randomtext.WordGenerator()

	require.Contains(expectedWords, generator())
}
