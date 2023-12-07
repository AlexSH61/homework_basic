package sensordata

import (
	"math/rand"
	"time"
)

func SecureRandomFloat64() float64 {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Float64()
}

func SensorRoutine(dataChannel chan<- float64) {
	defer close(dataChannel)
	for i := 0; i < 60; i++ {
		sensorData := SecureRandomFloat64() * 100
		dataChannel <- sensorData
		time.Sleep(time.Second)
	}
}
