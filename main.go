package main

import (
	"api"
	"cache"
	"fmt"
	"net/http"
)

func main() {
	newRouter := api.NewRouter()
	cache.InitTaskManager()
	err := http.ListenAndServe(":8080", newRouter)
	if err != nil {
		fmt.Printf("encountered error while listeningon port 8080: %s", err.Error())
	}
}
