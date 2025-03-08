package queue

import (
	"testing"
	"time"

	"github.com/truongtu268/distributePriorityQueue/model"
)

func TestEnqueueAndDequeueService(t *testing.T) {
	service := NewService(1 * time.Second)

	task1 := model.PriorityQueueTask{Priority: 1, AdID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, AdID: "Task 2"}

	service.Enqueue(task1)
	service.Enqueue(task2)

	dequeuedTask, ok := service.Dequeue()
	if !ok || dequeuedTask.AdID != "Task 2" {
		t.Errorf("Expected Task 2, got %v", dequeuedTask.AdID)
	}

	dequeuedTask, ok = service.Dequeue()
	if !ok || dequeuedTask.AdID != "Task 1" {
		t.Errorf("Expected Task 1, got %v", dequeuedTask.AdID)
	}
}

func TestPeekService(t *testing.T) {
	service := NewService(1 * time.Second)

	task := model.PriorityQueueTask{Priority: 1, AdID: "Task"}

	service.Enqueue(task)

	peekedTask, ok := service.Peek()
	if !ok || peekedTask.AdID != "Task" {
		t.Errorf("Expected Task, got %v", peekedTask.AdID)
	}

	// Ensure the task is still in the queue after peeking
	dequeuedTask, ok := service.Dequeue()
	if !ok || dequeuedTask.AdID != "Task" {
		t.Errorf("Expected Task, got %v", dequeuedTask.AdID)
	}
}

func TestIsClearQueueService(t *testing.T) {
	service := NewService(1 * time.Second)

	if !service.IsClearQueue() {
		t.Error("Expected queue to be clear")
	}

	task := model.PriorityQueueTask{Priority: 1, AdID: "Task"}
	service.Enqueue(task)

	if service.IsClearQueue() {
		t.Error("Expected queue not to be clear")
	}
}

func TestMergeQueues(t *testing.T) {
	service := NewService(1 * time.Second)

	task1 := model.PriorityQueueTask{Priority: 1, AdID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, AdID: "Task 2"}

	service.Enqueue(task1)
	service.Enqueue(task2)

	// Simulate merging
	service.mergeQueues()

	// After merging, OldAgeQueue should have the tasks
	dequeuedTask, ok := service.Dequeue()
	if !ok || dequeuedTask.AdID != "Task 2" {
		t.Errorf("Expected Task 2, got %v", dequeuedTask.AdID)
	}

	dequeuedTask, ok = service.Dequeue()
	if !ok || dequeuedTask.AdID != "Task 1" {
		t.Errorf("Expected Task 1, got %v", dequeuedTask.AdID)
	}
}

func TestStartMerging(t *testing.T) {
	// Use a short duration for testing
	service := NewService(100 * time.Millisecond)

	task1 := model.PriorityQueueTask{Priority: 1, AdID: "Task 1"}
	task2 := model.PriorityQueueTask{Priority: 2, AdID: "Task 2"}

	service.Enqueue(task1)
	service.Enqueue(task2)

	// Run StartMerging in a separate goroutine
	go service.StartMerging()

	// Wait for a bit longer than the merge duration to ensure merging happens
	time.Sleep(200 * time.Millisecond)

	// After merging, OldAgeQueue should have the tasks
	dequeuedTask, ok := service.Dequeue()
	if !ok || dequeuedTask.AdID != "Task 2" {
		t.Errorf("Expected Task 2, got %v", dequeuedTask.AdID)
	}

	dequeuedTask, ok = service.Dequeue()
	if !ok || dequeuedTask.AdID != "Task 1" {
		t.Errorf("Expected Task 1, got %v", dequeuedTask.AdID)
	}
}
