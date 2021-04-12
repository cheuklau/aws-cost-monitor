package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type RulesConfig struct {
	Team string `yaml:"team"`
	Ec2 struct {
		Sizes []string `yaml:"sizes,flow"`
		MaxCount []int `yaml:"maxCount,flow"`
	}
}

func (c *RulesConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func Run(filename string) RulesConfig {
	fmt.Println("Retrieving input parameters...")

	// Read yaml input file as data stream
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Parse data stream using defined rules structure
	var rules RulesConfig
	if err := rules.Parse(data); err != nil {
		log.Fatal(err)
	}

	return rules
}