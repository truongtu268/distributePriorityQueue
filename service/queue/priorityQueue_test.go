package queue

import (
	"testing"

	"github.com/truongtu268/distributePriorityQueue/model"
)

func TestEnqueueAndDequeue(t *testing.T) {
	pq := NewPriorityQueue()

	task1 := model.PriorityQueueTask{Priority: 1, ItemID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, ItemID: "Task 2"}

	pq.Enqueue(task1)
	pq.Enqueue(task2)

	dequeuedTask, found := pq.Dequeue()
	if !found {
		t.Fatal("Expected to find a task, but none was found")
	}

	if dequeuedTask.ItemID != "Task 2" {
		t.Errorf("Expected 'Task 2', got '%s'", dequeuedTask.ItemID)
	}

	dequeuedTask, found = pq.Dequeue()
	if !found {
		t.Fatal("Expected to find a task, but none was found")
	}

	if dequeuedTask.ItemID != "Task 1" {
		t.Errorf("Expected 'Task 1', got '%s'", dequeuedTask.ItemID)
	}
}

func TestPeek(t *testing.T) {
	pq := NewPriorityQueue()

	task1 := model.PriorityQueueTask{Priority: 1, ItemID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, ItemID: "Task 2"}

	pq.Enqueue(task1)
	pq.Enqueue(task2)

	peekedTask, found := pq.Peek()
	if !found {
		t.Fatal("Expected to find a task, but none was found")
	}

	if peekedTask.ItemID != "Task 2" {
		t.Errorf("Expected 'Task 2', got '%s'", peekedTask.ItemID)
	}

	// Ensure the task is not removed
	peekedTask, found = pq.Peek()
	if !found {
		t.Fatal("Expected to find a task, but none was found")
	}

	if peekedTask.ItemID != "Task 2" {
		t.Errorf("Expected 'Task 2', got '%s'", peekedTask.ItemID)
	}
}

func TestIsClearQueue(t *testing.T) {
	pq := NewPriorityQueue()

	if !pq.IsClearQueue() {
		t.Error("Expected queue to be clear initially")
	}

	task := model.PriorityQueueTask{Priority: 1, ItemID: "Task"}
	pq.Enqueue(task)

	if pq.IsClearQueue() {
		t.Error("Expected queue not to be clear after enqueue")
	}

	pq.Dequeue()

	if !pq.IsClearQueue() {
		t.Error("Expected queue to be clear after dequeueing all items")
	}
}

func TestMergeQueue(t *testing.T) {
	pq1 := NewPriorityQueue()
	pq2 := NewPriorityQueue()

	task1 := model.PriorityQueueTask{Priority: 1, ItemID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, ItemID: "Task 2"}
	task3 := model.PriorityQueueTask{Priority: 1, ItemID: "Task 3"}

	pq1.Enqueue(task1)
	pq2.Enqueue(task2)
	pq2.Enqueue(task3)

	pq1.MergeQueue(pq2)

	dequeuedTask, found := pq1.Dequeue()
	if !found || dequeuedTask.ItemID != "Task 2" {
		t.Errorf("Expected 'Task 2', got '%s'", dequeuedTask.ItemID)
	}

	dequeuedTask, found = pq1.Dequeue()
	if !found || dequeuedTask.ItemID != "Task 1" {
		t.Errorf("Expected 'Task 1', got '%s'", dequeuedTask.ItemID)
	}

	dequeuedTask, found = pq1.Dequeue()
	if !found || dequeuedTask.ItemID != "Task 3" {
		t.Errorf("Expected 'Task 3', got '%s'", dequeuedTask.ItemID)
	}
}
