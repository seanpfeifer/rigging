package fileload

import (
	"testing"

	. "github.com/seanpfeifer/rigging/assert"
)

const (
	jsonFile = "test_file.json"
	tomlFile = "test_file.toml"
)

type person struct {
	Name string
	Hat  bool
	Mass uint64
}
type wrongPerson struct {
	Name uint64
	Hat  string
	Mass float32
}

func TestJSONLoad(t *testing.T) {
	p, err := JSON[person](jsonFile)
	ExpectedActual(t, nil, err, "parsing person")
	validatePerson(t, p)
}

func TestJSONLoadError(t *testing.T) {
	_, err := JSON[wrongPerson](jsonFile)
	ExpectedActual(t, true, err != nil, "parsing wrongPerson")
}

func TestTOMLLoad(t *testing.T) {
	p, _, err := TOML[person](tomlFile)
	ExpectedActual(t, nil, err, "parsing person")
	validatePerson(t, p)
}

func TestTOMLLoadError(t *testing.T) {
	_, _, err := TOML[wrongPerson](tomlFile)
	ExpectedActual(t, true, err != nil, "parsing wrongPerson")
}

func validatePerson(t *testing.T, p *person) {
	t.Helper()
	ExpectedActual(t, true, p != nil, "non-nil person")
	ExpectedActual(t, "Sean", p.Name, "name")
	ExpectedActual(t, true, p.Hat, "hat")
	ExpectedActual(t, 42, p.Mass, "mass")
}
