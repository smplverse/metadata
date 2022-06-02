package config

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
)

type Base struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ExternalUrl string `json:"external_url,omitempty"`
	Image       string `json:"image,omitempty"`
}

func Get() (Base, error) {
	ex, err := os.Executable()
	if err != nil {
		return Base{}, err
	}
	ex = filepath.ToSlash(ex)
	var file *os.File
	file, err = os.Open(path.Join(path.Dir(ex) + "/" + "config.json"))
	if err != nil {
		return Base{}, err
	}

	defer file.Close()
	var config Base
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Base{}, err
	}
	return config, nil
}
