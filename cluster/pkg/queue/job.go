package queue

import (
	"time"

	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/runner"
	"go.uber.org/zap"
)

type PriorityQueue struct {
	queue []*Item
}

type Item struct {
	Value Allocate
	Index int
}

type Allocate struct {
	Who runner.Runner
	When time.Time
	Job *job.Job
}

func NewPriorityQueue() *PriorityQueue {
	h := &PriorityQueue{}

	return h
}

func (pq *PriorityQueue) Len() int {
	return len(pq.queue)
}

func (pq *PriorityQueue) Less(i,j int) bool {
	if pq.queue[i].Value.When.IsZero(){
		return true
	}

	if pq.queue[j].Value.When.IsZero(){
		return false
	}

	return pq.queue[i].Value.When.Before(pq.queue[j].Value.When)
}

func (pq *PriorityQueue) Push(item *Item) {
	n := len(pq.queue)
	item.Index = n
	pq.queue = append(pq.queue, item)
	pq.bubbleUp(len(pq.queue)-1)
}
 
func (pq *PriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
	pq.queue[i].Index = i
	pq.queue[j].Index = j
}

func (pq *PriorityQueue) Remove(item *job.Job) (bool) {
	
	tmp := make([]*Item, len(pq.queue))
	copy(tmp, pq.queue)

	Index := -1
	found := false

	for idx, val := range tmp {
		if val.Value.Job.TextId == item.TextId {
			Index = idx 
			found = true
			continue
		}

		if found {
			tmp[idx].Index--
		}
	}

	if !found {
		logger.Debug("Can't find item")
		return false
	}

	if len(tmp) == 1 {
		pq.queue = make([]*Item, 0)
		return true
	} 
	
	if Index == len(tmp)-1 {
		pq.queue = tmp[:Index]
		return true
	}
		
	pq.queue = append(tmp[:Index], tmp[Index+1])
	return true
}
 
func (pq *PriorityQueue) Pop() (*Item, bool) {
	if pq.Len() == 0 {
        return nil, false
    }

	if pq.Len() == 1 {
		item := pq.queue[0]
		pq.queue = []*Item{}
		return item, true
	}

    root := pq.queue[0]
	last := pq.queue[pq.Len()-1]

	pq.queue[0] = last
	pq.queue = pq.queue[:len(pq.queue)-1]
	pq.bubbleDown(0)

	return root, true
}

func (pq *PriorityQueue) SetAllocate(allocate *Allocate) {
	logger.Info("Set allocate", zap.Any("Allocate", allocate.Who.Who))
	current := time.Now()
	found := false
	for idx, val := range pq.queue {
		if val.Value.Job == allocate.Job {
			pq.queue[idx].Value.Who = allocate.Who
			pq.queue[idx].Value.When = current
			found = true
			break
		}
	}

	if !found {
		logger.Error("Can't allocate", zap.Any("allocate", allocate))
	}
}

func (pq *PriorityQueue) GetNotAllocate() (*job.Job, bool) {
	for _, item := range pq.queue {
		if item.Value.Who == (runner.Runner{}) && len(item.Value.Who.Who) == 0 {
			return item.Value.Job, true
		}
	}

	return nil, false
}

func (pq *PriorityQueue) CheckTimeOut(t int) []*Item {
    currentTime := time.Now()
    var timedOutItems []*Item
    var nonTimedOutItems []*Item

    for pq.Len() > 0 {
        tmp, ok := pq.Pop()

		if ok {
			if tmp.Value.When.IsZero() {
                logger.Error("time is weird", zap.Any("tmp", tmp))
            } else {
                duration := currentTime.Sub(tmp.Value.When)
                if duration >= time.Hour * time.Duration(t) {
                    logger.Info("Timeout!", zap.Any("content", tmp.Value.Job.Content), zap.Any("by", tmp.Value.Who))
                    timedOutItems = append(timedOutItems, tmp)
                } else {
                    nonTimedOutItems = append(nonTimedOutItems, tmp)
                }
            }
		}
    }

    for _, item := range nonTimedOutItems {
        pq.Push(item)
    }

    return timedOutItems
}

func (pq *PriorityQueue) bubbleUp(index int) {
	for index > 0 {
		parentIdx := (index - 1) /2
		
		if pq.queue[index].Value.When.After(pq.queue[parentIdx].Value.When) {
			break
		}

		pq.queue[index], pq.queue[parentIdx] = pq.queue[parentIdx], pq.queue[index]
		index = parentIdx
	}
}

func (pq *PriorityQueue) bubbleDown(index int) {
	currentTime := time.Now()

	for {
		leftIdx := index*2 + 1
		rightIdx := index*2 + 2

		min := index

		if leftIdx < pq.Len() {
			durationLeftIdx := currentTime.Sub(pq.queue[leftIdx].Value.When)
			durationMinIdx := currentTime.Sub(pq.queue[min].Value.When)
			
			if durationMinIdx < durationLeftIdx {
				min = leftIdx
			}
		}

		if rightIdx < pq.Len() {
			durationRightIdx := currentTime.Sub(pq.queue[rightIdx].Value.When)
			durationMinIdx := currentTime.Sub(pq.queue[min].Value.When)
			
			if durationMinIdx < durationRightIdx {
				min = rightIdx
			}
		}

		if min == index {
			break
		}

		pq.queue[index], pq.queue[min] = pq.queue[min], pq.queue[index]
		index = min
	}
}