package main

import (
	"fmt"

	processdataroutine "github.com/AlexSH61/homework_basic/hw10_motion_sensor/process_data"
	sensordata "github.com/AlexSH61/homework_basic/hw10_motion_sensor/sensor"
)

func main() {
	dataChannel := sensordata.SensorRoutine()
	processedDataChannel := make(chan float64)

	go processdataroutine.ProcessDataRoutine(dataChannel, processedDataChannel)

	for processedData := range processedDataChannel {
		fmt.Printf("Processed Data: %.2f\n", processedData)
	}
}
