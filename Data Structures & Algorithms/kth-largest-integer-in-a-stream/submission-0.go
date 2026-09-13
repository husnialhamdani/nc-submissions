// 1. Define a custom type for the min-heap
type IntHeap []int

// 2. Implement the heap.Interface methods
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap logic
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 3. Define the KthLargest struct using your new IntHeap
type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	
	// Push all elements into the heap
	for _, num := range nums {
		heap.Push(h, num)
	}
	
	// Keep only the k largest elements
	for h.Len() > k {
		heap.Pop(h)
	}
	
	return KthLargest{minHeap: h, k: k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	
	// If size exceeds k, remove the smallest element
	if this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	
	// The root of the min-heap is always the k-th largest element
	return (*this.minHeap)[0]
}