package helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHealth(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json") // Respond with something to simulate a 200
	}))
	defer server.Close()

	doctor := Hospital{MaxCheckups: 3}

	doctor.Patients = append(doctor.Patients, patient{Name: "Stuff", Address: server.URL, Healthy: false})

	CheckHealth(&doctor)

	if doctor.Patients[0].Healthy != false && doctor.Patients[0].Checkups != 1 {
		test.Fail()
	}
}

func TestSickResponse(test *testing.T) {
	doctor := Hospital{MaxCheckups: 3}
	patient := patient{Healthy: true, Checkups: 0}

	sickResponse(&patient, &doctor, "Fake cause")

	test.Logf("%+v", patient)

	if patient.Healthy != true || patient.Checkups != 1 {
		test.Fail()
	}
}

func TestHealthyResponse(test *testing.T) {
	doctor := Hospital{MaxCheckups: 3}
	patient := patient{Healthy: false, Checkups: 0}

	healthyResponse(&patient, &doctor)

	test.Logf("%+v", patient)

	if patient.Healthy != false || patient.Checkups != 1 {
		test.Fail()
	}
}
