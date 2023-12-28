package module

import (
	"fmt"
	"os"

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
			config[fullKey] = fmt.Sprintf("%s", value)
		}
	}
}
