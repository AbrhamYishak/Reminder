package scheduler

import (
	"backend/db"
	"container/heap"
	"backend/models"
)
func Loader(){
	db := db.Connection()
	var messages []models.Message
	db.Order("Time asc").Limit(100).Find(&messages)
	heap.Init(H)
	for _, m := range messages {
		heap.Push(H, m)
	}
}
