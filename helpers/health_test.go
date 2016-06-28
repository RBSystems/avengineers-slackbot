package helpers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestCheckHealthSuccess(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	doctor := Hospital{MaxCheckups: 3}

	doctor.Patients = append(doctor.Patients, patient{Name: "Stuff", Address: server.URL, Healthy: false})

	CheckHealth(&doctor)

	if doctor.Patients[0].Healthy != false && doctor.Patients[0].Checkups != 1 {
		test.Fail()
	}
}

func TestCheckHealthTimeout(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(3 * time.Second) // Simulate a timeout (anything above two seconds will timeout)
	}))
	defer server.Close()

	doctor := Hospital{MaxCheckups: 3}

	doctor.Patients = append(doctor.Patients, patient{Name: "Stuff", Address: server.URL, Healthy: false})

	CheckHealth(&doctor)

	if doctor.Patients[0].Healthy != false && doctor.Patients[0].Checkups != 1 {
		test.Fail()
	}
}

func TestCheckHealthFail(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	doctor := Hospital{MaxCheckups: 3}

	doctor.Patients = append(doctor.Patients, patient{Name: "Stuff", Address: server.URL, Healthy: false})

	CheckHealth(&doctor)

	if doctor.Patients[0].Healthy != false && doctor.Patients[0].Checkups != 1 {
		test.Fail()
	}
}

func TestSickResponseWhenHealthy(test *testing.T) {
	doctor := Hospital{MaxCheckups: 3}
	patient := patient{Healthy: true, Checkups: 0}

	sickResponse(&patient, &doctor, "Fake cause")

	if patient.Healthy != true || patient.Checkups != 1 {
		test.Fail()
	}
}

func TestSickResponseWhenSick(test *testing.T) {
	doctor := Hospital{MaxCheckups: 3}
	patient := patient{Healthy: false, Checkups: 0}

	sickResponse(&patient, &doctor, "Fake cause")

	if patient.Healthy != false || patient.Checkups != 0 {
		test.Fail()
	}
}

func TestSickResponseWhenHealthyWithMaxCheckups(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	os.Setenv("SLACKBOT_WEBHOOK", server.URL) // Don't post to Slack

	doctor := Hospital{MaxCheckups: 3}
	patient := patient{Healthy: true, Checkups: 3}

	sickResponse(&patient, &doctor, "Fake cause")

	if patient.Healthy != false || patient.Checkups != 0 {
		test.Fail()
	}
}

func TestHealthyResponseWhenHealthy(test *testing.T) {
	doctor := Hospital{MaxCheckups: 3}
	patient := patient{Healthy: true, Checkups: 0}

	healthyResponse(&patient, &doctor)

	if patient.Healthy != true || patient.Checkups != 0 {
		test.Fail()
	}
}

func TestHealthyResponseWhenSick(test *testing.T) {
	doctor := Hospital{MaxCheckups: 3}
	patient := patient{Healthy: false, Checkups: 0}

	healthyResponse(&patient, &doctor)

	if patient.Healthy != false || patient.Checkups != 1 {
		test.Fail()
	}
}

func TestSickResponseWhenSickWithMaxCheckups(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	os.Setenv("SLACKBOT_WEBHOOK", server.URL) // Don't post to Slack

	doctor := Hospital{MaxCheckups: 3}
	patient := patient{Healthy: false, Checkups: 3}

	healthyResponse(&patient, &doctor)

	if patient.Healthy != true || patient.Checkups != 0 {
		test.Fail()
	}
}
