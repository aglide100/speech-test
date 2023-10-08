package queue

import (
	"sync"
	"time"

	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/runner"
	"go.uber.org/zap"
)

type PriorityJobQueue struct {
	mutex sync.RWMutex
	queue []*Allocate
	length int
}

type Allocate struct {
	Who runner.Runner
	When time.Time
	Job job.Job
	Index int
}

func NewJobQueue(size int) *PriorityJobQueue {
	return &PriorityJobQueue{
		queue: []*Allocate{},
		length: size,
	}
}

func (pq *PriorityJobQueue) Length() int {
	return pq.length
}


func (pq *PriorityJobQueue) Size() int {
	return len(pq.queue)
}

func (pq *PriorityJobQueue) Less(i,j int) bool {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()
	if pq.queue[i].When.IsZero(){
		return true
	}

	if pq.queue[j].When.IsZero(){
		return false
	}

	pq.mutex.RLocker()
	return pq.queue[i].When.Before(pq.queue[j].When)
}

func (pq *PriorityJobQueue) Push(item *Allocate) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	n := len(pq.queue)
	item.Index = n
	pq.queue = append(pq.queue, item)
	logger.Info("Size", zap.Any("size", len(pq.queue)))
}
 
func (pq *PriorityJobQueue) Swap(i, j int) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
	pq.queue[i].Index = i
	pq.queue[j].Index = j
}

func (pq *PriorityJobQueue) Remove(item *Allocate) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	old := []*Allocate{}
	copy(old,pq.queue)

	index := -1
	found := false

	for idx, val := range old {
		if val == item {
			index = idx
			found = true
			continue
		}

		if found {
			old[idx].Index--
		}
	}

	if !found {
		logger.Error("Can't find item")
		return
	}

	old = append(old[:index], old[index+1])

	pq.queue = old
}
 
func (pq *PriorityJobQueue) Pop() *Allocate {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	old := pq
	n := len(old.queue)
	item := old.queue[n-1]
	old.queue[n-1] = nil
	pq.queue = old.queue[0:n-1]
	
	return item
}

func (pq *PriorityJobQueue) SetAllocate(item *Allocate) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	logger.Info("Set allocate", zap.Any("Allocate", item.Who.Who))
	current := time.Now()
	found := false
	for idx, val := range pq.queue {
		if val.Job == item.Job {
			pq.queue[idx].Who = item.Who
			pq.queue[idx].When = current
			found = true
			break
		}
	}

	if !found {
		logger.Error("Can't allocate", zap.Any("item", item))
	}
}

func (pq *PriorityJobQueue) GetNotAllocate() (job.Job, bool) {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	for _, item := range pq.queue {
		if item.Who == (runner.Runner{}) && len(item.Who.Who) == 0 {
			return item.Job, true
		}
	}

	return job.Job{}, false
}

func (pq *PriorityJobQueue) CheckTimeOut() {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	currentTime := time.Now()
	
	for idx, item := range pq.queue {
		if item.When.IsZero() {
			continue
		}
		duration := currentTime.Sub(pq.queue[idx].When)


		if duration >= time.Hour * 2 {
			logger.Info("Timeout!", zap.Any("item", pq.queue[idx]))

			pq.queue[idx].Who = runner.Runner{}
			pq.queue[idx].When = time.Time{}
		}
	}
}
