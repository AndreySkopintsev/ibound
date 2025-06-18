package cache

import (
	"fmt"
	"sync"
	"time"
)

type enum string

const (
	TaskDone       enum = "task done"
	TaskInProgress enum = "task in progress"
)

var (
	GlobalCacheManager *CacheManager
	ErrNotInCache      = fmt.Errorf("requested task is not in cache")
)

type CacheManager struct {
	cache map[string]*CacheModel
	mu    sync.RWMutex
}

func InitTaskManager() {
	GlobalCacheManager = &CacheManager{
		cache: make(map[string]*CacheModel),
	}
}

type CacheModel struct {
	Id          string    `json:"id"`
	TaskStatus  enum      `json:"taskStatus"`
	CreatedAt   time.Time `json:"createdAt"`
	ElapsedTime string    `json:"elapsedTime,omitempty"`
	mu          sync.Mutex
}

func (c *CacheManager) Save(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newCacheModel := &CacheModel{
		Id:         id,
		TaskStatus: TaskInProgress,
		CreatedAt:  time.Now(),
	}

	c.cache[id] = newCacheModel

	go func(c *CacheModel) {
		time.Sleep(10 * time.Second)
		c.setStatus(TaskDone)
	}(newCacheModel)
}

func (t *CacheModel) setStatus(status enum) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.TaskStatus = status
}

func (c *CacheManager) Delete(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, id)
}

func (c *CacheManager) Read(id string) (*CacheModel, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if _, ok := c.cache[id]; !ok {
		return &CacheModel{}, ErrNotInCache
	}
	cachedTask := c.cache[id]
	elapsedDuration := time.Since(cachedTask.CreatedAt)
	cachedTask.ElapsedTime = fmt.Sprintf("%.f:%.f:%.f", elapsedDuration.Hours(), elapsedDuration.Minutes(), elapsedDuration.Seconds())
	return cachedTask, nil
}
