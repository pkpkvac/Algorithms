package reorganizestring

import "container/heap"

type Task struct {
	task byte
	freq int
}

// type TasksWithExpiry struct {
// 	task Task
// }

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

func reorganizeString(s string) string {

	h := &MaxHeap{}
	heap.Init(h)
	q := []Task{}

	m := make(map[byte]int)

	for _, t := range s {
		m[byte(t)]++
	}

	for task, count := range m {
		hItem := Task{task: task, freq: count}
		heap.Push(h, hItem)
	}

	// time := 0
	// start processing items
	retStr := []byte{}
	for h.Len() > 0 || len(q) > 0 {

		// Move all ready tasks from queue back to heap
		for len(q) > 0 {
			task := q[0]
			q = q[1:]
			heap.Push(h, task)
		}

		if h.Len() > 0 {
			// Process the most frequent task

			task1 := heap.Pop(h).(Task)
			task1.freq--

			if h.Len() == 0 && task1.freq > 0 {
				return ""
			}

			retStr = append(retStr, task1.task)

			if task1.freq > 0 {
				q = append(q, Task{task: task1.task, freq: task1.freq})
			}

			if h.Len() > 0 {
				task2 := heap.Pop(h).(Task)
				task2.freq--
				retStr = append(retStr, task2.task)
				if task2.freq > 0 {
					q = append(q, Task{task: task2.task, freq: task2.freq})
				}
			}

		}
		// Always increment time (whether processing task or idling)
	}

	return string(retStr)
}

func reorganizeStringClean(s string) string {
	h := &MaxHeap{}
	heap.Init(h)

	// Count frequencies
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}

	// Build heap
	for char, count := range m {
		heap.Push(h, Task{task: char, freq: count})
	}

	var result []byte

	for h.Len() > 0 {
		// Pop most frequent
		first := heap.Pop(h).(Task)
		result = append(result, first.task)
		first.freq--

		// If only one char left and it has count > 0, impossible
		if h.Len() == 0 && first.freq > 0 {
			return ""
		}

		// Pop second most frequent if available
		if h.Len() > 0 {
			second := heap.Pop(h).(Task)
			result = append(result, second.task)
			second.freq--

			// Put second back if count > 0
			if second.freq > 0 {
				heap.Push(h, second)
			}
		}

		// Put first back if count > 0
		if first.freq > 0 {
			heap.Push(h, first)
		}
	}

	return string(result)
}
