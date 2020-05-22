func findKthLargest(nums []int, k int) int {
	buildHeap(nums, len(nums))
	for i := 0; i < k; i++ {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		heapify(nums, 0, len(nums)-i-1)
	}
	return nums[len(nums)-k]
}

func buildHeap(heap []int, n int) {
	for i := (n-1)/2; i >= 0; i-- {
		heapify(heap, i, len(heap))
	}
}

func heapify(heap []int, i, n int) {
	for {
		maxPos := i
		if i*2+1 < n && heap[i*2+1] > heap[maxPos] {
			maxPos = i*2+1
		}
		if i*2+2 < n && heap[i*2+2] > heap[maxPos] {
			maxPos = i*2+2
		}
		if maxPos == i {
			break
		}
		heap[i], heap[maxPos] = heap[maxPos], heap[i]
		i = maxPos
	}
}
