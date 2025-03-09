package model

type TaskStatus string

const (
	SubmittedStatus  TaskStatus = "submitted"
	QueuedStatus     TaskStatus = "queued"
	ProcessingStatus TaskStatus = "processing"
	CompletedStatus  TaskStatus = "completed"
)
