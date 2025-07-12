package scheduler
import (
	"backend/models"
)
type MessageHeap []models.Message
func (h MessageHeap) Len() int { return len(h) }
func (h MessageHeap) Less(i, j int) bool {
	return h[i].Time.Before(h[j].Time) // min-heap by send time
}
func (h MessageHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MessageHeap) Push(x any) {
	*h = append(*h, x.(models.Message))
}
func (h *MessageHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0:n-1]
	return item
}

var H = &MessageHeap{}
