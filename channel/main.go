package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Printf("Error shutting down server: %v\n", err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting server: %v\n", err)
	}
	fmt.Println("Server stopped")
}

func mainGo() {
	start := time.Now()
	for i := range 10 {
		if i%2 == 0 {
			go slowFunction(i)
			continue
		}
		slowFunction(i)
	}
	fmt.Println(time.Since(start))
}

func slowFunction(n int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(n)
}
