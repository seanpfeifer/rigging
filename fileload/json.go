// Package fileload contains commonly used file loading functions.
// For human-created files (eg, config files), I'd suggest using TOML (https://github.com/BurntSushi/toml),
// specifically for comment support and improved readability.
package fileload

import (
	"encoding/json"
	"os"
)

// JSON reads a file at the given location and attempts to unmarshal it as JSON into the given value pointed to by v.
func JSON(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
