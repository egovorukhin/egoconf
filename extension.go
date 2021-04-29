package egoconf

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

type Extension string

const (
	NONE Extension = ""
	JSON Extension = ".json"
	XML  Extension = ".xml"
	YAML Extension = ".yaml"
	YML  Extension = ".yml"
)

//Возвращаем расширение файла
func (ext Extension) String() string {
	return string(ext)
}

func (ext Extension) marshal(v interface{}) ([]byte, error) {

	switch ext {
	case YAML, YML:
		return yaml.Marshal(v)
	case JSON:
		return json.Marshal(v)
	case XML:
		return xml.Marshal(v)
	}

	return nil, errors.New("Extension is not correctly")
}

func (ext Extension) unmarshal(b []byte, v interface{}) error {

	switch ext {
	case YAML, YML:
		return yaml.Unmarshal(b, v)
	case JSON:
		return json.Unmarshal(b, v)
	case XML:
		return xml.Unmarshal(b, v)
	}

	return errors.New("Extension is not correctly")
}

//Получаем расширение файла и соотносим его с константой
func getFileExtension(path string) Extension {
	switch filepath.Ext(path) {
	case JSON.String():
		return JSON
	case XML.String():
		return XML
	case YAML.String():
		return YAML
	case YML.String():
		return YML
	default:
		return NONE
	}
}
