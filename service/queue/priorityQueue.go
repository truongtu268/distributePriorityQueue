package queue

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/truongtu268/distributePriorityQueue/model"
)

// PriorityQueue represents a thread-safe priority queue
type PriorityQueue struct {
	items   sync.Map
	isClean atomic.Bool
}

// NewPriorityQueue creates a new PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	isClean := atomic.Bool{}
	isClean.Store(true)
	return &PriorityQueue{
		isClean: isClean,
	}
}

// Enqueue adds an item with a given priority to the queue
func (pq *PriorityQueue) Enqueue(item model.PriorityQueueTask) {
	item.CreatedAt = time.Now() // Set the timestamp when enqueuing
	var priorityQueue []model.PriorityQueueTask
	queue, ok := pq.items.Load(item.Priority)
	if !ok {
		priorityQueue = []model.PriorityQueueTask{item}
	} else {
		priorityQueue = queue.([]model.PriorityQueueTask)
		priorityQueue = append(priorityQueue, item)
	}
	pq.items.Store(item.Priority, priorityQueue)
	if pq.isClean.Load() {
		pq.isClean.Store(false)
	}
}

// Dequeue removes and returns the item with the highest priority
func (pq *PriorityQueue) Dequeue() (model.PriorityQueueTask, bool) {
	//pq.AgeTasks() // Adjust priorities based on age before dequeuing

	var highestPriority int
	var highestItem model.PriorityQueueTask
	found := false
	countQueue := 0

	pq.items.Range(func(key, value interface{}) bool {
		priority := key.(int)
		countQueue++
		if !found || priority > highestPriority {
			highestPriority = priority
			highestItem = value.([]model.PriorityQueueTask)[0]
			found = true
		}
		return true
	})

	if found {
		queue, _ := pq.items.Load(highestPriority)
		priorityQueue := queue.([]model.PriorityQueueTask)
		if len(priorityQueue) > 1 {
			pq.items.Store(highestPriority, priorityQueue[1:])
		} else {
			pq.items.Delete(highestPriority)
			if countQueue == 1 {
				pq.isClean.Store(true)
			}
		}
	} else {
		pq.isClean.Store(true)
	}
	return highestItem, found
}

// Peek returns the item with the highest priority without removing it
func (pq *PriorityQueue) Peek() (model.PriorityQueueTask, bool) {
	var highestPriority int
	var highestItem model.PriorityQueueTask
	found := false

	pq.items.Range(func(key, value interface{}) bool {
		priority := key.(int)
		if !found || priority > highestPriority {
			highestPriority = priority
			highestItem = value.([]model.PriorityQueueTask)[0]
			found = true
		}
		return true
	})

	return highestItem, found
}

// PrintQueue prints all items in the queue for debugging
func (pq *PriorityQueue) PrintQueue() {
	pq.items.Range(func(key, value interface{}) bool {
		fmt.Printf("Priority: %d, Items: %v\n", key, value)
		return true
	})
}

// IsClearQueue returns state for priority queue
func (pq *PriorityQueue) IsClearQueue() bool {
	return pq.isClean.Load()
}

// MergeQueue merge main queue with another queue
func (pq *PriorityQueue) MergeQueue(q *PriorityQueue) {
	// queue with same priority of main queue and delete this key
	pq.items.Range(func(key, value interface{}) bool {
		priority := key.(int)
		mainQueue := value.([]model.PriorityQueueTask)
		val, ok := q.items.Load(priority)
		if ok {
			newQueue := val.([]model.PriorityQueueTask)
			mainQueue = append(mainQueue, newQueue...)
		}
		q.items.Delete(priority)
		pq.isClean.Store(false)
		return true
	})

	// copy value from priority not in main queue
	q.items.Range(func(key, value interface{}) bool {
		pq.items.Store(key, value)
		pq.isClean.Store(false)
		return true
	})
}
