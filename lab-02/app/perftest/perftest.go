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

	"../mtx"
	mult "../multiply"
)

func prettyProfilePrint(size int, alg_time, w_time, w_opt_time C.double) {
    fmt.Printf("%d ", size)
    C.CPrint(alg_time)
    C.fflush(C.stdout)
    C.CPrint(w_time)
    C.fflush(C.stdout)
    C.CPrint(w_opt_time)
    C.fflush(C.stdout)
    fmt.Println("")
}

func ProfileWrapper() {
    for i := 10; i <= 100; i+= 10 {
        var left mtx.Matrix = mtx.Allocate(i, i)
        var right mtx.Matrix = mtx.Allocate(i, i)

        alg_time := __algBench(left, right)
        w_time := __winBench(left, right)
        wo_time := __winOptBench(left, right)
        prettyProfilePrint(i, alg_time, w_time, wo_time)
    }
    for i := 150; i <= 600; i+= 50 {
        var left mtx.Matrix = mtx.Allocate(i, i)
        var right mtx.Matrix = mtx.Allocate(i, i)

        alg_time := __algBench(left, right)
        w_time := __winBench(left, right)
        wo_time := __winOptBench(left, right)
        prettyProfilePrint(i, alg_time, w_time, wo_time)
    }

    fmt.Println("--- odd ---")

    for i := 10; i <= 100; i+= 10 {
        var left mtx.Matrix = mtx.Allocate(i + 1, i + 1)
        var right mtx.Matrix = mtx.Allocate(i + 1, i + 1)

        alg_time := __algBench(left, right)
        w_time := __winBench(left, right)
        wo_time := __winOptBench(left, right)
        prettyProfilePrint(i, alg_time, w_time, wo_time)
    }

    for i := 150; i <= 600; i+= 50 {
        var left mtx.Matrix = mtx.Allocate(i + 1, i + 1)
        var right mtx.Matrix = mtx.Allocate(i + 1, i + 1)

        alg_time := __algBench(left, right)
        w_time := __winBench(left, right)
        wo_time := __winOptBench(left, right)
        prettyProfilePrint(i, alg_time, w_time, wo_time)
    }
}

func __algBench(left, right mtx.Matrix) C.double {
    var count int = 1000
    var start, end C.double
    var sum_elapsed C.double = 0

    for i := 0; i < count; i++ {
        start = C.GetCpuTime()
        mult.Algebraic(left, right)
        end = C.GetCpuTime()

        sum_elapsed += end - start
    }
    return sum_elapsed / C.double(count)
}

func __winBench(left, right mtx.Matrix) C.double {
    var count int = 1000
    var start, end C.double
    var sum_elapsed C.double = 0

    for i := 0; i < count; i++ {
        start = C.GetCpuTime()
        mult.Winograd(left, right)
        end = C.GetCpuTime()

        sum_elapsed += end - start
    }
    return sum_elapsed / C.double(count)
}

func __winOptBench(left, right mtx.Matrix) C.double {
    var count int = 1000
    var start, end C.double
    var sum_elapsed C.double = 0

    for i := 0; i < count; i++ {
        start = C.GetCpuTime()
        mult.WinogradOptimized(left, right)
        end = C.GetCpuTime()

        sum_elapsed += end - start
    }
    return sum_elapsed / C.double(count)
}