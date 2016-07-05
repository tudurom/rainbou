package util

import (
	"io/ioutil"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Stores metadata about the color scheme: its name and author
type Metadata struct {
	Name   string
	Author string
}

// A representation of a color scheme
type ColorScheme struct {
	Metadata Metadata
	Colors   map[string]*Color
}

// The structure of a color scheme in the yaml color scheme file
type RawColorScheme struct {
	Metadata Metadata
	Colors   map[string]string
}

// Converts a RawColorScheme to a ColorScheme
func NewColorScheme(raw RawColorScheme) ColorScheme {
	var scheme ColorScheme
	scheme.Colors = make(map[string]*Color)

	r, err := regexp.Compile(".{1,2}")
	if err != nil {
		panic(err)
	}

	scheme.Metadata = raw.Metadata

	for code, value := range raw.Colors {
		scheme.Colors[code] = new(Color)
		for i, color := range r.FindAll([]byte(value), -1) {
			decimalColor, err := strconv.ParseUint(string(color[:]), 16, 8)
			if err != nil {
				panic(err)
			}
			if i == 0 {
				scheme.Colors[code].R = uint8(decimalColor)
			} else if i == 1 {
				scheme.Colors[code].G = uint8(decimalColor)
			} else {
				scheme.Colors[code].B = uint8(decimalColor)
			}
		}
	}

	return scheme
}

// Reads the file and returns a ColorScheme object
func DecodeFromFile(fpath string) ColorScheme {
	contents, err := ioutil.ReadFile(fpath)

	if err != nil {
		panic(err)
	}

	scheme := RawColorScheme{}
	yaml.Unmarshal(contents, &scheme)

	return NewColorScheme(scheme)
}
