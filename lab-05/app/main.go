package main

import (
	"fmt"
	"runtime"

	"./pipeline"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// count := 20
	var count int
	fmt.Print("Queue size: ")
	fmt.Scan(&count)
	ch := make (chan int)
	pipeline_qu := pipeline.Pipeline(count, ch)
	_ = pipeline_qu
	<- ch


	pipeline_qu_sync := pipeline.Sync(count)
	_ = pipeline_qu_sync

	pipeline.PerfLog(pipeline_qu, pipeline_qu_sync)
	// pipeline.Log(pipeline_qu_sync)
	// pipeline.Log(pipeline_qu)	
}

