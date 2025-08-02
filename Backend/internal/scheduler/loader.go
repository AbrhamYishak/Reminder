package scheduler

import (
	"backend/db"
	"container/heap"
	"backend/models"
	"log"
)
func Loader(){
	db := db.Connection()
	var messages []models.Message
	if err := db.Order("Time asc").Limit(100).Find(&messages).Error; err != nil{
		log.Fatal("could not load the data")	
	}
	heap.Init(H)
	for _, m := range messages {
		heap.Push(H, m)
	}
}
