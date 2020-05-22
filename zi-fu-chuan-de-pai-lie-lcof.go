func permutation(s string) []string {
    nums := []byte(s)
    sort.Slice(nums, func (i, j int) bool { return nums[i] < nums[j] })
    return perm(string(nums))
}

func perm(s string) []string {
    if len(s) == 1 {
        return []string{s}
    }
    ans := []string{}
    for i := range s {
        if i > 0 && s[i] == s[i-1] {
            continue
        }
        for _, v := range perm(strings.Join([]string{s[:i], s[i+1:]}, "")) {
            ans = append(ans, string(s[i]) + v)
        }
    }
    return ans
}
