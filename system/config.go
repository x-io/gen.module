package system

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

func LoadConfig(config interface{}, file, dirs string) error {

	data, err := os.ReadFile(file)
	if err != nil {
		dir, err := os.Executable()
		if err != nil {
			return err
		}

		exPath := filepath.Dir(dir)
		data, err = os.ReadFile(exPath + dirs + file)
		if err != nil {
			data, err = os.ReadFile("." + dirs + file)
			if err != nil {
				return err
			}
		}
	}

	switch path.Ext(file) {
	case ".yaml":
		return yaml.Unmarshal(data, config)
	case ".json":
		return json.Unmarshal(data, config)
	default:
		return fmt.Errorf("扩展名不符")
	}
}
