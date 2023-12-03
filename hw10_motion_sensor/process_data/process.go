package processdataroutine

import "sync"

func ProcessDataRoutine(dataChannel <-chan float64, processedDataChannel chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum float64
	var count int
	var mu sync.Mutex

	for sensorData := range dataChannel {
		sum += sensorData
		count++

		if count == 10 {
			average := sum / 10
			mu.Lock()
			processedDataChannel <- average
			mu.Unlock()
			mu.Lock()
			count = 0
			sum = 0
			mu.Unlock()
		}
	}
}
