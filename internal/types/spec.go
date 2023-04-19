package types

import (
	"fmt"
	"log"
	"time"

	yml "gopkg.in/yaml.v3"
)

type Specs struct {
	Version string `yaml:"version"`
	Project struct {
		Name       string `yaml:"name"`
		Created_at string `yaml:"created_at"`
	}
	Services []Service
}

func NewSpecs(name string, services []Service, databases []Service) *Specs {
	specs := Specs{}
	specs.Version = "v1.0.0"
	specs.Project.Name = name
	specs.Project.Created_at = NewDate()
	specs.Services = services
	return &specs
}

func NewDate() string {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%d %s %d", d, m, y)
}

func ToYamlString(specs *Specs) ([]byte, error) {
	yaml, err := yml.Marshal(&specs)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return yaml, nil
}
