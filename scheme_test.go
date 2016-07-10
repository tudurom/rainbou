package main

import (
	"reflect"
	"testing"

	"github.com/tudurom/rainbou/util"
)

func is(actual, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func GetRawTestScheme() util.RawColorScheme {
	scheme := util.RawColorScheme{
		util.Metadata{
			"thunder",
			"tudurom",
		},
		map[string]string{
			"bg": "1d2a30",
			"fg": "c0c5ce",
			"cr": "c0c5ce",

			"00": "1d2a30",
			"08": "456472",

			"01": "a66363",
			"09": "c27171",

			"02": "63a690",
			"10": "6dc2a3",

			"03": "a6a663",
			"11": "bfc271",

			"04": "6385a6",
			"12": "719bc2",

			"05": "bf9c86",
			"13": "bf9c86",

			"06": "63a69b",
			"14": "71c2af",

			"07": "c0c5ce",
			"15": "eff1f5",
		},
	}

	return scheme
}

func GetTestScheme() util.ColorScheme {
	return util.NewColorScheme(GetRawTestScheme())
}

func TestContext(t *testing.T) {
	actual := util.GenerateContext(GetTestScheme())

	is(actual.Metadata, GetTestScheme().Metadata, t)
	is(actual.Colors, GetTestScheme().Colors, t)
}

func TestHexColors(t *testing.T) {
	actual := util.GenerateContext(GetTestScheme()).ColorMap
	expected := GetRawTestScheme().Colors

	for code, _ := range actual {
		is(actual[code]["hex"], expected[code], t)
	}
}
