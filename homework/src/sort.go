package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {    
	size := 10											              // (1)
	items := make([]int, size, size)					              // (2)
	rand.Seed(time.Now().UnixNano())					              // (3)
	for i := 0; i < size; i++ {							              // (4)
		items[i] = rand.Intn(999) - rand.Intn(999)		              // (5)
	}	              
	var (              
		gap = len(items)								              // (6)
		shrink = 1.3									              // (7)
		swapped = true									              // (8)			
	)              
	            
	for swapped {										              // (9)
			swapped = false									          // (10)
        	gap = int(float64(gap) / shrink)			              // (11)
        	if gap < 1 {								              // (12)
				gap = 1										          // (13)
			}              
		for i := 0; i + gap < size; i++ {					          // (14)
			if items[i] > items[i + gap] {				              // (15)
				items[i + gap], items[i] = items[i], items[i + gap]	  // (16)
				swapped = true									  	  // (17)
			}	
		}
	}
	fmt.Println(items)
}