package module

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

var config map[string]string

func InitConfigs() {
	filename := "resources/application.yml"
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var result map[interface{}]interface{}
	err = yaml.UnmarshalStrict(data, &result)
	if err != nil {
		panic(err)
	}

	config = make(map[string]string)
	initKeys("", result)
}

func formatKey(prefix string, key interface{}) string {
	if prefix == "" {
		return fmt.Sprintf("%v", key)
	}
	return fmt.Sprintf("%s.%v", prefix, key)
}

func initKeys(prefix string, data map[interface{}]interface{}) {
	for key, value := range data {
		fullKey := formatKey(prefix, key)

		if v, ok := value.(map[interface{}]interface{}); ok {
			initKeys(fullKey, v)
		} else {
			config[fullKey] = formatValue(value)
		}
	}
}


func formatValue(value interface{}) string {
	strValue := fmt.Sprintf("%s", value)
	re := regexp.MustCompile("\\$\\{([^:]+)(:([^\\}]+))?\\}")
	result := re.ReplaceAllStringFunc(strValue, func(match string) string {
		matches := re.FindStringSubmatch(match)
		envVar := matches[1]
		defaultVal := matches[3]

		if val, exists := os.LookupEnv(envVar); exists {
			return val
		}
		return defaultVal
	})

	return result
}