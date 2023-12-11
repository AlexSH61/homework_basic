package processdataroutine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessDataRoutine(t *testing.T) {
	testData := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	dataChannel := make(chan float64, len(testData))
	processedDataChannel := make(chan float64)

	go ProcessDataRoutine(dataChannel, processedDataChannel)

	for _, data := range testData {
		dataChannel <- data
	}
	close(dataChannel)

	processedData := make([]float64, 0, 60)
	for avg := range processedDataChannel {
		processedData = append(processedData, avg)
	}

	expectedProcessedData := []float64{5.5, 15.5}
	assert.ElementsMatch(t, expectedProcessedData, processedData)
}
