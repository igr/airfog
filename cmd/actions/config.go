package actions

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Host     string
	Scheme   string
	Username string
	Password string
}

func loadConfig() Configuration {
	file, _ := os.Open("./airfog.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	{
		err := decoder.Decode(&configuration)
		if err != nil {
			fmt.Println("Configuration error:", err)
		}
	}
	return configuration
}
