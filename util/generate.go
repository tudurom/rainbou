package util

import (
	"bytes"
	"html/template"
)

// The Context struct holds the information a template needs
type Context struct {
	Metadata Metadata
	Colors   map[string]*Color
	ColorMap map[string]map[string]interface{}
}

// Generates a Context from a given ColorScheme
func GenerateContext(scheme ColorScheme) Context {
	var ret Context
	ret.ColorMap = make(map[string]map[string]interface{})
	ret.Metadata, ret.Colors = scheme.Metadata, scheme.Colors

	for index, color := range scheme.Colors {
		ret.ColorMap[index] = make(map[string]interface{})
		ret.ColorMap[index]["hex"] = color.Hex()
		ret.ColorMap[index]["hexbgr"] = color.HexBGR()
		ret.ColorMap[index]["dhex"] = color.DHex()
		ret.ColorMap[index]["rgb"] = color.RGB()
		ret.ColorMap[index]["srgb"] = color.SRGB()
	}

	return ret
}

// Builds the theme by executing the template with a given Context
func BuildTheme(scheme ColorScheme, templateFileName string) string {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	ctx := GenerateContext(scheme)
	t.Execute(buf, ctx)

	return buf.String()
}
