package egoconf

import (
	"fmt"
	"testing"
)

type Config struct {
	Ports    Ports    `json:"ports"`
	Postgres Postgres `json:"postgres"`
}

type Ports struct {
	Http  string `json:"http"`
	Https string `json:"https"`
	Udp   string `json:"udp"`
}

type Postgres struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SSL      bool   `json:"ssl"`
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
	err := Save("config", YML, c)
	if err != nil {
		fmt.Println(err)
	}

	c1 := Config{}
	err = Load("config", &c1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c1)

}
