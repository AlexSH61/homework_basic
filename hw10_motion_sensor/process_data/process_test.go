package processdataroutine

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProcessDataRoutine(t *testing.T) {
	testData := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	dataChannel := make(chan float64, len(testData))
	processedDataChannel := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(1)

	go ProcessDataRoutine(dataChannel, processedDataChannel, &wg)

	for _, data := range testData {
		dataChannel <- data
	}

	close(dataChannel)

	timeout := time.After(5 * time.Second)

	go func() {
		wg.Wait()
		close(processedDataChannel)
	}()

	select {
	case <-timeout:
		t.Fatal("Test timeout exceeded")
	default:
	}

	processedData := make([]float64, 0, 60)
	for avg := range processedDataChannel {
		processedData = append(processedData, avg)
	}

	expectedProcessedData := []float64{5.5, 15.5}
	assert.ElementsMatch(t, expectedProcessedData, processedData)
}
