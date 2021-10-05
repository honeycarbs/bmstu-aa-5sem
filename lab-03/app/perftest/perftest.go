package perftest

/*
#include <pthread.h>
#include <time.h>
#include <stdio.h>

static double GetCpuTime() {
    struct timespec time;
    if (clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &time)) {
        perror("clock_gettime");
        return 0;
    }

    long seconds = time.tv_sec;
    long nanoseconds = time.tv_nsec;
    double elapsed = seconds + nanoseconds*1e-9;

    return elapsed;
}

static void CPrint(double num) {
    printf("%.2e ", num);
}
*/
import "C"

import (
	"fmt"

	_gen "../generator"
	_sort "../sort"
)
type ARRAY_STATE_ID int
const (
	RANDOM ARRAY_STATE_ID = iota
    SORTED
    REVERSED
)

func prettyProfilePrint(insert_time, select_time, heap_time C.double) {
    C.CPrint(insert_time)
    C.fflush(C.stdout)
    C.CPrint(select_time)
    C.fflush(C.stdout)
    C.CPrint(heap_time)
    C.fflush(C.stdout)
}

func ProfileWrapper() {
    data_random := __initTestingData(RANDOM)
    // data_sorted := __initTestingData(SORTED)
    // data_reversed := __initTestingData(REVERSED)

	for i:= 0; i < len(data_random); i++ { 
        fmt.Println(len(data_random[i]))

		// prettyProfilePrint(__insertionBench(data_random[i]),
        // __selectionsBench(data_random[i]), __heapBench(data_random[i]))

        // prettyProfilePrint(__insertionBench(data_sorted[i]),
        // __selectionsBench(data_sorted[i]), __heapBench(data_sorted[i]))

        // prettyProfilePrint(__insertionBench(data_reversed[i]),
        // __selectionsBench(data_reversed[i]), __heapBench(data_reversed[i]))

        // fmt.Println("")
	} 
}

func __reverse(array []int) []int {
	for i := 0; i < len(array)/2; i++ {
		j := len(array) - i - 1
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func __initTestingData(id ARRAY_STATE_ID) [][]int {
    var res [][]int

    var sorted bool = true

    if (id == RANDOM) {
        sorted = false
    }
    res = __initHelper(10, 50, 10, id, sorted, res)
    res = __initHelper(50, 500, 50, id, sorted, res)
    res = __initHelper(500, 2500, 100, id, sorted, res)

    return res
}

func __initHelper(start, stop, step int, id ARRAY_STATE_ID, sorted bool, res [][]int) [][]int {
    var item []int
    for i := start; i <= stop; i += step {
        item = _gen.Slice(i, sorted)
        if (id == REVERSED) {
            __reverse(item)
        }
        res = append(res, item)
    }
    return res
}

func __insertionBench(testArray _sort.IArray) C.double {
    var count int = 10000
    var start, end C.double
    var sum_elapsed C.double = 0

    for i := 0; i < count; i++ {
        var dummy _sort.IArray = make([]int, len(testArray))
        copy(dummy, testArray) 
        start = C.GetCpuTime()
        _sort.InsertionSort(dummy)
        end = C.GetCpuTime()

        sum_elapsed += end - start
    }
    return sum_elapsed / C.double(count)
}

func __selectionsBench(testArray _sort.IArray) C.double {
    var count int = 10000
    var start, end C.double
    var sum_elapsed C.double = 0

    for i := 0; i < count; i++ {
        var dummy _sort.IArray = make([]int, len(testArray))
        copy(dummy, testArray) 
        start = C.GetCpuTime()
        _sort.SelectionSort(dummy)
        end = C.GetCpuTime()

        sum_elapsed += end - start
    }
    return sum_elapsed / C.double(count)
}

func __heapBench(testArray _sort.IArray) C.double {
    var count int = 10000
    var start, end C.double
    var sum_elapsed C.double = 0

    for i := 0; i < count; i++ {
        var dummy _sort.IArray = make([]int, len(testArray))
        copy(dummy, testArray) 
        start = C.GetCpuTime()
        _sort.HeapSort(dummy)
        end = C.GetCpuTime()

        sum_elapsed += end - start
    }
    return sum_elapsed / C.double(count)
}
