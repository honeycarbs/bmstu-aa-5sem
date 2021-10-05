package generator
import (
	"math/rand"
	"time"
	"sort"
)

func Slice(size int, sorted bool) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	if sorted {
		sort.Ints(slice)
	}
	return slice
}