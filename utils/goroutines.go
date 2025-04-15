package utils

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func SimulateApiCalls(url string, times int8) {

	start := time.Now()

	for i := range times {
		resp, error := http.Get(url)
		if error != nil {
			panic(error)
		}
		fmt.Println("API called", i+1, "times")
		defer resp.Body.Close()
	}

	elapsed := time.Since(start)
	fmt.Printf("Tempo de execução: %s\n", elapsed)
}

func SimulateApiCallsWithGoroutines(url string, times int8) {
	start := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(int(times))

	for i := range times {
		go func() {
			resp, error := http.Get(url)
			if error != nil {
				panic(error)
			}
			fmt.Println("API called", i+1, "times")
			defer resp.Body.Close()
			defer wg.Done()
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Tempo de execução: %s\n", elapsed)

}
