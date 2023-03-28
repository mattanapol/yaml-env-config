package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	prefix := flag.String("prefix", "", "Prefix for the environment variables")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: yaml2env [--prefix prefix] <input_file>")
		os.Exit(1)
	}

	inputFilePath := flag.Arg(0)
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	var input map[string]interface{}
	err = yaml.Unmarshal(data, &input)
	if err != nil {
		panic(err)
	}

	envContent := yamlToEnv(input, *prefix)
	// get file name without extension
	fileName := strings.TrimSuffix(path.Base(inputFilePath), path.Ext(inputFilePath))
	outputFilePath := path.Base(fileName) + ".env"
	err = os.WriteFile(outputFilePath, []byte(envContent), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transformed YAML to ENV successfully!")
}

func yamlToEnv(input map[string]interface{}, parentKey string) string {
	var buffer bytes.Buffer
	for key, value := range input {
		key = strings.ReplaceAll(key, "-", "_")
		if parentKey != "" {
			key = parentKey + "_" + key
		}
		key = strings.ToUpper(key)

		switch reflect.ValueOf(value).Kind() {
		case reflect.Map:
			subMap := value.(map[string]interface{})
			buffer.WriteString(strings.TrimRight(yamlToEnv(subMap, key), "\n") + "\n\n")
		default:
			// if value is nil, we set it to an empty string
			if value == nil {
				value = ""
			}
			buffer.WriteString(fmt.Sprintf("%s=%v\n", key, value))
		}
	}
	return buffer.String()
}
