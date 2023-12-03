package sensordata

import (
	"sync"
	"time"
)

func SensorRoutine(dataChannel chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataChannel)

	for i := 1; i <= 60; i++ {
		sensorData := float64(i * 2)
		dataChannel <- sensorData
		time.Sleep(time.Second)
	}
}
