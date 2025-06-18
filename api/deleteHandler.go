package api

import (
	"cache"
	"net/http"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get(IdParam)
	if taskId == "" {
		w.Write([]byte("please provide an id"))
		return
	}
	cache.GlobalCacheManager.Delete(taskId)
	w.Write([]byte("task deleted"))
}
