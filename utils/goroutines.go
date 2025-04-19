package utils

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
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

func SimulateServerHttp(times int8) {
	start := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(int(times))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(11 * time.Second)
		fmt.Println("Server called")
	}))
	defer server.Close()

	for i := range times {
		go func(ctx context.Context) {
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, "GET", server.URL, nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("Request timed out")
				}

				return
			}
			defer resp.Body.Close()
			fmt.Println("API called", i+1, "times")
		}(ctx)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Tempo de execução: %s\n", elapsed)
}

func SimulateRace() {

	// Create a channel to know when the race ends
	finishLine := make(chan string)

	startTime := time.Now()

	// Create 3 race cars (goroutines)
	go func() {
		randTime := time.Duration(rand.Intn(10))
		time.Sleep(randTime * time.Second) // Red car is slower
		finishLine <- "Red Car"
	}()

	go func() {
		randTime := time.Duration(rand.Intn(10))
		time.Sleep(randTime * time.Second) // Blue car is faster
		finishLine <- "Blue Car"
	}()

	go func() {
		randTime := time.Duration(rand.Intn(10))
		time.Sleep(randTime * time.Second) // Green car is the slowest
		finishLine <- "Green Car"
	}()

	// Let's see who arrives first!
	for i := 0; i < 3; i++ {
		car := <-finishLine
		fmt.Printf("%s arrived in %d place! In: %s\n", car, i+1, time.Since(startTime))
	}

}
