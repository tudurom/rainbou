package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tudurom/rainbou/util"
)

func main() {
	var schemeFilePath, templateFilePath string

	if len(os.Args) != 3 {
		printUsage()
	}

	// Add the yaml suffix automatically
	schemeFilePath, templateFilePath = os.Args[1], os.Args[2]
	if !strings.HasSuffix(schemeFilePath, ".yaml") {
		schemeFilePath += ".yaml"
	}

	// Scheme and template files are relative to the db directory
	dir, err := filepath.Abs(os.Getenv("GOPATH") + "/src/github.com/tudurom/rainbou")
	if err != nil {
		panic(err)
	}

	const schemeDir = "db/colors"
	const templateDir = "db/templates"

	if schemeFilePath == "" {
		fmt.Println("You must specify a color scheme file")
		os.Exit(1)
	}

	if templateFilePath == "" {
		fmt.Println("You must specify a template file")
		os.Exit(1)
	}

	var schemeFile string
	absSchemeFile, err := filepath.Abs(schemeFilePath)
	if err != nil {
		panic(err)
	}
	if filepath.IsAbs(schemeFilePath) && fileExists(schemeFilePath) {
		// Case 1: the supplied path is already absolute
		schemeFile = schemeFilePath
	} else if fileExists(dir + "/" + schemeDir + "/" + schemeFilePath) {
		// Case 2: the path is relative to the db dir (only the color scheme was supplied)
		schemeFile = dir + "/" + schemeDir + "/" + schemeFilePath
	} else if fileExists(absSchemeFile) {
		// Case 3: the path is relative to the working directory
		schemeFile = absSchemeFile
	} else {
		fmt.Println("Color scheme file not found")
		os.Exit(1)
	}

	// Same cases here
	var templateFile string
	absTemplateFile, err := filepath.Abs(templateFilePath)
	if err != nil {
		panic(err)
	}
	if filepath.IsAbs(templateFilePath) && fileExists(templateFilePath) {
		templateFile = templateFilePath
	} else if fileExists(dir + "/" + templateDir + "/" + templateFilePath) {
		templateFile = dir + "/" + templateDir + "/" + templateFilePath
	} else if fileExists(absTemplateFile) {
		templateFile = absTemplateFile
	} else {
		fmt.Println("Template file not found")
		os.Exit(1)
	}

	fmt.Print(util.BuildTheme(util.DecodeFromFile(schemeFile), templateFile))
}

func printUsage() {
	fmt.Printf("Usage: %s <scheme_file> <template_file>\n\n", os.Args[0])
	fmt.Println("Where scheme_file is the name or the path of the color scheme file and\ntemplate_file is the name or the path of the color scheme file.")
	os.Exit(1)
}

// Returns true if the given file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
