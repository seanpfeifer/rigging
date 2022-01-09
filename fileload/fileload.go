// Package fileload contains commonly used file loading functions.
// For human-created files (eg, config files), I'd suggest using TOML (https://github.com/BurntSushi/toml),
// specifically for comment support and improved readability.
package fileload

import (
	"encoding/json"
	"os"

	"github.com/BurntSushi/toml"
)

// JSON reads a file at the given location and attempts to unmarshal it as JSON with the given type.
// The generic type adds more convenience so you can simplify file loading down to something like:
//   cfg, err := fileload.JSON[ServiceCfg]("cfg.json")
//   logging.FatalIfError(err, "loading service config")
//
// ie, this saves you a line of explicitly declaring a var, because we'll declare it and return a pointer here.
func JSON[T any](filename string) (*T, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var v T
	err = json.Unmarshal(data, &v)
	return &v, err
}

// TOML reads a file at the given location and attempts to unmarshal it as TOML with the given type.
func TOML[T any](filename string) (*T, toml.MetaData, error) {
	var v T
	meta, err := toml.DecodeFile(filename, &v)
	return &v, meta, err
}
