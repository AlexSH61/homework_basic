package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}
	wg.Wait()
	assert.Equal(t, 5, counter, "counter value: 5")
}
