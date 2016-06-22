package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

type monitor struct {
	Monitor []service `json:"monitor"`
}

type service struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func CheckHealth() error {
	monitor := monitor{}

	config, err := os.Open("config.json")
	if err != nil {
		return errors.New("Error opening config file: " + err.Error())
	}

	jsonParser := json.NewDecoder(config)
	err = jsonParser.Decode(&monitor)
	if err != nil {
		return errors.New("Error parsing config file: " + err.Error())
	}

	for i := range monitor.Monitor {
		request, err := http.NewRequest("GET", monitor.Monitor[i].Address, nil)
		request.Header.Set("Accept", "application/json")
		request.Header.Set("Authorization", "Bearer "+os.Getenv("WSO2_TOKEN"))

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return err
		}

		log.Println("Checking " + monitor.Monitor[i].Name + ": " + strconv.Itoa(response.StatusCode))

		if response.StatusCode != 200 {
			log.Println(monitor.Monitor[i].Name, response.StatusCode)
		}
	}

	return nil
}
