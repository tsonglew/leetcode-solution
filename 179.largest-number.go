import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.SliceStable(nums, func(a, b int) bool {
		as := strconv.Itoa(nums[a])
		bs := strconv.Itoa(nums[b])
		flag := false
		if len(as) < len(bs) {
			if as[0] < as[len(as)-1] {
				flag = true
			}
			as = as + strings.Repeat(string(as[0]), len(bs)-len(as))
		}
		if len(as) > len(bs) {
			if bs[0] < bs[len(bs)-1] {
				flag = true
			}
			bs = bs + strings.Repeat(string(bs[0]), len(as)-len(bs))
		}
		if as == bs {
			if flag {
				return nums[a] < nums[b]
			}
			return nums[a] > nums[b]
		}
		return as > bs
	})
	result := ""
	for i := range nums {
		result += strconv.Itoa(nums[i])
	}
	result = strings.TrimLeft(result, "0")
	if result == "" {
		result = "0"
	}
	return result
}
