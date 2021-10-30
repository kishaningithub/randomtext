package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/kishaningithub/randomtext"
	"github.com/pkg/profile"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sizePtr := flag.String("size", "1MB", "Size of generated random text in KB, MB, GB, TB")
	typePtr := flag.String("type", "chars", "Type of text to be generated - chars, words, zeros")
	profileTypePtr := flag.String("profile", "", "Type of profile - cpu,mem")
	flag.Parse()
	profileType := *profileTypePtr
	switch profileType {
	case "cpu":
		defer profile.Start(profile.ProfilePath(".")).Stop()
	case "mem":
		defer profile.Start(profile.ProfilePath("."), profile.MemProfile).Stop()
	}
	sizeInBytes, err := parseSize(strings.ToLower(*sizePtr))
	handleErr(err)
	generator, err := parseType(strings.ToLower(*typePtr))
	handleErr(err)
	randomtext.Generate(sizeInBytes, generator, os.Stdout)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseSize(size string) (int, error) {
	reg := regexp.MustCompile("\\d+")
	givenSize, err := strconv.Atoi(reg.FindString(size))
	if err != nil {
		return 0, errors.New("invalid size")
	}
	if strings.Contains(size, "kb") {
		return givenSize * 1024, nil
	}
	if strings.Contains(size, "mb") {
		return givenSize * 1024 * 1024, nil
	}
	if strings.Contains(size, "gb") {
		return givenSize * 1024 * 1024 * 1024, nil
	}
	if strings.Contains(size, "tb") {
		return givenSize * 1024 * 1024 * 1024 * 1024, nil
	}
	return givenSize, nil
}

func parseType(typeStr string) (func() string, error) {
	switch typeStr {
	case "chars":
		return randomtext.CharGenerator(), nil
	case "words":
		return randomtext.WordGenerator(), nil
	case "zeros":
		return randomtext.ZeroGenerator(), nil
	default:
		return nil, errors.New("not a valid value")
	}
}
