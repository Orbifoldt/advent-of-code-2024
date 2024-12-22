package util

type Entry struct {
	V     Vec
	Dist  int
	index int
}

func NewEntry(v Vec, dist int) *Entry {
	return &Entry{v, dist, 0}
}

type PriorityQueueVec []*Entry // See https://pkg.go.dev/container/heap#example-package-PriorityQueue

func (pq PriorityQueueVec) Len() int { return len(pq) }
func (pq PriorityQueueVec) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}

func (pq PriorityQueueVec) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueVec) Push(x any) {
	n := len(*pq)
	item := x.(*Entry)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueVec) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
