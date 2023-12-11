package sensordata

import (
	"testing"
	"time"
)

func TestSensorRoutine(t *testing.T) {
	dataChannel := SensorRoutine()
	timeout := 60 * time.Second
	select {
	case <-time.After(timeout):
		t.Fatalf("Test timed out after %s", timeout)
	case <-dataChannel:
	}
}
