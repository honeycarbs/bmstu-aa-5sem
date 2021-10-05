package sort

type IArray []int 

func InsertionSort(items IArray) {
	size := len(items)
	for i := 0; i < size; i++ {
		key := items[i]
		j := i
		for j > 0 && items[j-1] > key {
			items[j] = items[j-1]
			j--
		}
		items[j] = key
	}
}

func SelectionSort(items IArray) {
	l := len(items)
	for i := 0; i < l; i++ {
		min_i := i
		for j := i; j < l; j++ {
			if items[j] < items[min_i] {
				min_i = j
			}
		}
		items[i], items[min_i] = items[min_i], items[i]
	}
}

func HeapSort(items IArray) {

	for i := (len(items) - 1) / 2; i >= 0; i-- {
		_siftDown(items, i, len(items))
	}

	for i := len(items) - 1; i > 0; i-- {
		items[0], items[i] = items[i], items[0]
		_siftDown(items, 0, i)
	}
}


func _siftDown(heap []int, bttm, top int) {
	root := bttm
	for {
		child := root*2 + 1
		if child >= top {
			break
		}
		if child + 1 < top && heap[child] < heap[child + 1] {
			child++
		}
		if heap[root] < heap[child] {
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		} else {
			break
		}

	}
}