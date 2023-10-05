package queue

import (
	"time"

	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/runner"
)

type JobQueue struct {
	queue chan Allocate
}

type Allocate struct {
	Who runner.Runner
	When time.Time
	Job job.Job
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{
		queue: make(chan Allocate, size),
	}
}

func (q *JobQueue) Enqueue(item Allocate) {
	q.queue <- item
}

func (q *JobQueue) Dequeue() (Allocate, bool) {
	select {
	case item,ok := <-q.queue:
		return item, ok
	default:
		return Allocate{}, false
	}
}

func (q *JobQueue) GetNotAllocate() (job.Job, bool) {
	select {
	case item, ok := <-q.queue:
		if ok {
			if (len(item.Who.Who) >= 1) {
				q.Enqueue(item)
			} else {
				return item.Job, false
			}
		}

		return job.Job{}, false
	default:
		return job.Job{}, false
	}
}

func (q *JobQueue) SetAllocate(allocate Allocate) {
	q.Enqueue(allocate)
}

func (q *JobQueue) CheckTimeOut() {
	select {
	case item, ok := <-q.queue:
		if ok {
			if len(item.Who.Who) >= 1 {
				current := time.Now()

				if item.When.Sub(current) > time.Hour {
					newItem := &Allocate{
						Job: item.Job,
					}

					q.Enqueue(*newItem)
				} else {
					q.Enqueue(item)
				}
			}
		}
	}
}

func (q *JobQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *JobQueue) Size() int {
	return len(q.queue)
}