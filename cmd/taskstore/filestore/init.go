package filestore

import (
	"sync"
)

var (
	once     sync.Once
	instance *TaskStore
)

func GetTaskFileStoreInstance() *TaskStore {
	once.Do(func() {
		instance = NewTaskStore()
	})
	
	return instance
}
