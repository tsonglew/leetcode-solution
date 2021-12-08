func maxSumOfThreeSubarrays(nums []int, k int) []int {
    array1Index := 0
    array2IndexStart, _ := getArray23Index(array1Index, k)
    array1Sum, array2Sum, array3Sum := 0, 0, 0
    array1Max, array2Max, array3Max := 0, 0, 0
    maxSum := 0
    result := []int{0, 0, 0}

    array1MaxStart, array2MaxStart, array3MaxStart := -1, -1, -1
    for ;array1Index <= len(nums) - 3 * k; array1Index++ {
        array2Index, array3Index := getArray23Index(array1Index, k)
        array1Sum += nums[array1Index]
        array2Sum += nums[array2Index]
        array3Sum += nums[array3Index]

        if array1Index < array2IndexStart - 1 {
            continue
        }

        if array1Sum > array1Max {
            array1Max = array1Sum
            array1MaxStart = array1Index-k+1
        }
        if array2Sum > array2Max {
            array2Max = array2Sum
            array2MaxStart = array2Index-k+1
        }
        if array3Sum > array3Max {
            array3Max = array3Sum
            array3MaxStart = array3Index-k+1
        }
        if array1Sum + array2Sum + array3Sum > maxSum {
            maxSum = array1Sum + array2Sum + array3Sum
            result = []int{array1MaxStart, array2MaxStart, array3MaxStart}
        }
        array1Sum -= nums[array1Index-k+1]
        array2Sum -= nums[array2Index-k+1]
        array3Sum -= nums[array3Index-k+1]
    }
    return result
}
