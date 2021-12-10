package pipeline

import (
	"time"

	"../bmsearch"
)

func gen_string_sync(task *PipeTask) *PipeTask {
	task.generated = true

	task.start_generating = time.Now()

	task.source = bmsearch.GenerateRune(STRING_SIZE)
	task.pattern = bmsearch.GeneratePattern(task.source, PATTERN_SIZE)

	task.end_generatig = time.Now()

	return task
}

func get_table_sync(task *PipeTask) *PipeTask {
	task.skip_table_made = true

	task.start_table = time.Now()

	task.skip_table = bmsearch.ConstructSkipTable(task.pattern)
	task.end_table = time.Now()

	return task
}

func match_sync(task *PipeTask) *PipeTask {
	task.pattern_mached = true

	task.start_match = time.Now()
	task.pattern_index = bmsearch.FindFirstIndex(task.source, task.pattern, task.skip_table)
	task.end_match = time.Now()

	return task
}


func Sync(count int) *Queue {
	line_first := initQueue(count)
	line_second := initQueue(count)
	line_third := initQueue(count)

	for i := 0; i < count; i++ {
		pipe_task := new(PipeTask)
		pipe_task = gen_string_sync(pipe_task)
		line_first.enqueue(pipe_task)
		if (len(line_first.queue) != 0) {
			pipe_task = get_table_sync(line_first.dequeue())
			line_second.enqueue(pipe_task)
			if (len(line_second.queue) != 0) {
				pipe_task = match_sync(line_second.dequeue())
				line_third.enqueue(pipe_task)
			}
		}
	}
	return line_third
}