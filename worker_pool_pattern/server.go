package main

import (
	"log"
	"net/http"

	"github.com/D4vidRV/worker_pool_pattern/workerpool"
)

func main() {
	// Init the queue
	workerpool.InitQueue(20)

	// Init the dispacher and keep it listening.
	workerpool.NewDispatcher(15).Run(true)

	// Init a server
	s := http.Server{}
	s.Addr = ":8080"
	addJobRoute()
	log.Fatal(s.ListenAndServe())
}

func addJobRoute() {
	http.HandleFunc("/jobs", JobHandler)
}
