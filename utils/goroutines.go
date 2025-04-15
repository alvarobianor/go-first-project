package utils

import (
	"context"
	"errors"
	"fmt"
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
