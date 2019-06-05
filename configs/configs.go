package configs

import (
	"fmt"
	"log"
	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
	}
}

func New() {
	t := T{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t.B.RenamedC)
}