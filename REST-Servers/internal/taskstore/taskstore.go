package taskstore

import (
	"fmt"
	"sync"
	"time"
)

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
	sync.Mutex

	tasks  map[int]Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:   ts.nextId,
		Text: text,
		Due:  due,
	}
	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	// ts.tasksはmap型で保存されている
	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}

func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	// get task by id
	task, ok := ts.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("task with id = %d not found", id)
	}

	return task, nil
}

func (ts *TaskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	_, ok := ts.tasks[id]
	if !ok {
		return fmt.Errorf("task with id = %d not found", id)
	}

	delete(ts.tasks, id)
	return nil
}

func (ts *TaskStore) DeleteAllTasks() error {
	ts.Lock()
	defer ts.Unlock()

	// reassigne the map struct == delete all
	ts.tasks = make(map[int]Task)
	return nil
}

func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	tasks := []Task{}

	for _, val := range ts.tasks {
		tasks = append(tasks, val)
	}
	return tasks
}

func (ts *TaskStore) GetTaskByTag(tag string) []Task {
	ts.Lock()
	defer ts.Unlock()

	tasks := []Task{}

	for _, val := range ts.tasks {
		for _, tag_in_struct := range val.Tags {
			if tag == tag_in_struct {
				tasks = append(tasks, val)
				break
			}
		}
	}
	return tasks
}

func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
	ts.Lock()
	defer ts.Unlock()

	tasks := []Task{}

	for _, task := range ts.tasks {
		y, m, d := task.Due.Date()
		if y == year && m == month && d == day {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
