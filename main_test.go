package main

import (
	"strings"
	"testing"
)

func TestYamlToEnv(t *testing.T) {
	input := map[string]interface{}{
		"app": map[string]interface{}{
			"name":      "pim-service",
			"page-size": 10,
		},
		"logging": map[string]interface{}{
			"level":       "debug",
			"json-format": false,
		},
	}

	expected := `APP_NAME=pim-service
APP_PAGE_SIZE=10
LOGGING_LEVEL=debug
LOGGING_JSON_FORMAT=false
`

	actual := yamlToEnv(input, "")
	if !compareEnvStrings(expected, actual) {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, actual)
	}
}

func compareEnvStrings(s1, s2 string) bool {
	s1Lines := strings.Split(strings.Trim(s1, "\n"), "\n")
	s2Lines := strings.Split(strings.Trim(s2, "\n"), "\n")

	if len(s1Lines) != len(s2Lines) {
		return false
	}

	for _, line := range s1Lines {
		if !strings.Contains(strings.Join(s2Lines, "\n"), line) {
			return false
		}
	}

	return true
}
