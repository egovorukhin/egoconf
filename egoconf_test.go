package egoconf

import (
	"fmt"
	"testing"
)

type Config struct {
	Ports    Ports    `json:"ports" yaml:"ports"`
	Postgres Postgres `json:"postgres" yaml:"postgres"`
}

type Ports struct {
	Http  string `json:"http" yaml:"http"`
	Https string `json:"https" yaml:"https"`
	Udp   string `json:"udp" yaml:"udp"`
}

type Postgres struct {
	Server   string `json:"server" yaml:"server"`
	Port     string `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Name     string `json:"name" yaml:"name"`
	SSL      bool   `json:"ssl" yaml:"ssl"`
}

func TestLoad(t *testing.T) {

	c := Config{
		Ports: Ports{
			Http:  "80",
			Https: "443",
			Udp:   "6565",
		},
		Postgres: Postgres{
			Server:   "server",
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

	//Xml
	err = Save("config.xml", c)
	if err != nil {
		t.Error(err)
	}

	err = Load("config.xml", &c1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Xml: %v\n", c1)
}
