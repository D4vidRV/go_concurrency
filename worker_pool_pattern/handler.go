package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/D4vidRV/worker_pool_pattern/workerpool"
)

type HeavyWork struct {
	Name   string `json:"name"`
	number int
}

func (p *HeavyWork) Job() error {
	// time.Sleep(500 * time.Millisecond)
	fmt.Printf("heavy job is running %d \n", p.number)
	return nil
}

func JobHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getting job...")
	q := r.URL.Query().Get("q")

	if q == "" {
		q = "default"
	}

	for i := 0; i < 1000; i++ {
		work := HeavyWork{Name: q, number: i}
		workerpool.AddUnit(&work)
	}
}
