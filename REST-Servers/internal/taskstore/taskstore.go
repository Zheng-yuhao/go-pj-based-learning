package taskstore

import "time"

/*
db model

- Id
- Text
- Tags
- Due
*/

type Task struct {
	Id   int       `json:"int"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type TaskStore struct {
}

func New() *TaskStore

func (ts *TaskStore) CreateTask(test string, tags []string, due time.Time) int

func (ts *TaskStore) GetTask(id int) (Task, error)

func (ts *TaskStore) DeleteTask(id int) error

func (ts *TaskStore) DeleteAllTasks() error

func (ts *TaskStore) GetAllTasks() []Task

func (ts *TaskStore) GetTaskByTag(tag string) []Task

func (ts *TaskStore) GetTasksByDueDate(due time.Time) []Task
