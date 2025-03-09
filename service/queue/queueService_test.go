package queue

import (
	"testing"
	"time"

	"github.com/truongtu268/distributePriorityQueue/model"
)

func TestEnqueueAndDequeueService(t *testing.T) {
	service := NewService(1 * time.Second)

	task1 := model.PriorityQueueTask{Priority: 1, ItemID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, ItemID: "Task 2"}

	service.Enqueue(task1)
	service.Enqueue(task2)

	dequeuedTask, ok := service.Dequeue()
	if !ok || dequeuedTask.ItemID != "Task 2" {
		t.Errorf("Expected Task 2, got %v", dequeuedTask.ItemID)
	}

	dequeuedTask, ok = service.Dequeue()
	if !ok || dequeuedTask.ItemID != "Task 1" {
		t.Errorf("Expected Task 1, got %v", dequeuedTask.ItemID)
	}
}

func TestPeekService(t *testing.T) {
	service := NewService(1 * time.Second)

	task := model.PriorityQueueTask{Priority: 1, ItemID: "Task"}

	service.Enqueue(task)

	peekedTask, ok := service.Peek()
	if !ok || peekedTask.ItemID != "Task" {
		t.Errorf("Expected Task, got %v", peekedTask.ItemID)
	}

	// Ensure the task is still in the queue after peeking
	dequeuedTask, ok := service.Dequeue()
	if !ok || dequeuedTask.ItemID != "Task" {
		t.Errorf("Expected Task, got %v", dequeuedTask.ItemID)
	}
}

func TestIsClearQueueService(t *testing.T) {
	service := NewService(1 * time.Second)

	if !service.IsClearQueue() {
		t.Error("Expected queue to be clear")
	}

	task := model.PriorityQueueTask{Priority: 1, ItemID: "Task"}
	service.Enqueue(task)

	if service.IsClearQueue() {
		t.Error("Expected queue not to be clear")
	}
}

func TestMergeQueues(t *testing.T) {
	service := NewService(1 * time.Second)

	task1 := model.PriorityQueueTask{Priority: 1, ItemID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, ItemID: "Task 2"}

	service.Enqueue(task1)
	service.Enqueue(task2)

	// Simulate merging
	service.mergeQueues()

	// After merging, OldAgeQueue should have the tasks
	dequeuedTask, ok := service.Dequeue()
	if !ok || dequeuedTask.ItemID != "Task 2" {
		t.Errorf("Expected Task 2, got %v", dequeuedTask.ItemID)
	}

	dequeuedTask, ok = service.Dequeue()
	if !ok || dequeuedTask.ItemID != "Task 1" {
		t.Errorf("Expected Task 1, got %v", dequeuedTask.ItemID)
	}
}

func TestStartMerging(t *testing.T) {
	// Use a short duration for testing
	service := NewService(100 * time.Millisecond)

	task1 := model.PriorityQueueTask{Priority: 1, ItemID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, ItemID: "Task 2"}

	service.Enqueue(task1)
	service.Enqueue(task2)

	// Run StartMerging in a separate goroutine
	go service.StartMerging()

	// Wait for a bit longer than the merge duration to ensure merging happens
	time.Sleep(200 * time.Millisecond)

	task3 := model.PriorityQueueTask{Priority: 3, ItemID: "Task 3"}
	service.Enqueue(task3)
	// After merging, OldAgeQueue should have the tasks
	dequeuedTask, ok := service.Dequeue()
	if !ok || dequeuedTask.ItemID != "Task 2" {
		t.Errorf("Expected Task 2, got %v", dequeuedTask.ItemID)
	}

	dequeuedTask, ok = service.Dequeue()
	if !ok || dequeuedTask.ItemID != "Task 1" {
		t.Errorf("Expected Task 1, got %v", dequeuedTask.ItemID)
	}
}
