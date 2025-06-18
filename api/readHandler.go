package api

import (
	"cache"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get(IdParam)
	if taskId == "" {
		w.Write([]byte("please provide an id"))
		return
	}
	task, err := cache.GlobalCacheManager.Read(taskId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	marshaledBody, err := json.Marshal(task)
	if err != nil {
		marshErr := fmt.Errorf("encountered error while unmarshaling: %w", err)
		fmt.Println(marshErr)
		w.Write([]byte(marshErr.Error()))
		return
	}
	w.Write(marshaledBody)
}
