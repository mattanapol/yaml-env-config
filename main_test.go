package main

import (
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

	actual := yamlToEnv(input, "")
	assertEqual(t, actual["APP_NAME"], "pim-service")
	assertEqual(t, actual["APP_PAGE_SIZE"], "10")
	assertEqual(t, actual["LOGGING_LEVEL"], "debug")
	assertEqual(t, actual["LOGGING_JSON_FORMAT"], "false")
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, actual)
	}
}
