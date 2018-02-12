package config

import (
	"os"
	"log"
	"encoding/json"
	"structure"
)

type Config struct {
	Service		string `json:"service"`
	Host		string `json:"host"`
	Port 		uint16 `json:"port"`
	Heartbeat	uint16 `json:"heartbeat"`
	Strategy    string `json:"strategy"`
	MaxProcess  int    `json:"maxprocessor"`
	Backends 	[]structure.Backend `json:"backends"`
}

func Load(fileName string) (*Config, error) {
	var config Config
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("load config faild:", err)
	} else {
		buff := make([]byte, 1024)
		end, _ := file.Read(buff)
		err = json.Unmarshal(buff[:end], &config)
		if err != nil {
			log.Println("decode json config failed:", err)
		}
	}
	log.Println("success load config file:", fileName)
	return &config, err
}

