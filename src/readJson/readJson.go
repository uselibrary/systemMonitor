package readJson

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Name     string `json:"name"`
	Telegram struct {
		Token  string `json:"token"`
		ChatID string `json:"chat_id"`
	} `json:"telegram"`
	Disk   string `json:"disk"`
	Status struct {
		CPU              float64 `json:"cpu"`
		DiskPercentage   float64 `json:"diskpercentage"`
		Network          float64 `json:"network"`
		MemoryPercentage float64 `json:"memorypercentage"`
	} `json:"status"`
}

// read config file
// variable: file path
// return: config struct
func ReadConfig(path string) *config {
	// open config file
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	// decode config file
	decoder := json.NewDecoder(file)
	config := new(config)
	err = decoder.Decode(config)
	if err != nil {
		log.Println(err)
		return nil
	}
	return config
}