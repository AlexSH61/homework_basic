package processdataroutine

func ProcessDataRoutine(dataChannel <-chan float64, processedDataChannel chan<- float64) {
	var sum float64
	var count int
	defer close(processedDataChannel)
	for sensorData := range dataChannel {
		sum += sensorData
		count++

		if count == 10 {
			average := (sum) / 10
			processedDataChannel <- average
			count = 0
			sum = 0
		}
	}
}
