package egoconf

import (
	"fmt"
	"testing"
)

type Config struct {
	Ports    `json:"ports" yaml:"ports" ini:"ports"`
	Postgres `json:"postgres" yaml:"postgres" ini:"postgres"`
}

type Ports struct {
	Http  string `json:"http" yaml:"http" ini:"http"`
	Https string `json:"https" yaml:"https" ini:"https"`
	Udp   string `json:"udp" yaml:"udp" ini:"udp"`
}

type Postgres struct {
	Server   string `json:"server" yaml:"server" ini:"server"`
	Port     string `json:"port" yaml:"port" ini:"port"`
	Username string `json:"username" yaml:"username" ini:"username"`
	Password string `json:"password" yaml:"password" ini:"password"`
	Name     string `json:"name" yaml:"name" ini:"name"`
	SSL      bool   `json:"ssl" yaml:"ssl" ini:"ssl"`
}

func TestLoad(t *testing.T) {

	c := Config{
		Ports: Ports{
			Http:  "80",
			Https: "443",
			Udp:   "6565",
		},
		Postgres: Postgres{
			Server:   "localhost",
			Port:     "5432",
			Username: "user",
			Password: "pass",
			Name:     "db",
			SSL:      false,
		},
	}

	//Yaml
	err := Save("config.yml", c)
	if err != nil {
		t.Error(err)
	}

	c1 := Config{}
	err = Load("config.yml", &c1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Yml: %v\n", c1)

	//Json
	err = Save("config.json", c)
	if err != nil {
		t.Error(err)
	}

	err = Load("config.json", &c1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Json: %v\n", c1)

	//xml
	err = Save("config.xml", c)
	if err != nil {
		t.Error(err)
	}

	err = Load("config.xml", &c1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Xml: %v\n", c1)

	//ini
	var cfg Config
	err = Load("config.ini", &cfg)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Ini: %v\n", cfg)
}
