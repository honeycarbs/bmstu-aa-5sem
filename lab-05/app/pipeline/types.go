package pipeline

import "time"

const STRING_SIZE = 10
const PATTERN_SIZE = 5


type PipeTask struct {
	// helpers
	num  			int
	generated		bool
	skip_table_made bool
	pattern_mached	bool


	// time labels
	start_generating time.Time
	end_generatig    time.Time
	start_table		 time.Time
	end_table		 time.Time
	start_match		 time.Time
	end_match		 time.Time

	// data
	source        []rune
	pattern       []rune
	skip_table    map[rune]int
	pattern_index int
}

type Queue struct {
	queue [](*PipeTask)
	size  int
}
