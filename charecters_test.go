package randomtext_test

import (
	"strings"
	"testing"

	"github.com/kishaningithub/randomtext"
	"github.com/stretchr/testify/require"
)

func TestCharGeneratorShouldGenerateRandomCharacters(t *testing.T) {
	require := require.New(t)
	expectedCharacters := strings.Split(randomtext.Characters, "\n")

	generator := randomtext.CharGenerator()

	require.Contains(expectedCharacters, generator())
}
