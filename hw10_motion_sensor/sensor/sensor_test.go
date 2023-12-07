package sensordata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSensorRoutine(t *testing.T) {
	dataChannel := make(chan float64, 60)
	go SensorRoutine(dataChannel)

	for i := 0; i < 60; i++ {
		sensorData := <-dataChannel
		assert.InDelta(t, sensorData, 0, 100, "Value out of range")
	}
	close(dataChannel)
	time.Sleep(100 * time.Millisecond)
}
