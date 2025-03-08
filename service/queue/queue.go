package queue

import "github.com/truongtu268/distributePriorityQueue/model"

type Queue interface {
	Enqueue(model.PriorityQueueTask)
	Dequeue() (model.PriorityQueueTask, bool)
	Peek() (model.PriorityQueueTask, bool)
	IsClearQueue() bool
}
