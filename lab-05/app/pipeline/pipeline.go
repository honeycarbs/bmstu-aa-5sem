package pipeline

import (
	"time"

	"../bmsearch"
)



func Pipeline(count int, ch chan int) *Queue {
	first := make(chan *PipeTask, count)
	second := make(chan *PipeTask, count)
	third := make(chan *PipeTask, count)

	line := initQueue(count)

	
	// funcs
	gen_string := func() {
		for {
			select {
			case pipe_task := <-first:
				pipe_task.generated = true

				pipe_task.start_generating = time.Now()

				pipe_task.source = bmsearch.GenerateRune(STRING_SIZE)
				pipe_task.pattern = bmsearch.GeneratePattern(pipe_task.source, PATTERN_SIZE)

				pipe_task.end_generatig = time.Now()

				second <- pipe_task
			}
		}
	}

	get_table := func() {
		for {
			select {
			case pipe_task := <-second:
				pipe_task.skip_table_made = true

				pipe_task.start_table = time.Now()

				pipe_task.skip_table = bmsearch.ConstructSkipTable(pipe_task.pattern)
				pipe_task.end_table = time.Now()

				third <- pipe_task
			}
		}
	}

	match := func() {
		for {
			select {
			case pipe_task := <-third:
				pipe_task.pattern_mached = true

				pipe_task.start_match = time.Now()
				pipe_task.pattern_index = bmsearch.FindFirstIndex(pipe_task.source, pipe_task.pattern, pipe_task.skip_table)
				pipe_task.end_match = time.Now()

				line.enqueue(pipe_task)
				if (pipe_task.num == count - 1) {
					ch <- 0
				}
			}
		}
	}
	go gen_string()
	go get_table()
	go match()


	for i := 0; i < count; i++ {
		pipe_task := new(PipeTask)
		pipe_task.num = i + 1
		first <- pipe_task
	}
	return line
}
