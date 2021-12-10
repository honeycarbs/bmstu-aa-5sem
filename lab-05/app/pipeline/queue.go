package pipeline

func initQueue(size int) *Queue {
	qu := new(Queue)
	qu.queue = make([](*PipeTask), size)
	qu.size = -1

	return qu
}

func (qu *Queue) enqueue (t *PipeTask) {
	if (qu.size != len(qu.queue) - 1) {
		qu.queue[qu.size + 1] = t
		qu.size++
	}
}

func (qu *Queue) dequeue () *PipeTask {
	t := qu.queue[0]
	qu.queue = qu.queue[1:]
	qu.size--
	return t
}