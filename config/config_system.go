package config

import "fmt"

type System struct {
	Mode string `yaml:"mode"`
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%s", s.IP, s.Port)
}
