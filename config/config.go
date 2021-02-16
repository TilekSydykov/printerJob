package config

import (
	"encoding/json"
	"io/ioutil"
	"printsServer/filesystem"
)

const PrinterAddr = "NPI6F9308"
const PrinterPort = "9100"

const filePath = "./printer.conf.json"

type Config struct {
	Ip        string `json:"ip"`
	LocalGate string `json:"local_gate"`
}

func (c *Config) GetConfig() error {
	file, err := filesystem.RetrieveROM(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, c)
}

func (c *Config) WriteConfig() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}
