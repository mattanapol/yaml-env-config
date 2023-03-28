package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	prefix := flag.String("prefix", "", "Prefix for the environment variables")
	outputFormat := flag.String("format", "env", "Output format (env, yaml)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: yaml2env [--prefix prefix] [--format format] <input_file>")
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

	envMap := yamlToEnv(input, *prefix)
	// get file name without extension
	fileName := strings.TrimSuffix(path.Base(inputFilePath), path.Ext(inputFilePath))
	var outputFilePath, content string
	keyValue := createSortedKeyValueFromMap(envMap)
	if *outputFormat == "yaml" {
		outputFilePath = fileName + "_env.yaml"
		content = mapToYaml(keyValue)
	} else {
		outputFilePath = fileName + ".env"
		content = mapToEnv(keyValue)
	}
	err = os.WriteFile(outputFilePath, []byte(content), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transformed YAML to ENV successfully!")
}

func yamlToEnv(input map[string]interface{}, parentKey string) map[string]string {
	var result = make(map[string]string)
	for key, value := range input {
		key = strings.ReplaceAll(key, "-", "_")
		if parentKey != "" {
			key = parentKey + "_" + key
		}
		key = strings.ToUpper(key)

		switch reflect.ValueOf(value).Kind() {
		case reflect.Map:
			subMap := value.(map[string]interface{})
			// append the result of the submap to the result
			for k, v := range yamlToEnv(subMap, key) {
				result[k] = v
			}
		default:
			// if value is nil, we set it to an empty string
			if value == nil {
				value = ""
			}
			result[key] = fmt.Sprintf("%v", value)
		}
	}
	return result
}

func mapToEnv(input []KeyValue) string {
	var buffer bytes.Buffer
	for _, entry := range input {
		buffer.WriteString(fmt.Sprintf("%s=%v\n", entry.Key, entry.Value))
	}
	return buffer.String()
}

func mapToYaml(input []KeyValue) string {
	var buffer bytes.Buffer
	for _, entry := range input {
		buffer.WriteString(fmt.Sprintf("%s: \"%v\"\n", entry.Key, entry.Value))
	}
	return buffer.String()
}

func createSortedKeyValueFromMap(input map[string]string) []KeyValue {
	// Extract keys from map
	var keys []string
	for k := range input {
		keys = append(keys, k)
	}

	// Sort keys alphabetically
	sort.Strings(keys)

	var result []KeyValue
	for _, key := range keys {
		result = append(result, KeyValue{
			Key:   key,
			Value: input[key],
		})
	}
	return result
}

type KeyValue struct {
	Key   string
	Value string
}
