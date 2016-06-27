package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Hospital struct {
	MaxCheckups int       `json:"max_checkups"` // The number of times to allow a failure before reporting to Slack
	Patients    []patient `json:"patients"`
}

type patient struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Healthy  bool   `json:"healthy"`
	Checkups int    `json:"checkups"`
}

func LoadConfig(doctor *Hospital) error {
	config, err := os.Open("config.json")
	if err != nil {
		return errors.New("Error opening config file: " + err.Error())
	}

	jsonParser := json.NewDecoder(config)
	err = jsonParser.Decode(&doctor)
	if err != nil {
		return errors.New("Error parsing config file: " + err.Error())
	}

	for i := range doctor.Patients {
		doctor.Patients[i].Healthy = true // Set all health values to true initially
	}

	return nil
}

func CheckHealth(doctor *Hospital) {
	log.Println("----- Checking -----")

	for i := range doctor.Patients {
		request, _ := http.NewRequest("GET", doctor.Patients[i].Address, nil)
		request.Header.Set("Accept", "application/json")
		request.Header.Set("Authorization", "Bearer "+os.Getenv("WSO2_TOKEN"))

		client := &http.Client{}
		client.Timeout = time.Duration(2 * time.Second) // Set the timeout length to 2 seconds

		response, err := client.Do(request)
		if err != nil { // If our health request times out
			doctor.Patients[i] = sickResponse(&doctor.Patients[i], doctor, "timeout")
		} else {
			if response.StatusCode != 200 { // If we get a bad response back
				doctor.Patients[i] = sickResponse(&doctor.Patients[i], doctor, "bad response ("+strconv.Itoa(response.StatusCode)+")")
			} else { // If we get a good response
				doctor.Patients[i] = healthyResponse(&doctor.Patients[i], doctor)
			}
		}
	}
}

func sickResponse(patient *patient, doctor *Hospital, cause string) patient {
	if patient.Healthy == true {
		if patient.Checkups < doctor.MaxCheckups {
			patient.Checkups++
		} else {
			PostToSlack(patient.Name + " is sick")
			patient.Healthy = false
			patient.Checkups = 0
		}
	} else {
		patient.Checkups = 0
	}

	log.Printf("%+v", patient)

	return *patient
}

func healthyResponse(patient *patient, doctor *Hospital) patient {
	if patient.Healthy == false {
		if patient.Checkups < doctor.MaxCheckups {
			patient.Checkups++
		} else {
			PostToSlack(patient.Name + " is healthy")
			patient.Healthy = true
			patient.Checkups = 0
		}
	} else {
		patient.Checkups = 0
	}

	log.Printf("%+v", patient)

	return *patient
}
