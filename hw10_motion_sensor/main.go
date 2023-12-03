package main

import (
	"fmt"
	"sync"

	processdataroutine "github.com/AlexSH61/homework_basic/hw10_motion_sensor/process_data"
	sensorData "github.com/AlexSH61/homework_basic/hw10_motion_sensor/sensor"
)

func main() {
	dataChannel := make(chan float64)
	processedDataChannel := make(chan float64)
	var wg sync.WaitGroup

	wg.Add(2)

	go sensorData.SensorRoutine(dataChannel, &wg)
	go processdataroutine.ProcessDataRoutine(dataChannel, processedDataChannel, &wg)

	go func() {
		wg.Wait()
		close(processedDataChannel)
	}()

	for processedData := range processedDataChannel {
		fmt.Printf("Processed Data: %.2f\n", processedData)
	}
}
