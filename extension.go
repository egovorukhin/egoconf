package egoconf

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"gopkg.in/ini.v1"

	"gopkg.in/yaml.v3"
	"path/filepath"
)

type Extension string

const (
	NONE Extension = ""
	JSON           = ".json"
	XML            = ".xml"
	YAML           = ".yaml"
	YML            = ".yml"
	INI            = ".ini"
)

// Возвращаем расширение файла
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
	case INI:
		return ini.MapTo(v, string(b))
	}

	return errors.New("Extension is not correctly")
}

// Получаем расширение файла и соотносим его с константой
func getFileExtension(path string) Extension {
	return Extension(filepath.Ext(path))
	/*	switch filepath.Ext(path) {
		case JSON:
			return JSON
		case XML:
			return XML
		case YAML:
			return YAML
		case YML:
			return YML
		case INI:
			return INI
		default:
			return NONE
		}*/
}
