package workerpool

import "log"

type Unit interface {
	Job() error
}

// cola de trabajo compartida
var jobQueue chan Unit

func InitQueue(maxQueue int) {
	jobQueue = make(chan Unit, maxQueue)
}

type worker struct {
	//chanel of chanels of Units
	pool  chan chan Unit
	jobCh chan Unit
}

// Pool de workers
type dispatcher struct {
	pool    chan chan Unit
	workers int
}

// Constructor of worker
func newWorker(pool chan chan Unit) *worker {
	return &worker{
		jobCh: make(chan Unit),
		pool:  pool,
	}
}

func (w *worker) start() {
	go func() {
		for {
			// register the actual worker in the queue.
			w.pool <- w.jobCh
			select {
			case job := <-w.jobCh:
				// do the actual job here
				err := job.Job()

				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	}()
}

//Constructor of dispatcher (pool of workers)
func NewDispatcher(maxWorkers int) *dispatcher {
	return &dispatcher{
		pool:    make(chan chan Unit, maxWorkers),
		workers: maxWorkers,
	}
}

// Run is the starting point. This should be called by the client.
func (d *dispatcher) Run(async bool) {
	for i := 0; i < d.workers; i++ {
		w := newWorker(d.pool)
		w.start()
	}

	if async {
		go d.dispatchAsync()
	} else {
		go d.dispatch()
	}
}

// dispatch async
func (d *dispatcher) dispatchAsync() {
	for job := range jobQueue {
		go func(j Unit) {
			jobChannel := <-d.pool
			jobChannel <- j
		}(job)
	}
}

// dispatch not async
func (d *dispatcher) dispatch() {
	// goroutine responsable of assingn workers availables at income tasks
	go func() {
		for {
			select {
			case job, ok := <-jobQueue: // tomar tareas de una cola de tareas
				if ok {
					jobChannel := <-d.pool //asignarlas a un trabajador disponible del grupo
					jobChannel <- job
				}
			}
		}
	}()
}

func AddUnit(u Unit) {
	jobQueue <- u
}
