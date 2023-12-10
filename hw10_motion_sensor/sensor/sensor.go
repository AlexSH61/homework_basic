package sensordata

import (
	"math/rand"
	"time"
)

func SecureRandomFloat64() float64 {
	source := rand.NewSource(time.Now().UnixNano())
	// #nosec G404
	random := rand.New(source)
	return random.Float64()
}

func SensorRoutine() chan float64 {
	dataChannel := make(chan float64, 60)
	go func() {
		defer close(dataChannel)
		for i := 0; i < 60; i++ {
			sensorData := SecureRandomFloat64() * 100
			dataChannel <- sensorData
			time.Sleep(time.Second)
		}
	}()
	return dataChannel
}
