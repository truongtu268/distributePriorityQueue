package queue

import (
	"fmt"
	"time"

	"github.com/truongtu268/distributePriorityQueue/model"
)

type Service struct {
	OldAgeQueue   *PriorityQueue
	NewAgeQueue   *PriorityQueue
	RetryQueue    *PriorityQueue
	mergeDuration time.Duration
}

func NewService(switchDuration time.Duration) *Service {
	return &Service{
		OldAgeQueue:   NewPriorityQueue(),
		NewAgeQueue:   NewPriorityQueue(),
		RetryQueue:    NewPriorityQueue(),
		mergeDuration: switchDuration,
	}
}

func (s *Service) Enqueue(task model.PriorityQueueTask) {
	s.NewAgeQueue.Enqueue(task)
}

func (s *Service) Dequeue() (model.PriorityQueueTask, bool) {
	if !s.RetryQueue.IsClearQueue() {
		task, isValid := s.RetryQueue.Dequeue()
		if isValid {
			return task, true
		}
	}

	if !s.OldAgeQueue.IsClearQueue() {
		task, isValid := s.OldAgeQueue.Dequeue()
		if isValid {
			return task, true
		}
	}

	return s.NewAgeQueue.Dequeue()
}

func (s *Service) Peek() (model.PriorityQueueTask, bool) {
	if !s.RetryQueue.IsClearQueue() {
		task, isValid := s.RetryQueue.Peek()
		if isValid {
			return task, true
		}
	}

	if !s.OldAgeQueue.IsClearQueue() {
		task, isValid := s.OldAgeQueue.Peek()
		if isValid {
			return task, true
		}
	}

	return s.NewAgeQueue.Peek()
}

func (s *Service) IsClearQueue() bool {
	return s.OldAgeQueue.IsClearQueue() &&
		s.NewAgeQueue.IsClearQueue() &&
		s.RetryQueue.IsClearQueue()
}

// StartMerging starts the process of merging between the queues
func (s *Service) StartMerging() {
	ticker := time.NewTicker(s.mergeDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.mergeQueues()
		}
	}
}

// mergeQueues switches the contents of OldAgeQueue and NewAgeQueue
func (s *Service) mergeQueues() {
	fmt.Println("Merging queues...")
	s.OldAgeQueue.MergeQueue(s.NewAgeQueue)
	s.NewAgeQueue = NewPriorityQueue()
	fmt.Println("Queues Merged.")
}
