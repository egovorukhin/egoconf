package egoconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"strings"
)

type IExtension interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(b []byte, v interface{}) error
	String() string
}

type Extension struct {
	Name string
}

const (
	JSON = ".json"
	XML  = ".xml"
	YAML = ".yaml"
	YML  = ".yml"
	INI  = ".ini"
)

// Получаем расширение файла и соотносим его с константой
func getFileExtension(path string) (IExtension, error) {

	switch filepath.Ext(path) {
	case YAML, YML:
		return &Yaml{
			Names: []string{YAML, YML},
		}, nil
	case JSON:
		return &Json{
			Name: JSON,
		}, nil
	case XML:
		return &Xml{
			Name: XML,
		}, nil
	case INI:
		return &Ini{
			Name: INI,
			Path: path,
		}, nil
	}

	return nil, errors.New("extension is not correctly")
}

type Json Extension

func (Json) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (Json) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func (j Json) String() string {
	return fmt.Sprintf("name: %s", j.Name)
}

type Yaml struct {
	Names []string
}

func (Yaml) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Yaml) Unmarshal(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

func (y Yaml) String() string {
	return fmt.Sprintf("names: [%s]", strings.Join(y.Names, ","))
}

type Xml Extension

func (Xml) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Xml) Unmarshal(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

func (x Xml) String() string {
	return fmt.Sprintf("name: %s", x.Name)
}

type Ini struct {
	Name string
	Path string
}

func (Ini) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}

func (i Ini) Unmarshal(b []byte, v interface{}) error {
	return ini.MapTo(v, i.Path)
}

func (i Ini) String() string {
	return fmt.Sprintf("name: %s, path: %s", i.Name, i.Path)
}
