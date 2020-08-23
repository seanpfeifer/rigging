// Package fileload contains commonly used file loading functions.
package fileload

import (
	"encoding/json"
	"io/ioutil"
)

// JSON reads a file at the given location and attempts to unmarshal it as JSON into the given value pointed to by v.
func JSON(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
