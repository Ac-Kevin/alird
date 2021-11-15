package config

import (
	"fmt"
)

type DoMain struct {
	Name string
	Host []string
}

func (cfg *Config) GetDoMain() ([]DoMain, error) {
	secs, err := cfg.Ini.SectionsByName("domain")
	if err != nil {
		return nil, fmt.Errorf("ini read domain error:%s", err.Error())
	}
	var datas []DoMain
	for _, sec := range secs {
		if !sec.HasKey("name") {
			return nil, fmt.Errorf("ini domain no have name")
		}
		if !sec.HasKey("host") {
			return nil, fmt.Errorf("ini domain no have host")
		}
		name, _ := sec.GetKey("name")
		host, _ := sec.GetKey("host")
		datas = append(datas, DoMain{Name: name.String(), Host: host.Strings(";")})
	}
	return datas, nil
}
