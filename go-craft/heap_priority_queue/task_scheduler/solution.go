package taskscheduler

import "container/heap"

type Task struct {
	task byte
	freq int
}

type TasksWithExpiry struct {
	task Task
	time int
}

type MaxHeap []Task

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MaxHeap) Peek() Task {
	return h[0]
}

func leastInterval(tasks []byte, n int) int {

	h := &MaxHeap{}

	heap.Init(h)
	q := []TasksWithExpiry{}

	m := make(map[byte]int)

	for _, t := range tasks {
		m[t]++
	}

	for task, count := range m {
		hItem := Task{task: task, freq: count}
		heap.Push(h, hItem)
	}

	time := 0
	// start processing items
	for h.Len() > 0 || len(q) > 0 {

		// Move all ready tasks from queue back to heap
		for len(q) > 0 && time >= q[0].time {
			task := q[0].task
			q = q[1:]
			heap.Push(h, task)
		}

		if h.Len() > 0 {
			// Process the most frequent task
			task := heap.Pop(h).(Task)
			task.freq--

			// Add back to queue if more instances remain
			if task.freq > 0 {
				q = append(q, TasksWithExpiry{task: task, time: time + n + 1})
			}
		}
		// Always increment time (whether processing task or idling)
		time++
	}

	return time

}
