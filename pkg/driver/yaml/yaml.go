package yaml

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadFromFile(file string, into interface{}) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no file exists at path %s", file)
		}
		return fmt.Errorf("could not read file %s: %w", file, err)
	}

	if err := yaml.Unmarshal(bytes, into); err != nil {
		return fmt.Errorf("could not parse file YAML at %s: %w", file, err)
	}
	return nil
}

func WriteToFile(file string, from interface{}) error {
	bytes, err := yaml.Marshal(from)
	if err != nil {
		return fmt.Errorf("could not marshal to YAML: %w", err)
	}
	if err := os.WriteFile(file, bytes, 0777); err != nil {
		return fmt.Errorf("could not write YAML to file %s: %w", file, err)
	}
	return nil
}
